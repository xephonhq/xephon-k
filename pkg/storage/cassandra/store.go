package cassandra

import (
	"sync"

	"github.com/gocql/gocql"
	"github.com/xephonhq/xephon-k/pkg/common"
	"fmt"
)

var storeMap StoreMap

type StoreMap struct {
	mu     sync.RWMutex
	stores map[string]*Store
}

func init() {
	storeMap.stores = make(map[string]*Store, 1)
}

// GetDefaultCassandraStore will connect to cassandra if it is not found
// NOTE: we don't do it in init because it would break other stores, mem, mysql etc.
// TODO: we should return error and allow retry etc.
func GetDefaultCassandraStore() *Store {
	storeMap.mu.RLock()
	defer storeMap.mu.RUnlock()

	store, ok := storeMap.stores["default"]
	if ok {
		return store
	} else {
		log.Info("default cassandra store not found, connecting to cassandra now")
		storeMap.stores["default"] = NewCassandraStore()
		return storeMap.stores["default"]
	}
}

type Store struct {
	session *gocql.Session
}

func NewCassandraStore() *Store {
	store := &Store{}
	// connect to cassandra
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = naiveKeySpace
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("can't connect to cassandra %s", err)
		return store
	}
	store.session = session
	return store
}

// StoreType implements store interface
func (store Store) StoreType() string {
	return "cassandra"
}

// TODO: query

var writeStmtTmpl =
	`
	INSERT INTO %s.metrics (metric_name, metric_timestamp, tags, value) VALUES (?, ?, ?, ?)
	`

// WriteIntSeries implements Store interface
func (store Store) WriteIntSeries(series []common.IntSeries) error {
	// build the statement
	// TODO: let's not consider prepare for now
	// Use many goroutines when doing inserts, the driver is asynchronous but provides a synchronous API,
	// it can execute many queries concurrently
	session := store.session
	writeStmt := fmt.Sprintf(writeStmtTmpl, naiveKeySpace)
	for _, oneSeries := range series {
		batch := session.NewBatch(gocql.UnloggedBatch)
		for _, p := range oneSeries.Points {
			// TODO: can it handle map?
			// http://stackoverflow.com/questions/35401344/passing-a-map-as-a-value-to-insert-into-cassandra
			batch.Query(writeStmt, oneSeries.Name, p.TimeNano, oneSeries.Tags, p.V)
		}
		err := session.ExecuteBatch(batch)
		if err != nil {
			// TODO: better error handling
			log.Warn(err)
		}
	}
	return nil
}
