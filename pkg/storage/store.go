package storage

import "github.com/xephonhq/xephon-k/pkg/common"

type Store interface {
	StoreType() string
	WriteIntSeries([]common.IntSeries) error
}
