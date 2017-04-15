package util

import (
	asst "github.com/stretchr/testify/assert"
	"testing"
)

func TestMockVar(t *testing.T) {
	assert := asst.New(t)
	id := MockStringVar(&dummyVar, "ha")
	assert.Equal(dummyVar, "ha")
	RecoverMockedStringVar(id)
	assert.Equal("dummy", dummyVar)
}
