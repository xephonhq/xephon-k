package common

import (
	"testing"

	"encoding/json"

	asst "github.com/stretchr/testify/assert"
	"sort"
)

func TestIntPoint_MarshalJSON(t *testing.T) {
	assert := asst.New(t)

	//1492565887423026057 UnixNano
	//1492566023000       Unix() * 1000
	//1359788400000
	p := IntPoint{T: 1359788400000, V: 1}
	// http://stackoverflow.com/questions/21390979/custom-marshaljson-never-gets-called-in-go
	// j, err := json.Marshal(p)
	// TODO: what happens when i use decoder instead of json.Marshal
	j, err := json.Marshal(&p)
	assert.Nil(err)
	assert.Equal("[1359788400000,1]", string(j))
}

func TestIntPoint_UnmarshalJSON(t *testing.T) {
	assert := asst.New(t)

	p := IntPoint{T: 1359788400000, V: 1}
	j, err := json.Marshal(&p)
	assert.Nil(err)
	var p2 IntPoint
	err = json.Unmarshal(j, &p2)
	assert.Nil(err)
	assert.Equal(p, p2)
}

func TestDoublePoint_MarshalJSON(t *testing.T) {
	assert := asst.New(t)

	p := DoublePoint{T: 1359788400000, V: 1.2}
	//p := DoublePoint{T: 1359788400000, V: 1} // it's still 1.000000
	j, err := json.Marshal(&p)
	assert.Nil(err)
	assert.Equal("[1359788400000,1.200000]", string(j))
}

func TestDoublePoint_UnmarshalJSON(t *testing.T) {
	assert := asst.New(t)

	p := DoublePoint{T: 1359788400000, V: 1.2}
	j, err := json.Marshal(&p)
	assert.Nil(err)
	var p2 DoublePoint
	err = json.Unmarshal(j, &p2)
	assert.Equal(p, p2)
}

func TestIntPoints_Less(t *testing.T) {
	assert := asst.New(t)
	p1 := IntPoint{T: 1359788400000, V: 1}
	p2 := IntPoint{T: 1359788401000, V: 1}
	p3 := IntPoint{T: 1359788400200, V: 1}
	p4 := IntPoint{T: 1459788400000, V: 1}
	p := []IntPoint{p2, p1, p4, p3}
	sort.Slice(p, func(i, j int) bool {
		return p[i].T < p[j].T
	})
	assert.Equal(int64(1359788400000), p[0].T)
	assert.Equal(int64(1459788400000), p[3].T)
}

// 10000000	       222 ns/op
func BenchmarkIntPoint_MarshalJSON(b *testing.B) {
	p := IntPoint{T: 1359788400000, V: 1}
	var err error
	for i := 0; i < b.N; i++ {
		_, err = p.MarshalJSON()
	}
	if err != nil {
		err.Error()
	}
}

// 5000000	       306 ns/op
func BenchmarkIntPoint_MarshalJSON2(b *testing.B) {
	p := IntPoint{T: 1359788400000, V: 1}
	var err error
	for i := 0; i < b.N; i++ {
		_, err = p.MarshalJSON2()
	}
	if err != nil {
		err.Error()
	}
}

// 10000000	       139 ns/op
func BenchmarkIntPoint_MarshalJSON3(b *testing.B) {
	p := IntPoint{T: 1359788400000, V: 1}
	var err error
	for i := 0; i < b.N; i++ {
		_, err = p.MarshalJSON3()
	}
	if err != nil {
		err.Error()
	}
}
