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
	// TODO: find the right series and write to it
	log.Info("need to write sth!")
	return nil
}
