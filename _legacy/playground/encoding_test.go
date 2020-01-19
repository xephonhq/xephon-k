package playground

import (
	"testing"
	"encoding/json"
	"os"
)

// https://golang.org/pkg/encoding/json/#Encoder.Encode
func TestEncoding_JSON(t *testing.T) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(1)
	encoder.Encode(map[string]string{"foo": "bar"})
}
