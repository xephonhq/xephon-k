package generator

import (
	"testing"
	"time"

	asst "github.com/stretchr/testify/assert"
)

var fixedTimeInterval = []struct {
	opt      Option // generator option
	expected int64  // expected time interval
}{
	{opt: DefaultSecondOption, expected: time.Second.Nanoseconds() / 1000000000},
	{opt: DefaultOption, expected: time.Millisecond.Nanoseconds() / 1000000},
	{opt: DefaultNanosecondOption, expected: int64(time.Nanosecond)},
}

func TestConstantValueFixedInterval_NextIntPoint(t *testing.T) {
	assert := asst.New(t)
	// table driven test https://dave.cheney.net/2013/06/09/writing-table-driven-tests-in-go
	for _, tt := range fixedTimeInterval {
		gen := NewConstantValueFixedInterval(tt.opt)
		p1 := gen.NextIntPoint()
		p2 := gen.NextIntPoint()
		assert.Equal(tt.expected, p2.TimeNano-p1.TimeNano)
	}
}

func TestConstantValueFixedInterval_NextDoublePoint(t *testing.T) {
	assert := asst.New(t)
	for _, tt := range fixedTimeInterval {
		gen := NewConstantValueFixedInterval(tt.opt)
		p1 := gen.NextDoublePoint()
		p2 := gen.NextDoublePoint()
		assert.Equal(tt.expected, p2.TimeNano-p1.TimeNano)
	}
}
