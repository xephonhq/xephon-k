package generator

import (
	"testing"
	"time"

	asst "github.com/stretchr/testify/assert"
)

var fixedTimeInterval = []struct {
	o        Option // generator option
	expected int64  // expected time interval
}{
	{o: DefaultOption, expected: time.Millisecond.Nanoseconds() / 1000000},
	// FIXME: why using the default Option get 0s .... because init is not called? nop
	// Found it https://coderwall.com/p/g6vuqq/go-s-init-function
	// package variable is executed before init function, so fixedTimeInterval got an empty option
	// {o: DefaultNanosecondOption, expected: int64(time.Nanosecond)},
	 {o: Option{startTime: time.Now(), interval: time.Millisecond, precision: time.Millisecond}, expected: int64(time.Nanosecond)},
	{o: NewDefaultOption(), expected: time.Millisecond.Nanoseconds() / 1000000},
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
	gen := NewConstantValueFixedInterval(DefaultOption)
	p1 := gen.NextDoublePoint()
	p2 := gen.NextDoublePoint()
	assert.Equal(DefaultOption.GetInterval().Nanoseconds()/1000000, p2.TimeNano-p1.TimeNano)
}
