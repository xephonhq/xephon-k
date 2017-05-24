package disk

import (
	asst "github.com/stretchr/testify/assert"
	"github.com/xephonhq/xephon-k/pkg/common"
	"github.com/xephonhq/xephon-k/pkg/util"
	"os"
	"testing"
)

func TestNewLocalFileIndexWriter(t *testing.T) {
	assert := asst.New(t)
	f := util.TempFile(t, "xephon")
	defer os.Remove(f.Name())

	w := NewLocalFileWriter(f, -1)
	assert.NotNil(w.Close())
}

func TestLocalFileWriter_WriteSeries(t *testing.T) {
	assert := asst.New(t)
	f := util.TempFile(t, "xephon")

	w := NewLocalFileWriter(f, -1)
	s := common.NewIntSeries("s")
	s.Tags = map[string]string{"os": "ubuntu", "machine": "machine-01"}
	s.Points = []common.IntPoint{{T: 1359788400000, V: 1}, {T: 1359788500000, V: 2}}
	assert.Nil(w.WriteSeries(s))
	w.Flush()

	f.Close()
	// FIXME: the file is always empty
}
