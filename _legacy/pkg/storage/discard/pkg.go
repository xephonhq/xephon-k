package discard

import (
	"github.com/xephonhq/xephon-k/pkg/common"
	"github.com/xephonhq/xephon-k/pkg/util"
)

var log = util.Logger.NewEntryWithPkg("k.storage.discard")

type Store struct {
}

func NewDiscardStore() *Store {
	return &Store{}
}

func (Store) StoreType() string {
	return "discard"
}

func (Store) QuerySeries([]common.Query) ([]common.QueryResult, []common.Series, error) {
	result := make([]common.QueryResult, 0)
	series := make([]common.Series, 0)
	return result, series, nil
}

func (Store) WriteIntSeries([]common.IntSeries) error {
	return nil
}

func (Store) WriteDoubleSeries([]common.DoubleSeries) error {
	return nil
}

func (Store) Shutdown() {
	log.Info("shutting down discard store, nothing to do")
}
