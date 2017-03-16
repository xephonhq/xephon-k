package common

import (
	"testing"

	"encoding/json"

	"github.com/stretchr/testify/assert"
)

func TestQuery_JSON(t *testing.T) {
	asst := assert.New(t)
	tags := make(map[string]string)
	tags["os"] = "ubuntu"
	q := Query{Tags: tags, Name: "cpu.idle", MatchPolicy: "exact"}
	j, err := json.Marshal(q)
	asst.Nil(err)
	asst.Equal(`{"name":"cpu.idle","tags":{"os":"ubuntu"},"match_policy":"exact"}`, string(j))
	qr := QueryResult{Query: q, Matched: 1}
	j, err = json.Marshal(qr)
	asst.Nil(err)
	asst.Equal(`{"name":"cpu.idle","tags":{"os":"ubuntu"},"match_policy":"exact","matched":1}`, string(j))
}

func TestQuery_Hash(t *testing.T) {
	asst := assert.New(t)

	tags1 := make(map[string]string)
	tags1["os"] = "ubuntu"
	tags1["region"] = "us"
	tags2 := make(map[string]string)
	tags2["os"] = "ubuntu"
	tags2["region"] = "us"

	q := Query{Tags: tags1, Name: "cpu.idle", MatchPolicy: "exact"}
	s := IntSeries{Tags: tags2, Name: "cpu.idle"}
	asst.Equal(q.Hash(), s.Hash())
}
