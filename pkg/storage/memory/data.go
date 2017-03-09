package memory

import (
	"github.com/xephonhq/xephon-k/pkg/common"
	"sync"
)

type SeriesID string

// TODO: should be able to allow double etc later
type Data map[SeriesID]IntSeriesStore

type IntSeriesStore struct {
	mu     sync.RWMutex
	series common.IntSeries
}

func (store IntSeriesStore) WriteSeries(newSeries common.IntSeries) {
	store.mu.Lock()
	defer store.mu.Unlock()
	// merge the old series with new series
	// TODO: efficient merge sort
	// store.series should already be sorted

}