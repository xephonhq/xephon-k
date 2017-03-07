package common

import (
	"testing"

	"encoding/json"

	"github.com/stretchr/testify/assert"
)

func TestIntPoint_MarshalJSON(t *testing.T) {
	asst := assert.New(t)

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
