package memory

import (
	asst "github.com/stretchr/testify/assert"
	"github.com/xephonhq/xephon-k/pkg/common"
	"testing"
)

func TestIntersect(t *testing.T) {
	assert := asst.New(t)
	l1 := []common.SeriesID{1, 2, 3}
	l2 := []common.SeriesID{1, 2, 3, 4}
	l3 := []common.SeriesID{1, 2}
	l4 := []common.SeriesID{2}

	assert.Equal([]common.SeriesID{2}, Intersect(l1, l2, l3, l4))

	// longer list reaches end before short list does
	// NOTE: should see `break outer in 1 th list` if the log is enabled
	assert.Equal([]common.SeriesID{1}, Intersect([]common.SeriesID{1, 8, 11}, []common.SeriesID{1, 4, 5, 6}))
}

func TestUnion(t *testing.T) {
	assert := asst.New(t)

	// simple union without duplication
	l1 := []common.SeriesID{1, 2}
	l2 := []common.SeriesID{3}
	assert.Equal([]common.SeriesID{1, 2, 3}, Union(l1, l2))

	// duplication
	l3 := []common.SeriesID{1, 2}
	l4 := []common.SeriesID{1, 2, 3}
	// FIXED: result is n1, n2, n2, by adding
	//if lastVal == smallestVal {
	//	continue
	//}
	assert.Equal([]common.SeriesID{1, 2, 3}, Union(l3, l4))

	//log.Logger.EnableSourceLine()
	// three way
	l5 := []common.SeriesID{9}
	//fmt.Println("index_test.go:44") // works for IDEA
	//fmt.Println("source=index_test.go:44") // does not works for IDEA
	//fmt.Println("source= index_test.go:44") // works for IDEA
	// FIXED: {1, 2, 9, 3} would show up randomly, sometimes the test just pass
	// Got!! when we deal with dup, we need to use its next value to compare if there is any, the random is caused by picking the first value
	assert.Equal([]common.SeriesID{1, 2, 3, 9}, Union(l3, l4, l5))
	//log.Logger.DisableSourceLine()

	// just want to write one more test
	l6 := []common.SeriesID{1, 9}
	assert.Equal([]common.SeriesID{1, 2, 3, 9}, Union(l3, l4, l5, l6))
}

func TestIndex_Filter(t *testing.T) {
	assert := asst.New(t)
	idx := NewIndex(5)
	idx.Add(common.SeriesID(1), "app", "nginx")
	idx.Add(common.SeriesID(1), nameTagKey, "response_time")
	idx.Add(common.SeriesID(2), "app", "apache")
	idx.Add(common.SeriesID(2), nameTagKey, "response_time")
	filterResponse := common.Filter{Type: "tag_match", Key: nameTagKey, Value: "response_time"}
	assert.Equal([]common.SeriesID{1, 2}, idx.Filter(&filterResponse))
	filterNginxResponse := common.Filter{Type: "and",
		LeftOperand:  &common.Filter{Type: "tag_match", Key: nameTagKey, Value: "response_time"},
		RightOperand: &common.Filter{Type: "tag_match", Key: "app", Value: "nginx"}}
	assert.Equal([]common.SeriesID{1}, idx.Filter(&filterNginxResponse))
}

func TestIndex_Get(t *testing.T) {
	assert := asst.New(t)
	idx := NewIndex(1)
	idx.Add(common.SeriesID(1), "app", "nginx")
	assert.Equal([]common.SeriesID{1}, idx.Get("app", "nginx"))
	assert.Equal([]common.SeriesID{}, idx.Get("foo", "bar"))
	assert.Equal(0, len(idx.Get("foo", "bar")))
}

func TestIndex_Add(t *testing.T) {
	assert := asst.New(t)
	idx := NewIndex(1)
	idx.Add(common.SeriesID(1), "app", "nginx")
	// NOTE: they just share one tag pair, so they can have different series ID
	idx.Add(common.SeriesID(2), "app", "nginx")
	idx.Add(common.SeriesID(3), "app", "apache")
	assert.Equal(true, idx.tagKeyIndex["app"]["nginx"])
	assert.Equal(true, idx.tagKeyIndex["app"]["apache"])
	assert.Equal(false, idx.tagKeyIndex["app"]["IIS"])
	// FIXME: currently we don't add separator between tagKey and tagValue
	assert.Equal([]common.SeriesID{1, 2}, idx.invertedIndexes["appnginx"].Postings)
}

func TestInvertedIndex_Add(t *testing.T) {
	assert := asst.New(t)
	iidx := newInvertedIndex("foo")
	iidx.Add(common.SeriesID(5))
	iidx.Add(common.SeriesID(1))
	iidx.Add(common.SeriesID(2))
	iidx.Add(common.SeriesID(2))
	iidx.Add(common.SeriesID(4))
	assert.Equal(iidx.Postings, []common.SeriesID{1, 2, 4, 5})
}

func TestMin(t *testing.T) {
	assert := asst.New(t)
	assert.Equal(1, min(1, 2))
	assert.Equal(1, min(2, 1))
}
