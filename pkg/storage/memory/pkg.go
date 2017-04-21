package memory

import (
	"sync"

	"github.com/xephonhq/xephon-k/pkg/util"
)

var log = util.Logger.NewEntryWithPkg("k.s.mem")
var initSeriesCount = 10

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
