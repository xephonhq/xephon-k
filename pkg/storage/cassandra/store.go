package cassandra

import (
	"github.com/gocql/gocql"
	"github.com/xephonhq/xephon-k/pkg/common"
	//"time"
	"github.com/pkg/errors"
)

// Store contains a cassandra session
type Store struct {
	config  Config
	session *gocql.Session
}

// NewCassandraStore creates a new cassandra store connecting to localhost cassandra
func NewCassandraStore(config Config) (*Store, error) {
	store := &Store{
		config: config,
	}
	// connect to cassandra
	cluster := gocql.NewCluster(config.Host)
	cluster.Port = config.Port
	cluster.Keyspace = defaultKeySpace
	session, err := cluster.CreateSession()
	if err != nil {
		return nil, errors.Wrapf(err, "can't connect to %s:%d", config.Host, config.Port)
	} else {
		log.Infof("connected to cassandra %s:%d", config.Host, config.Port)
	}
	store.session = session
	return store, nil
}

// StoreType implements store interface
func (store Store) StoreType() string {
	return "cassandra"
}

func (store Store) QuerySeries(queries []common.Query) ([]common.QueryResult, []common.Series, error) {
	log.Panic("not implemented!")

	result := make([]common.QueryResult, 0)
	series := make([]common.Series, 0)
	return result, series, nil
}

// QueryIntSeriesBatch implements Store interface
func (store Store) QueryIntSeriesBatch(queries []common.Query) ([]common.QueryResult, []common.IntSeries, error) {
	log.Panic("not implemented!")

	result := make([]common.QueryResult, 0)
	series := make([]common.IntSeries, 0)
	return result, series, nil
}

// QueryIntSeries implements Store interface
// Deprecated: Use QueryIntSeriesBatch instead
func (store Store) QueryIntSeries(query common.Query) ([]common.IntSeries, error) {
	series := make([]common.IntSeries, 0)
	session := store.session

	if query.MatchPolicy == "exact" {
		iter := session.Query(selectIntStmt, query.Name, query.Tags).Iter()
		// NOTE: both time and int64 works
		// var metricTimestamp time.Time
		var metricTimestamp int64
		var metricValue int64
		oneSeries := common.IntSeries{}
		// TODO: may specify capacity to improve performance
		oneSeries.Points = make([]common.IntPoint, 0)
		for iter.Scan(&metricTimestamp, &metricValue) {
			oneSeries.Points = append(oneSeries.Points, common.IntPoint{T: metricTimestamp, V: metricValue})
			//log.Infof("%v %d", metricTimestamp, metricValue)
		}
		if err := iter.Close(); err != nil {
			return series, err
		}
		return series, nil
	}
	log.Warn("non exact match is not supported!")

	return series, nil
}

// WriteIntSeries implements Store interface
func (store Store) WriteIntSeries(series []common.IntSeries) error {
	// TODO: we should use different goroutine for different series according to official doc
	// Use many goroutines when doing inserts, the driver is asynchronous but provides a synchronous API,
	// it can execute many queries concurrently
	session := store.session
	for _, oneSeries := range series {
		batch := session.NewBatch(gocql.UnloggedBatch)
		for _, p := range oneSeries.Points {
			// http://stackoverflow.com/questions/35401344/passing-a-map-as-a-value-to-insert-into-cassandra
			batch.Query(insertIntStmt, oneSeries.Name, p.T, oneSeries.Tags, p.V)
		}
		err := session.ExecuteBatch(batch)
		if err != nil {
			// TODO: better error handling, we should have an error aggregator
			log.Warn(err)
		}
	}
	return nil
}

// WriteDoubleSeries implements Store interface
func (store Store) WriteDoubleSeries(series []common.DoubleSeries) error {
	// TODO: copied from write int series
	session := store.session
	for _, oneSeries := range series {
		batch := session.NewBatch(gocql.UnloggedBatch)
		for _, p := range oneSeries.Points {
			// http://stackoverflow.com/questions/35401344/passing-a-map-as-a-value-to-insert-into-cassandra
			batch.Query(insertDoubleStmt, oneSeries.Name, p.T, oneSeries.Tags, p.V)
		}
		err := session.ExecuteBatch(batch)
		if err != nil {
			// TODO: better error handling, we should have an error aggregator
			log.Warn(err)
		}
	}
	return nil
}

func (store Store) Shutdown() {
	log.Info("shutting down cassandra store, close connection")
	store.session.Close()
	log.Info("shutdown complete")
}
