package generator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstantGenerator_NextIntPoint(t *testing.T) {
	asst := assert.New(t)
	gen := NewConstantGenerator()
	p1 := gen.NextIntPoint()
	p2 := gen.NextIntPoint()
	asst.Equal(constantGeneratorDefaultTimeInterval, p2.TimeNano-p1.TimeNano)
}
