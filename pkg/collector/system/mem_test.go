package system

import (
	asst "github.com/stretchr/testify/assert"
	"github.com/xephonhq/xephon-k/pkg/util"
	"testing"
)

func TestMeminfoCollector_Update(t *testing.T) {
	id := util.MockStringVar(&meminfoPath, "fixtures/meminfo")
	defer util.RecoverMockedStringVar(id)
	assert := asst.New(t)
	mem := MeminfoCollector{}
	err := mem.Update()
	assert.Nil(err)
	assert.Equal(uint64(32858548), mem.MemTotal)
}
