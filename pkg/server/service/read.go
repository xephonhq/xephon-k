package service

import (
	"github.com/xephonhq/xephon-k/pkg/common"
	"github.com/xephonhq/xephon-k/pkg/storage"
)

type ReadService struct {
	store storage.Store
}

func NewReadService(store storage.Store) *ReadService {
	return &ReadService{
		store: store,
	}
}

func (r *ReadService) QuerySeries(queries []common.Query) ([]common.QueryResult, []common.Series, error) {
	return r.store.QuerySeries(queries)
}
