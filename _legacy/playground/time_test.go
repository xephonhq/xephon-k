package playground

import (
	"testing"
	"time"
	"math"
	"reflect"
	"github.com/xephonhq/xephon-k/pkg/bench/generator"
)

func TestTime_Unix(t *testing.T) {
	// https://github.com/xephonhq/xephon-k/issues/35
	// Different time stamp precision
	t.Log(math.MaxInt32)
	t.Log(time.Now().Unix())
	t.Log(time.Now().Unix() * 1000)
	t.Log(time.Now().UnixNano())
	t.Log(math.MaxInt64)
}

func TestTime_Type(t *testing.T) {
	t.Log(reflect.TypeOf(time.Millisecond))
}

func TestTime_Match(t *testing.T) {
	var tm time.Duration

	//tm = time.Second
	//tm = time.Millisecond
	tm = generator.DefaultOption.GetPrecision()

	switch tm {
	case time.Second:
		t.Log("second")
	case time.Millisecond:
		t.Log("millisecond")
	case time.Nanosecond:
		t.Log("nanosecond")
	}

	table := []struct {
		opt generator.Option
	}{
		{opt: generator.DefaultOption},
	}

	for _, tt := range table {
		t.Log(tt.opt.GetPrecision())
	}
}
