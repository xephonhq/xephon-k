package memory

import (
	"github.com/pkg/errors"
	"github.com/xephonhq/xephon-k/pkg/common"
)

// Store is the in memory storage with data and index
type Store struct {
	config Config
	data   *Data
	index  *Index // TODO: might change to value instead of pointer (why I said that?)
}

// NewMemStore creates an in memory storage with small allocated space
// Deprecated
func NewMemStore() *Store {
	store := &Store{}
	store.data = NewData(initSeriesCount)
	store.index = NewIndex(initSeriesCount)
	return store
}

func NewMemStore2(config Config) *Store {
	return &Store{
		config: config,
		data:   NewData(initSeriesCount),
		index:  NewIndex(initSeriesCount),
	}
}

// StoreType implements Store interface
func (store *Store) StoreType() string {
	return "memory"
}

// QuerySeries implements Store interface
func (store *Store) QuerySeries(queries []common.Query) ([]common.QueryResult, []common.Series, error) {
	result := make([]common.QueryResult, 0, len(queries))
	series := make([]common.Series, 0, len(queries))
	// TODO:
	// - first look up the series id
	// - add match number
	// - read the data by time range
	// - apply the aggregator when look up?
	// - test it in non e2e test
	for i := 0; i < len(queries); i++ {
		query := queries[i]
		queryResult := common.QueryResult{Query: query, Matched: 0}
		switch query.MatchPolicy {
		case "exact":
			seriesID := common.Hash(&query)
			s, ok, err := store.data.ReadSeries(seriesID, query.StartTime, query.EndTime)
			if ok {
				queryResult.Matched = 1
				series = append(series, s)
			}
			if err != nil {
				// TODO: wrap the error
				return result, series, err
			}
		case "filter":
			// TODO: we should also expose a HTTP API for query series ID only
			// FIXME: this is a dirty hack to be compatible with the Name filed in the query, it is treated as __name__ tag
			// need to make a shallow copy, otherwise it will refer to itself and cause stackoverflow
			originalFilter := query.Filter
			query.Filter = common.Filter{Type: "and", LeftOperand: &common.Filter{Type: "tag_match", Key: nameTagKey, Value: query.Name},
				RightOperand: &originalFilter}
			seriesIDs := store.index.Filter(&query.Filter)
			queryResult.Matched = len(seriesIDs)
			for j := 0; j < len(seriesIDs); j++ {
				// TODO: let's just assume all series in the index is all in the memory, so we don't check the data map
				seriesID := seriesIDs[j]
				s, ok, err := store.data.ReadSeries(seriesID, query.StartTime, query.EndTime)
				if ok {
					series = append(series, s)
				}
				if err != nil {
					// TODO: wrap the error
					return result, series, err
				}
			}
		default:
			// TODO: return error to warn the user
			log.Warn("unsupported match policy %s", query.MatchPolicy)
		}
		result = append(result, queryResult)
	}
	return result, series, nil
}

// WriteIntSeries implements Store interface
func (store *Store) WriteIntSeries(series []common.IntSeries) error {
	for i := 0; i < len(series); i++ {
		id := common.Hash(&series[i])
		// Write Data
		err := store.data.WriteIntSeries(id, series[i])
		if err != nil {
			return errors.Wrapf(err, "write data failed for %s %v", series[i].Name, series[i].Tags)
		}
		// Write Index
		// TODO: write index and write data can be parallel, though I don't know if it has performance boost
		// TODO: write index should also have error
		// NOTE: we store series name as special tag
		store.index.Add(id, nameTagKey, series[i].Name)
		for k, v := range series[i].Tags {
			store.index.Add(id, k, v)
		}
	}
	return nil
}

// WriteDoubleSeries implements Store interface
func (store *Store) WriteDoubleSeries(series []common.DoubleSeries) error {
	for i := 0; i < len(series); i++ {
		id := common.Hash(&series[i])
		// Write Data
		err := store.data.WriteDoubleSeries(id, series[i])
		if err != nil {
			return errors.Wrapf(err, "write data failed for %s %v", series[i].Name, series[i].Tags)
		}
		// Write Index
		// TODO: write index and write data can be parallel, though I don't know if it has performance boost
		// TODO: write index should also have error
		// NOTE: we store series name as special tag
		store.index.Add(id, nameTagKey, series[i].Name)
		for k, v := range series[i].Tags {
			store.index.Add(id, k, v)
		}
	}
	return nil
}

// Shutdown
func (store *Store) Shutdown() {
	// TODO: ask user if they want to flush in memory data to disk
	log.Info("shutting down memory store, nothing to do, have a nice weekend~")
}
