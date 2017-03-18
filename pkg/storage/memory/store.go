package memory

import (
	"github.com/xephonhq/xephon-k/pkg/common"
	"sync"
)

// singleton
var storeMap StoreMap

// MemStore need to be shared between read and write services
type StoreMap struct {
	mu     sync.RWMutex
	stores map[string]*Store
}

func init() {
	storeMap.stores = make(map[string]*Store, 1)
	storeMap.stores["default"] = NewMemStore()
}

func GetDefaultMemStore() *Store {
	storeMap.mu.RLock()
	defer storeMap.mu.RUnlock()
	return storeMap.stores["default"]
}

// Store is the in memory storage with data and index
type Store struct {
	data  Data
	index Index
}

// NewMemStore creates an in memory storage with small allocated space
func NewMemStore() *Store {
	store := Store{}
	// TODO: add a function to create the data and index?
	store.data = make(Data, initSeriesCount)
	store.index = make(Index, initSeriesCount)
	return &store
}

// StoreType implements Store interface
func (store Store) StoreType() string {
	return "memory"
}

// QueryIntSeries implements Store interface
// TODO: this definitely won't work
func (store Store) QueryIntSeries(query common.Query) ([]common.IntSeries, error) {
	series := make([]common.IntSeries, 0)
	// TODO: use switch and not hard coded string
	if query.MatchPolicy == "exact" {
		// fetch the series
		// TODO: should we make a copy of the points, what would happen if there are
		// write when we are encoding it to json
		seriesID := SeriesID(query.Hash())
		// TODO: there is mutex on IntSeries store, how does prometheus etc. handle this?
		// should we have a get method or things like that?
		// prometheus use Iterator .... maybe we need custom implements
		oneSeries, ok := store.data[seriesID]
		if ok {
			// TODO: apply the time range filter
			//series = append(series, oneSeries.series)
			// TODO: use the start and end time from query
			// TODO: rename to ReadAll, ReadByStartEndTime etc.
			series = append(series, *oneSeries.ReadSeries(0, 1447884000020))
		}
		return series, nil
	}
	log.Warn("non exact match is not supported!")
	// TODO: query the index to do the filter
	return series, nil
}

// WriteIntSeries implements Store interface
func (store Store) WriteIntSeries(series []common.IntSeries) error {
	for _, oneSeries := range series {
		id := SeriesID(oneSeries.Hash())
		// TODO: this should return error and we should handle it somehow
		store.data.WriteIntSeries(id, oneSeries)
		// TODO: write the index
	}
	return nil
}

func (store Store) Shutdown() {
	log.Info("shutting down memoery store, nothing to do, have a nice weekend~")
}
