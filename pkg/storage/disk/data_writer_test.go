package disk

import (
	"os"
	"testing"

	asst "github.com/stretchr/testify/assert"
	"github.com/xephonhq/xephon-k/pkg/common"
	"github.com/xephonhq/xephon-k/pkg/util"
)

func TestNewLocalFileIndexWriter(t *testing.T) {
	assert := asst.New(t)
	f := util.TempFile(t, "xephon")
	defer os.Remove(f.Name())

	w, err := NewLocalFileWriter(f, -1)
	assert.Nil(err)
	assert.NotNil(w.Close())
}

func TestLocalFileWriter_WriteSeries(t *testing.T) {
	assert := asst.New(t)
	f := util.TempFile(t, "xephon")
	defer os.Remove(f.Name())

	// f, err = os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
	w, err := NewLocalFileWriter(f, -1)
	assert.Nil(err)

	s := common.NewIntSeries("s")
	s.Tags = map[string]string{"os": "ubuntu", "machine": "machine-01"}
	s.Points = []common.IntPoint{{T: 1359788400000, V: 1}, {T: 1359788500000, V: 2}}

	assert.Nil(w.WriteSeries(s))
	assert.Equal(uint64(9+2+16*2), w.n)
	assert.Equal(ErrNotFinalized, w.Close())
	assert.Nil(w.WriteIndex())
	assert.Nil(w.Close())

	// NOTE: need to re-open the file because the writer has already closed it
	f, err = os.OpenFile(f.Name(), os.O_RDONLY, 0666)
	assert.Nil(err)
	r, err := NewLocalDataFileReader(f)
	assert.Nil(err)
	assert.Nil(r.ReadIndexOfIndexes())
	assert.Equal(1, r.SeriesCount())
	assert.Nil(r.ReadAllIndexEntries())
	assert.Nil(r.Close())
}
