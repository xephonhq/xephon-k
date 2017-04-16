package system

import (
	asst "github.com/stretchr/testify/assert"
	"github.com/xephonhq/xephon-k/pkg/util"
	"runtime"
	"testing"
)

// test CPU using fixture
func TestGlobalStat_Update(t *testing.T) {
	// NOTE: https://dave.cheney.net/2016/05/10/test-fixtures-in-go
	id := util.MockStringVar(&statPath, "fixtures/stat")
	defer util.RecoverMockedStringVar(id)
	assert := asst.New(t)
	stat := StatCollector{}
	err := stat.Update()
	assert.Nil(err)
	assert.Equal(runtime.NumCPU(), len(stat.CPUs))
}
