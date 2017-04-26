package memory

import (
	"github.com/xephonhq/xephon-k/pkg/common"
)

// Index is a map of inverted index with tag name as key and tag value as term for the inverted index
type Index struct {
	tagKeyIndex     map[string]map[string]bool // map[string]bool is used as set
	invertedIndexes map[string]*InvertedIndex
}

// InvertedIndex use Term for tag value postings for a list of sorted series ID
// TODO: Series ID should use locality sensitive hashing https://en.wikipedia.org/wiki/Locality-sensitive_hashing
type InvertedIndex struct {
	Term     string
	Postings []common.SeriesID
}

var initialPostingSize = 10

func NewIndex(capacity int) *Index {
	return &Index{
		tagKeyIndex:     make(map[string]map[string]bool, capacity),
		invertedIndexes: make(map[string]*InvertedIndex, capacity),
	}
}

func newInvertedIndex(term string) *InvertedIndex {
	return &InvertedIndex{
		Term:     term,
		Postings: make([]common.SeriesID, 0, initialPostingSize),
	}
}

func (idx *Index) Get(tagKey string, tagValue string) []common.SeriesID {
	term := tagKey + tagValue
	iidx, ok := idx.invertedIndexes[term]
	if ok {
		return iidx.Postings
	} else {
		return []common.SeriesID{}
	}
}

func (idx *Index) Add(id common.SeriesID, tagKey string, tagValue string) {
	// update tagKeyIndex
	_, ok := idx.tagKeyIndex[tagKey]
	if !ok {
		idx.tagKeyIndex[tagKey] = make(map[string]bool)
	}
	idx.tagKeyIndex[tagKey][tagValue] = true

	// TODO: should add separator, in Prometheus `db.go` it's `const sep = '\xff'`
	term := tagKey + tagValue
	// create the inverted index if not exists
	_, ok = idx.invertedIndexes[term]
	if !ok {
		idx.invertedIndexes[term] = newInvertedIndex(term)
	}
	idx.invertedIndexes[term].Add(id)
}

// TODO: actually we can have a fixed size map to cache the hot series, so there is no need to lookup if the id is already in there
func (iidx *InvertedIndex) Add(id common.SeriesID) {
	// binary search and insert the value if not found
	low, high := 0, len(iidx.Postings)
	for low < high {
		// TODO: use custom compare function or compare it directly using <
		mid := low + (high-low)/2 // avoid overflow, copied from `src/sort/search.go` sort.Search
		if iidx.Postings[mid] >= id {
			high = mid
		} else {
			low = mid + 1
		}
	}

	// not found
	if low == len(iidx.Postings) {
		iidx.Postings = append(iidx.Postings, id)
		return
	} else if iidx.Postings[low] != id {
		// insert it to the slice https://github.com/golang/go/wiki/SliceTricks#insert
		iidx.Postings = append(iidx.Postings, id) // we append id here, but any value is ok, it will be overwritten by following copy
		copy(iidx.Postings[low+1:], iidx.Postings[low:])
		iidx.Postings[low] = id

	}

	// found
	// TODO: should have some sort of cache
	return
}

