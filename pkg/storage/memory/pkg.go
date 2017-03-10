package memory

import (
	"github.com/xephonhq/xephon-k/pkg/common"
	"github.com/xephonhq/xephon-k/pkg/util"
)

var log = util.Logger.NewEntryWithPkg("k.s.mem")
var initSeriesNumber = 10

// Store is the in memory storage with data and index
type Store struct {
	data  Data
	index Index
}

// NewMemStore creates an in memory storage with small allocated space
func NewMemStore() *Store {
	store := Store{}
	// TODO: add a function to create the data and index?
	store.data = make(Data, initSeriesNumber)
	store.index = make(Index, initSeriesNumber)
	return &store
}

// StoreType implements Store interface
func (store Store) StoreType() string {
	return "memory"
}

// WriteIntSeries implements Store interface
func (store Store) WriteIntSeries(series []common.IntSeries) error {
	// FIXME: it seems it is an empty series ...
	// TODO: find the right series and write to it
	for _, oneSeries := range series {
		id := SeriesID(oneSeries.Hash())
		// TODO: this logic should be moved to data
		seriesStore, ok := store.data[id]
		if ok {
			log.Info("merge with existing series")
			seriesStore.WriteSeries(oneSeries)
		} else {
			log.Info("create new entry in map")
			// FIXME: unify pointer and value
			store.data[id] = *NewIntSeriesStore()
			// FIXED: http://stackoverflow.com/questions/32751537/why-do-i-get-a-cannot-assign-error-when-setting-value-to-a-struct-as-a-value-i
			// store.data[id].series = oneSeries
			seriesStore = store.data[id]
			seriesStore.series = oneSeries
		}
	}

	log.Info("need to write sth!")
	return nil
}
