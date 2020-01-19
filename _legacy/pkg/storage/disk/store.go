package disk

import (
	"github.com/pkg/errors"
	"github.com/xephonhq/xephon-k/pkg/common"
	"github.com/xephonhq/xephon-k/pkg/encoding"
	"io/ioutil"
	"sync"
)

type Store struct {
	mu     sync.RWMutex
	config Config
	writer DataFileWriter
}

func init() {
	storeMap.stores = make(map[string]*Store, 1)
}

func NewDiskStore(config Config) (*Store, error) {
	root := config.Folder
	if root == "" {
		log.Warn("root is empty, using tmp folder")
		root = "/tmp"
	}
	store := &Store{config: config}
	// FIXME: we should not use temp file, but since we don't read from disk for now, we don't scan the directory
	f, err := ioutil.TempFile(root, "xephonk-data")
	if err != nil {
		return nil, errors.Wrap(err, "can't create file!")
	}

	// config encoding
	// TODO: do we still need to check, we have validate in
	log.Infof("encoding: %s", config.Encoding)
	timeCodec, _ := encoding.Str2Codec(config.Encoding["time"])
	intCodec, _ := encoding.Str2Codec(config.Encoding["int"])
	doubleCodec, _ := encoding.Str2Codec(config.Encoding["double"])

	opt, err := NewEncodingOption(func(o *EncodingOption) {
		o.TimeCodec = timeCodec
		o.IntValueCodec = intCodec
		o.DoubleValueCodec = doubleCodec
	})
	if err != nil {
		return nil, errors.Wrap(err, "can't set encoding options")
	}
	w, err := NewLocalFileWriter(f, config.FileBufferSize, opt)
	if err != nil {
		return nil, errors.Wrap(err, "can't create local file writer")
	}
	store.writer = w
	return store, nil
}

func (store *Store) StoreType() string {
	return "disk"
}

func (store *Store) QuerySeries(queries []common.Query) ([]common.QueryResult, []common.Series, error) {
	log.Panic("DiskStore does not implement QuerySeries")
	return nil, nil, nil
}

func (store *Store) WriteIntSeries(series []common.IntSeries) error {
	store.mu.Lock()
	defer store.mu.Unlock()

	for i := 0; i < len(series); i++ {
		err := store.writer.WriteSeries(&series[i])
		if err != nil {
			return errors.Wrapf(err, "write data failed for %s %v", series[i].Name, series[i].Tags)
		}
	}

	return nil
}

func (store *Store) WriteDoubleSeries(series []common.DoubleSeries) error {
	store.mu.Lock()
	defer store.mu.Unlock()

	for i := 0; i < len(series); i++ {
		err := store.writer.WriteSeries(&series[i])
		if err != nil {
			return errors.Wrapf(err, "write data failed for %s %v", series[i].Name, series[i].Tags)
		}
	}

	return nil
}

func (store *Store) Shutdown() {
	log.Info("shutting down on disk store")
	if err := store.writer.WriteIndex(); err != nil {
		log.Warn("can't write index")
		log.Warn(err)
	}
	if err := store.writer.Close(); err != nil {
		log.Warn("can't close writer")
		log.Warn(err)
	}
	log.Info("shutdown complete")
}
