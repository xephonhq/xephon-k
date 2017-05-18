package playground

import (
	"testing"
	"time"
	"math"
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