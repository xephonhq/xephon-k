package generator

import (
	"github.com/xephonhq/xephon-k/pkg/common"
	"strconv"
	"time"
)

type ConstantGenerator struct {
	config        Config
	currentSeries int
	timestamp     int64
	interval      int64
}

func NewConstantGenerator(config Config) *ConstantGenerator {
	return &ConstantGenerator{
		config:        config,
		currentSeries: 0,
		// TODO: we only allow millisecond for now
		timestamp: time.Now().Unix() * 1000,
		interval:  int64(config.TimeInterval) * 1000,
	}
}

func (g *ConstantGenerator) NextInt() *common.IntSeries {
	s := common.NewIntSeries("xephon-bench")
	s.Tags["series-id"] = strconv.Itoa(g.currentSeries)
	s.Points = make([]common.IntPoint, g.config.PointsPerSeries)
	for i := 0; i < g.config.PointsPerSeries; i++ {
		s.Points[i].T = g.timestamp
		// TODO: allow update this constant
		s.Points[i].V = 1
		g.timestamp += g.interval
	}
	g.currentSeries++
	if g.currentSeries == g.config.NumSeries {
		g.currentSeries = 0
	}
	return s
}
