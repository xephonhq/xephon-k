package xephonk

import (
	"github.com/xephonhq/xephon-k/pkg/util"
	"testing"
)

func TestSerializer_WriteInt(t *testing.T) {
	xks := Serializer{}
	xks.Start()
	xks.WriteInt(util.CreateDummyIntPoints())
	xks.WriteInt(util.CreateDummyIntPoints())
	xks.End()
	//log.Info(string(xks.Data()))
}
