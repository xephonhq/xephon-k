package storage

import (
	"github.com/xephonhq/xephon-k/pkg/common"
	"github.com/xephonhq/xephon-k/pkg/storage/cassandra"
	"github.com/xephonhq/xephon-k/pkg/storage/memory"
)

// check interface
var _ Store = (*memory.Store)(nil)
var _ Store = (*cassandra.Store)(nil)

// Store is the base interface for all type of storages
type Store interface {
	StoreType() string
	// TODO: support double
	QueryIntSeriesBatch([]common.Query) ([]common.QueryResult, []common.IntSeries, error)
	// Deprecated: Use QueryIntSeriesBatch instead
	QueryIntSeries(common.Query) ([]common.IntSeries, error)
	WriteIntSeries([]common.IntSeries) error
	Shutdown()
}
