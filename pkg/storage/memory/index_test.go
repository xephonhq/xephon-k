package memory

import (
	"fmt"
	asst "github.com/stretchr/testify/assert"
	"github.com/xephonhq/xephon-k/pkg/common"
	"testing"
)

func TestIntersect(t *testing.T) {
	l1 := []common.SeriesID{"n1", "n2", "n3"}
	l2 := []common.SeriesID{"n1", "n2", "n3", "n4"}
	l3 := []common.SeriesID{"n1", "n2"}
	l4 := []common.SeriesID{"n2"}
	l5 := []common.SeriesID{"n5", "n6", "n7"}
	// test if print will sort the slice, nop
	fmt.Println([][]common.SeriesID{l1, l2, l3, l4})
	//Intersect(l1, l2)
	Intersect(l3, l5)
	//Intersect(l1, l2, l3, l4)
}

func TestIndex_Get(t *testing.T) {
	assert := asst.New(t)
	idx := NewIndex(1)
	idx.Add(common.SeriesID("n1"), "app", "nginx")
	assert.Equal([]common.SeriesID{"n1"}, idx.Get("app", "nginx"))
	assert.Equal([]common.SeriesID{}, idx.Get("foo", "bar"))
	assert.Equal(0, len(idx.Get("foo", "bar")))
}

func TestIndex_Add(t *testing.T) {
	assert := asst.New(t)
	idx := NewIndex(1)
	idx.Add(common.SeriesID("n1"), "app", "nginx")
	idx.Add(common.SeriesID("n2"), "app", "nginx")
	idx.Add(common.SeriesID("a1"), "app", "apache")
	assert.Equal(true, idx.tagKeyIndex["app"]["nginx"])
	assert.Equal(true, idx.tagKeyIndex["app"]["apache"])
	assert.Equal(false, idx.tagKeyIndex["app"]["IIS"])
	// FIXME: currently we don't add separator between tagKey and tagValue
	assert.Equal([]common.SeriesID{"n1", "n2"}, idx.invertedIndexes["appnginx"].Postings)
}

func TestInvertedIndex_Add(t *testing.T) {
	assert := asst.New(t)
	iidx := newInvertedIndex("foo")
	iidx.Add(common.SeriesID("d"))
	iidx.Add(common.SeriesID("e"))
	iidx.Add(common.SeriesID("a"))
	iidx.Add(common.SeriesID("a"))
	iidx.Add(common.SeriesID("b"))
	assert.Equal(iidx.Postings, []common.SeriesID{"a", "b", "d", "e"})
}