// Intersect is used for AND, i.e. app=nginx AND os=ubuntu
// - sort lists by length
// - loop through the element in the shortest list,
// 	 - use exponential search to find if the element exists in other lists, only add it to result if it appears in all lists
//   - if any list reaches its end, the outer loop breaks
// NOTE:
// - we didn't use the algorithm in the VLDB paper, just a naive one with some similar ideas
// - in fact, this is just the `join` operation in RDBMS
// TODO:
// - it is also possible to sort by value range
// Ref
// - https://www.quora.com/Which-is-the-best-algorithm-to-merge-k-ordered-lists
// 	 - 'adaptive list intersection'
// - Improving performance of List intersection http://www.vldb.org/pvldb/2/vldb09-pvldb37.pdf
// 	 - Dynamic probe
// - Exponential (galloping) search https://en.wikipedia.org/wiki/Exponential_search
func Intersect(postings ...[]common.SeriesID) []common.SeriesID {
	// posting is a sorted list, see InvertedIndex
	// sort by list length using selection sort, assume the number of list is small
	listCount := len(postings)
	allLength := make([]int, listCount)
	// NOTE: probeStart is not used by sorting lists, we just use the loop to initialize all element to 1,
	// because exponential search can't start from 0, 0 * 2 = 0
	probeStart := make([]int, listCount)
	for i := 0; i < listCount; i++ {
		shortestIndex := i
		shortestLength := len(postings[i])
		for j := i + 1; j < listCount; j++ {
			curLength := len(postings[j])
			if curLength < shortestLength {
				shortestIndex = j
				shortestLength = curLength
			}
		}
		// swap if needed
		if i != shortestIndex {
			postings[i], postings[shortestIndex] = postings[shortestIndex], postings[i]
		}
		allLength[i] = shortestLength
		probeStart[i] = 1
	}

	// walk all the elements in the shortest list
	// assume the intersection is same length as the shortest list, allocate the space
	intersection := make([]common.SeriesID, 0, allLength[0])
OUTER:
	for i := 0; i < allLength[0]; i++ {
		cur := postings[0][i]
		// probe all the other lists, if one of them don't met, go to next loop
		for k := 1; k < listCount; k++ {
			// exponential search, use a smaller range for following binary search
			bound := probeStart[k]
			size := allLength[k]
			for bound < size && postings[k][bound] < cur {
				bound *= 2
			}

			// binary search
			low := bound / 2
			// NOTE: Go does not have `(a < b)? a : b` http://stackoverflow.com/questions/19979178/what-is-the-idiomatic-go-equivalent-of-cs-ternary-operator
			high := min(bound, size)
			for low < high {
				mid := low + (high-low)/2
				if postings[k][mid] >= cur {
					high = mid
				} else {
					low = mid + 1
				}
			}
			// this list reaches end, no need to continue the outer loop
			if low == size {
				// http://relistan.com/continue-statement-with-labels-in-go/
				//log.Infof("break outer in %d th list", k)
				break OUTER
			}
			probeStart[k] = low + 1
			// got the nearest one, but not the same one, no need to check other lists, continue the outer loo
			if postings[k][low] != cur {
				continue OUTER
			}
		}
		// if you made it here, then you are in all the lists
		intersection = append(intersection, cur)
	}
	return intersection
}

// Union is used for OR, i.e. app=nginx OR app=apache
// - sort all the lists by length? or just pick the smallest one?
// - get first len(smallest) elements of each array into an array and sort it? this is nk * log(k)
// NOTE
// - Linear search merge duplicate compare
// - Divide and Conquer merge requires extra space
// - Heap merge requires using Heap (e... such a brainless note, a.k.a I don't know how to write heap)
// - need to exclude lists when they reaches the end, might use a map
// Ref
// - https://en.wikipedia.org/wiki/K-Way_Merge_Algorithms
// - https://github.com/prometheus/tsdb/issues/50
// - k-way merging and k-ary sorts http://cs.uno.edu/people/faculty/bill/k-way-merge-n-sort-ACM-SE-Regl-1993.pdf
// - https://www.cs.cmu.edu/~adamchik/15-121/lectures/Binary%20Heaps/heaps.html
func Union(postings ...[]common.SeriesID) []common.SeriesID {
	listCount := len(postings)
	remainLists := make(map[int]bool, listCount)
	posList := make([]int, listCount)
	allLength := make([]int, listCount)
	for i := 0; i < listCount; i++ {
		remainLists[i] = true
		posList[i] = 0
		allLength[i] = len(postings[i])
	}

	// FIXME: this is linear search merge, the slowest one, nk, but when k is small, this is fine
	// TODO: it seems there is not need for sorting
	// TODO: capacity, sum of all lists?
	// TODO: need to handle duplication, the union should also be a set, and there could be duplication for sure
	union := make([]common.SeriesID, 0)
	impossibleLargeVal := common.SeriesID("ZZZZZZ")
	lastVal := impossibleLargeVal
	j := 0
	for len(remainLists) > 0 {
		log.Info(remainLists)
		if j > 5 {
			break
		}else {
			j++
		}
		// TODO: using map, you can't pick the first one as the smallest, unless you add a flag
		smallestVal := impossibleLargeVal
		smallestIndex := 0
		for i := range remainLists {
			curVal := postings[i][posList[i]]
			if curVal == lastVal {
				// duplication
				posList[i]++
				if posList[i] == allLength[i] {
					delete(remainLists, i)
				}
			} else if curVal < smallestVal {
				// smaller value
				smallestVal = curVal
				smallestIndex = i
			}
		}
		posList[smallestIndex]++
		if posList[smallestIndex] == allLength[smallestIndex] {
			delete(remainLists, smallestIndex)
		}
		union = append(union, smallestVal)
	}
	return union
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
