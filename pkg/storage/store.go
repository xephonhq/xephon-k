package storage

import (
	"github.com/xephonhq/xephon-k/pkg/common"
	"github.com/xephonhq/xephon-k/pkg/storage/cassandra"
	"github.com/xephonhq/xephon-k/pkg/storage/discard"
	"github.com/xephonhq/xephon-k/pkg/storage/disk"
	"github.com/xephonhq/xephon-k/pkg/storage/memory"
)

// check interface
var _ Store = (*discard.Store)(nil)
var _ Store = (*memory.Store)(nil)
var _ Store = (*disk.Store)(nil)
var _ Store = (*cassandra.Store)(nil)

// Store is the base interface for all type of storages
// TODO: each store should maintains some counter for internal metrics
type Store interface {
	StoreType() string
	QuerySeries([]common.Query) ([]common.QueryResult, []common.Series, error)
	WriteIntSeries([]common.IntSeries) error
	WriteDoubleSeries([]common.DoubleSeries) error
	Shutdown()
}

func CreateStore(engine string, config Config) (Store, error) {
	switch engine {
	case "discard", "null":
		return discard.NewDiscardStore(), nil
	case "memory":
		return memory.CreateStore(config.Memory)
	case "disk":
		return disk.CreateStore(config.Disk)
	case "cassandra":
		return cassandra.CreateStore(config.Cassandra)
	default:
		log.Fatalf("unknown storage %s", engine)
		return nil, nil
	}
}
