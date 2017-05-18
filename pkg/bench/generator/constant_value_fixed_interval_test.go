package generator

import (
	asst "github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var fixedTimeInterval = []struct {
	o        Option // generator option
	expected int64  // expected time interval
}{
	{o: defaultOption, expected: time.Millisecond.Nanoseconds() / 1000000},
	{o: defaultNanosecondOption, expected: int64(time.Nanosecond)},
}

func TestConstantValueFixedInterval_NextIntPoint(t *testing.T) {
	assert := asst.New(t)
	// table driven test https://dave.cheney.net/2013/06/09/writing-table-driven-tests-in-go
	for _, tt := range fixedTimeInterval {
		gen := NewConstantValueFixedInterval(tt.o)
		p1 := gen.NextIntPoint()
		p2 := gen.NextIntPoint()
		assert.Equal(tt.expected, p2.TimeNano-p1.TimeNano)
	}
}

func TestConstantValueFixedInterval_NextDoublePoint(t *testing.T) {
	assert := asst.New(t)
	gen := NewConstantValueFixedInterval(defaultOption)
	p1 := gen.NextDoublePoint()
	p2 := gen.NextDoublePoint()
	assert.Equal(defaultOption.GetInterval().Nanoseconds()/1000000, p2.TimeNano-p1.TimeNano)
}
