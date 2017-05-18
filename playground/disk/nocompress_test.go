package disk

import (
	"testing"
	"github.com/xephonhq/xephon-k/pkg/storage/disk"
	"io/ioutil"
	"os"
	"bytes"
	"encoding/binary"
)

// writing series to disk without any compression and then read it out
type fileHeader struct {
	version          uint8
	timeCompression  uint8
	valueCompression uint8
}

func TestNoCompress_Header(t *testing.T) {
	header := fileHeader{version: 1, timeCompression: disk.CompressionNone, valueCompression: disk.CompressionNone}
	tmpfile, err := ioutil.TempFile("", "xephon-no-compress")
	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(tmpfile.Name())

	var buf bytes.Buffer
	// TODO: binary endian
	// - influxdb writer seems to be using BigEndian
	// - old prometheus seems to be using LittleEndian
	// - new prometheus tsdb seems to be using BigEndian
	// lscpu | grep Endian
	// - shows little endian
	// Big endian byte ordering has been chosen as the "neutral" or standard for network data exchange and thus Big Endian byte ordering is also known as the "Network Byte Order"
	// http://www.sqlite.org/fileformat2.html
	// sqlite seems to be using big endian
	// https://docs.oracle.com/cd/E17275_01/html/programmer_reference/am_misc_faq.html
	// - BDB sort integer as byte strings, and it works badly for integer on little-endian architectures
	binary.Write(&buf, binary.LittleEndian, header.version)
	binary.Write(&buf, binary.LittleEndian, header.timeCompression)
	binary.Write(&buf, binary.BigEndian, header.valueCompression)
	n, err := tmpfile.Write(buf.Bytes())
	t.Logf("written %d bytes", n)
	if err != nil {
		t.Fatal(err)
	}
	tmpfile.Close()

	// read stuff back
	// f, err := os.Open(tmpfile.Name())

}
