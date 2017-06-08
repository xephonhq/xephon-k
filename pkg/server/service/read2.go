package service

import (
	"github.com/xephonhq/xephon-k/pkg/common"
	"github.com/xephonhq/xephon-k/pkg/storage"
)

type ReadService2 struct {
	store storage.Store
}

func NewReadService(store storage.Store) *ReadService2 {
	return &ReadService2{
		store: store,
	}
}

func (r *ReadService2) QuerySeries(queries []common.Query) ([]common.QueryResult, []common.Series, error) {
	return r.store.QuerySeries(queries)
}
