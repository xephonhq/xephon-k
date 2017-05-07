package memory

import "github.com/xephonhq/xephon-k/pkg/common"

// Store is the in memory storage with data and index
type Store struct {
	data  Data
	index *Index // TODO: might change to value instead of pointer
}

// NewMemStore creates an in memory storage with small allocated space
func NewMemStore() *Store {
	store := Store{}
	// TODO: add a function to create the data?
	store.data = make(Data, initSeriesCount)
	store.index = NewIndex(initSeriesCount)
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
	// TODO: not hard coded string
	switch query.MatchPolicy {
	case "exact":
		// fetch the series
		seriesID := query.Hash()
		// TODO: should we make a copy of the points, what would happen if there are
		// write when we are encoding it to json
		// TODO: there is mutex on IntSeries store, how does prometheus etc. handle this?
		// should we have a get method or things like that?
		// prometheus use Iterator .... maybe we need custom implements, I think it also have blocks
		oneSeries, ok := store.data[seriesID]
		if ok {
			series = append(series, *oneSeries.ReadByStartEndTime(query.StartTime, query.EndTime))
		}
		return series, nil
	case "filter":
		// TODO: real filter
		log.Warn("TODO: write code for filter")
	default:
		// TODO: query the index to do the filter
		log.Warn("non exact match is not supported!")
	}
	return series, nil
}

// WriteIntSeries implements Store interface
func (store Store) WriteIntSeries(series []common.IntSeries) error {
	for _, oneSeries := range series {
		id := oneSeries.Hash()
		// TODO: this should return error and we should handle it somehow
		// Write Data
		store.data.WriteIntSeries(id, oneSeries)
		// Write Index
		// NOTE: we store series name as special tag
		store.index.Add(id, nameTagKey, oneSeries.Name)
		for k, v := range oneSeries.Tags {
			store.index.Add(id, k, v)
		}
	}
	return nil
}

// Shutdown TODO: gracefully flush in memory data to disk
func (store Store) Shutdown() {
	log.Info("shutting down memoery store, nothing to do, have a nice weekend~")
}
