package disk

import (
	"testing"
	"github.com/xephonhq/xephon-k/pkg/storage/disk"
	"io/ioutil"
	"os"
	"bytes"
	//"encoding/binary"
	"fmt"
)

// writing series to disk without any compression and then read it out
type fileHeader struct {
	version          uint8
	timeCompression  uint8
	valueCompression uint8
}

// NOTE: must pass a pointer of buffer
func (header *fileHeader) write(buf *bytes.Buffer) {
	buf.WriteByte(header.version)
	buf.WriteByte(header.timeCompression)
	buf.WriteByte(header.valueCompression)
}

func TestNoCompress_Header(t *testing.T) {
	header := fileHeader{version: 1, timeCompression: disk.CompressionNone, valueCompression: disk.CompressionNone}
	//header := fileHeader{version: 1, timeCompression: disk.CompressionGzip, valueCompression: disk.CompressionZlib}
	tmpfile, err := ioutil.TempFile("", "xephon-no-compress")
	if err != nil {
		t.Fatal(err)
	}

	//defer os.Remove(tmpfile.Name())

	var buf bytes.Buffer
	// TODO: Endianness problem https://github.com/xephonhq/xephon-k/issues/34
	// but it seems for single uint8, this is not a problem
	//binary.Write(&buf, binary.LittleEndian, header.version)
	//binary.Write(&buf, binary.LittleEndian, header.timeCompression)
	//binary.Write(&buf, binary.LittleEndian, header.valueCompression)

	header.write(&buf)

	n, err := tmpfile.Write(buf.Bytes())
	t.Logf("written %d bytes\n", n)
	if err != nil {
		t.Fatal(err)
	}
	tmpfile.Close()

	// read stuff back
	f, err := os.Open(tmpfile.Name())
	readBuf := make([]byte, 3)
	n, err = f.Read(readBuf)
	t.Logf("read %d bytes\n", n)
	if err != nil {
		t.Fatal(err)
	}
	f.Close()
	// convert to header
	newHeader := fileHeader{}
	newHeader.version = uint8(readBuf[0])
	newHeader.timeCompression = uint8(readBuf[1])
	newHeader.valueCompression = uint8(readBuf[2])
	fmt.Printf("version %d, time compression %d, value compression %d\n",
		newHeader.version, newHeader.timeCompression, newHeader.valueCompression)
}
