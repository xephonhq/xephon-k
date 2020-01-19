package cassandra

import (
	"github.com/xephonhq/xephon-k/pkg/util"
	"sync"
)

var log = util.Logger.NewEntryWithPkg("k.storage.cassandra")

const defaultKeySpace = "xephon"

var storeMap StoreMap

// StoreMap is used to allow multiple cassandra session, it is also used as a singleton when you are just using the default store.
// its methods use a RWMutex
type StoreMap struct {
	mu     sync.RWMutex
	stores map[string]*Store
}

func init() {
	storeMap.stores = make(map[string]*Store, 1)
}

func CreateStore(config Config) (*Store, error) {
	storeMap.mu.Lock()
	defer storeMap.mu.Unlock()
	if err := config.Validate(); err != nil {
		return nil, err
	}
	store, err := NewCassandraStore(config)
	if err != nil {
		return nil, err
	}
	storeMap.stores["default"] = store
	return store, nil
}
