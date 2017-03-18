package storage

import "github.com/xephonhq/xephon-k/pkg/common"

// Store is the base interface for all type of storages
type Store interface {
	StoreType() string
	QueryIntSeries(common.Query) ([]common.IntSeries, error)
	WriteIntSeries([]common.IntSeries) error
	// TODO: maybe we should add a graceful shutdown method
}
