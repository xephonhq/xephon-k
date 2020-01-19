package playground

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"io"
	"os"
	"testing"
	//"time"
)

// https://golang.org/pkg/compress/gzip/#example__writerReader
func TestGzip(t *testing.T) {
	var buf bytes.Buffer

	data := "I am going to duplicate I am I am I am I am I am\n"

	t.Logf("data length %d\n", len(data))

	zw := gzip.NewWriter(&buf)
	//zw.Name = "a-new-hope.txt"
	//zw.Comment = "an epic space opera by George Lucas"
	//zw.ModTime = time.Date(1977, time.May, 25, 0, 0, 0, 0, time.UTC)

	_, err := zw.Write([]byte(data))
	if err != nil {
		t.Fatal(err)
	}
	if err := zw.Close(); err != nil {
		t.Fatal(err)
	}
	// the size actually increased since the data is too small
	t.Logf("after gzip %d\n", buf.Len())

	zr, err := gzip.NewReader(&buf)
	if err != nil {
		t.Fatal(err)
	}

	// NOTE: output comes from here
	if _, err := io.Copy(os.Stdout, zr); err != nil {
		t.Fatal(err)
	}

	if err := zr.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestZlib(t *testing.T) {
	var buf bytes.Buffer

	data := "I am going to duplicate I am I am I am I am I am\n"

	t.Logf("data length %d\n", len(data))

	zw := zlib.NewWriter(&buf)

	_, err := zw.Write([]byte(data))
	if err != nil {
		t.Fatal(err)
	}
	if err := zw.Close(); err != nil {
		t.Fatal(err)
	}
	// the size actually increased since the data is too small
	t.Logf("after zlib %d\n", buf.Len())

	zr, err := zlib.NewReader(&buf)
	if err != nil {
		t.Fatal(err)
	}

	// NOTE: output comes from here
	if _, err := io.Copy(os.Stdout, zr); err != nil {
		t.Fatal(err)
	}

	if err := zr.Close(); err != nil {
		t.Fatal(err)
	}
}