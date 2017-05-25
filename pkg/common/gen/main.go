package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func assert(err error) {
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}

func main() {
	fmt.Println("generate other series based on int series")

	// int series is used as template
	tplBytes, err := ioutil.ReadFile("../series_int.go")
	assert(err)
	tpl := string(tplBytes)
	otherSeriesTypes := map[string]string{"RawSeries": "../series_raw.go", "DoubleSeries": "../series_double.go"}

	for newType, newFile := range otherSeriesTypes {
		content := strings.Replace(tpl, "IntSeries", newType, -1)
		data := append([]byte("// Generated from series_int.go DO NOT EDIT!\n"), []byte(content)...)
		err = ioutil.WriteFile(newFile, data, 0644)
		assert(err)
	}

	fmt.Println("finished")

}
