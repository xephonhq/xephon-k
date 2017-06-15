package generator

import "testing"

func TestConstantGenerator_NextInt(t *testing.T) {
	c := Config{
		TimeInterval:    1,
		PointsPerSeries: 2,
		NumSeries:       3,
	}
	g := NewConstantGenerator(c)
	t.Log(g.NextInt())
	t.Log(g.NextInt())
	t.Log(g.NextInt())
	t.Log(g.NextInt())
	t.Log(g.NextInt())
}
