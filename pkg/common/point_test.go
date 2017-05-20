package common

import (
	"testing"

	"encoding/json"

	"github.com/stretchr/testify/assert"
	"sort"
)

func TestIntPoint_MarshalJSON(t *testing.T) {
	asst := assert.New(t)

	//1492565887423026057 UnixNano
	//1492566023000       Unix() * 1000
	//1359788400000
	p := IntPoint{TimeNano: 1359788400000, V: 1}
	// http://stackoverflow.com/questions/21390979/custom-marshaljson-never-gets-called-in-go
	// j, err := json.Marshal(p)
	// TODO: what happens when i use decoder instead of json.Marshal
	j, err := json.Marshal(&p)
	asst.Nil(err)
	asst.Equal("[1359788400000,1]", string(j))
}

func TestIntPoint_UnmarshalJSON(t *testing.T) {
	asst := assert.New(t)

	p := IntPoint{TimeNano: 1359788400000, V: 1}
	j, err := json.Marshal(&p)
	asst.Nil(err)
	var p2 IntPoint
	err = json.Unmarshal(j, &p2)
	asst.Nil(err)
	asst.Equal(p, p2)
}

func TestByTime_Less(t *testing.T) {
	asst := assert.New(t)
	p1 := IntPoint{TimeNano: 1359788400000, V: 1}
	p2 := IntPoint{TimeNano: 1359788401000, V: 1}
	p3 := IntPoint{TimeNano: 1359788400200, V: 1}
	p4 := IntPoint{TimeNano: 1459788400000, V: 1}
	p := []IntPoint{p2, p1, p4, p3}
	sort.Sort(ByTime(p))
	asst.Equal(int64(1359788400000), p[0].TimeNano)
	asst.Equal(int64(1459788400000), p[3].TimeNano)
}

// 5000000	       318 ns/op
func BenchmarkIntPoint_MarshalJSON(b *testing.B) {
	p := IntPoint{TimeNano: 1359788400000, V: 1}
	var err error
	for i := 0; i < b.N; i++ {
		_, err = p.MarshalJSON()
	}
	if err != nil {
		err.Error()
	}
}

// 10000000	       225 ns/op
func BenchmarkIntPoint_MarshalJSON2(b *testing.B) {
	p := IntPoint{TimeNano: 1359788400000, V: 1}
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
	p := IntPoint{TimeNano: 1359788400000, V: 1}
	var err error
	for i := 0; i < b.N; i++ {
		_, err = p.MarshalJSON3()
	}
	if err != nil {
		err.Error()
	}
}
