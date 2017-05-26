package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func assert(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// NOTE: it should be called from upper folder
// Use `go run gen/main.go` instead of `go run main.go`
func main() {
	fmt.Println("generate other series based on int series")

	// int series is used as template
	tplBytes, err := ioutil.ReadFile("series_store_int.go")
	assert(err)
	tpl := string(tplBytes)
	otherSeriesTypes := map[string]map[string]string{"series_store_double.go": {"IntSeries": "DoubleSeries", "IntPoint": "DoublePoint"}}

	for newFile, replacements := range otherSeriesTypes {
		content := tpl
		for old, newStr := range replacements {
			content = strings.Replace(content, old, newStr, -1)
		}
		data := append([]byte("// Generated from series_store_int.go DO NOT EDIT!\n"), []byte(content)...)
		err = ioutil.WriteFile(newFile, data, 0644)
		assert(err)
	}

	fmt.Println("finished")

}
