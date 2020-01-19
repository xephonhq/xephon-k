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
	tplBytes, err := ioutil.ReadFile("series_int.go")
	assert(err)
	tpl := string(tplBytes)
	// FIXME: raw series can not have methods like GetMaxMinTime
	// otherSeriesTypes := map[string]string{"RawSeries": "series_raw.go", "DoubleSeries": "series_double.go"}
	otherSeriesTypes := map[string]string{"DoubleSeries": "series_double.go"}

	for newType, newFile := range otherSeriesTypes {
		content := strings.Replace(tpl, "IntSeries", newType, -1)
		data := append([]byte("// Generated from series_int.go DO NOT EDIT!\n"), []byte(content)...)
		err = ioutil.WriteFile(newFile, data, 0644)
		assert(err)
	}

	fmt.Println("finished")

}
