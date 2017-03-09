package common

type IntSeries struct {
	Name   string            `json:"name"`
	Tags   map[string]string `json:"tags"`
	Points []IntPoint        `json:"points"`
}

