package memory

import (
	"sort"
	"sync"

	"github.com/xephonhq/xephon-k/pkg/common"
)

// IntSeriesStore protects the underlying IntSeries with a RWMutex
type IntSeriesStore struct {
	mu     sync.RWMutex
	series common.IntSeries
	length int
}

type DoubleSeriesStore struct {
	mu     sync.RWMutex
	series common.DoubleSeries
	length int
}

// NewIntSeriesStore creates a IntSeriesStore
func NewIntSeriesStore() *IntSeriesStore {
	series := common.IntSeries{}
	series.Points = make([]common.IntPoint, initPointsLength)
	return &IntSeriesStore{series: series, length: 0}
}

// WriteSeries merges the new series with existing one and replace old points with new points if their timestamp matches
// TODO: what happens when no memory is available? maybe this function should return error
func (store *IntSeriesStore) WriteSeries(newSeries common.IntSeries) {
	store.mu.Lock()
	defer store.mu.Unlock()

	// merge the old series with new series
	// TODO: check if they are the same series
	// TODO: efficient merge sort
	// TODO: actually we can remove duplicate when merge by comparing with previous point
	// TODO: if the start of the new series is smaller than the end of existing and we make sure there is no duplication we can simply append it

	// TODO: add a flag to series, so we don't sort points that are already sorted
	// store.series should already be sorted, so we only sort the newSeries
	sort.Sort(common.IntPoints(newSeries.Points))
	i := 0
	j := 0
	k := 0
	// NOTE: we can't use len(store.series.Points) because there might be duplicate
	oldLength := store.length
	newLength := len(newSeries.Points)
	// log.Infof("ol %d nl %d", oldLength, newLength)
	points := make([]common.IntPoint, oldLength+newLength)
	for i < oldLength && j < newLength {
		if store.series.Points[i].TimeNano < newSeries.Points[j].TimeNano {
			points[k] = store.series.Points[i]
			i++
		} else if store.series.Points[i].TimeNano == newSeries.Points[j].TimeNano {
			// if there is duplicate, overwrite the old point with new point
			points[k] = newSeries.Points[j]
			i++
			j++
		} else {
			points[k] = newSeries.Points[j]
			j++
		}
		// log.Infof("value in loop is %v", points[k].T)
		k++
	}
	// log.Infof("i %d j %d k %d", i, j, k)
	store.length = k

	// copy what is left, should only have one array left
	// https://github.com/golang/go/wiki/SliceTricks
	if i < oldLength {
		// FIXED: should cut ... instead of simply append
		points = append(points[:k], store.series.Points[i:]...)
		store.length = k + oldLength - i
	}
	if j < newLength {
		points = append(points[:k], newSeries.Points[j:]...)
		store.length = k + newLength - j
	}
	log.Debugf("length %d", store.length)

	// TODO: maybe using pool is a good idea since there are a lot of merge when inserting series
	store.series.Points = points

	//n := 0
	//for n < store.length {
	//	log.Infof("time in store %v", store.series.Points[n].T)
	//	n++
	//}
}

// ReadByStartEndTime filters and return a copy of the data
// TODO: we were previously returning *common.IntSeries, but there should not have any copy of the underlying points I
// suppose?
func (store *IntSeriesStore) ReadByStartEndTime(startTime int64, endTime int64) common.IntSeries {
	store.mu.RLock()
	defer store.mu.RUnlock()
	log.Trace("read the series!")
	log.Tracef("store length %d", store.length)

	// TODO: may use pool for points
	returnSeries := common.IntSeries{
		Name: store.series.Name,
		Tags: store.series.Tags,
	}
	// TODO: we can assume the data has fixed interval and start from a more close position instead of zero
	// TODO: we can simply copy by the start and end instead of going it one by one since the int series store should already be sorted
	// FIXME: at least use binary/exponential search for start and end instead of compare one by one
	// TODO: may use skip list
	// TODO: run the aggregation here?
	for i := 0; i < store.length; i++ {
		// FIXED: wrong fake time in test
		//1359788400002
		//144788400002
		if store.series.Points[i].TimeNano >= startTime && store.series.Points[i].TimeNano <= endTime {
			// TODO: maybe we should add a append method to IntSeries and let it create a new copy of the point
			log.Tracef("need to append the points!")
			returnSeries.Points = append(returnSeries.Points, store.series.Points[i])
		}
	}
	return returnSeries
}
