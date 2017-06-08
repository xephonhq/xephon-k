package memory

import (
	"sync"

	"github.com/pkg/errors"
	"github.com/xephonhq/xephon-k/pkg/util"
)

var log = util.Logger.NewEntryWithPkg("k.s.mem")
var initSeriesCount = 10

const (
	nameTagKey = "__name__"
)

// singleton map for share memory store between multiple go-kit services
var storeMap StoreMap

// StoreMap protects underlying mem store with a RWMutex
type StoreMap struct {
	mu     sync.RWMutex
	stores map[string]*Store
}

func init() {
	storeMap.stores = make(map[string]*Store, 1)
	storeMap.stores["default"] = NewMemStore()
}

// GetDefaultMemStore returns the default mem store initialized when package starts
func GetDefaultMemStore() *Store {
	storeMap.mu.RLock()
	defer storeMap.mu.RUnlock()
	return storeMap.stores["default"]
}

func CreateStore(config Config) error {
	storeMap.mu.Lock()
	defer storeMap.mu.Unlock()

	if err := config.Validate(); err != nil {
		return err
	}
	// TODO: rename to NewMemStore2 after old one is removed
	storeMap.stores["default"] = NewMemStore2(config)
	return nil
}

func GetStore() (*Store, error) {
	storeMap.mu.RLock()
	defer storeMap.mu.RUnlock()
	s, ok := storeMap.stores["default"]
	if !ok {
		return nil, errors.New("default store is not created! call CreateStore first")
	}
	return s, nil
}
