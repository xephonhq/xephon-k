package playground

import (
	"testing"
	"time"
)

func TestTime_Unix(t *testing.T) {
	t.Log(time.Now().Unix())
	t.Log(time.Now().Unix() * 1000)
}