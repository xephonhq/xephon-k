package xephonk

import (
	"testing"

	"github.com/xephonhq/xephon-k/pkg/util"
)

func TestSerializer_WriteInt(t *testing.T) {
	xks := NewSerializer()
	xks.WriteInt(util.CreateDummyIntPoints())
	xks.WriteInt(util.CreateDummyIntPoints())
	xks.End()
	log.Info(string(xks.Data()))
}
