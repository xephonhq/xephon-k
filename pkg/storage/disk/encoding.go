package disk

import (
	"encoding/binary"
	"io"
	"reflect"

	"github.com/pkg/errors"
	"github.com/xephonhq/xephon-k/pkg/common"
	"github.com/xephonhq/xephon-k/pkg/encoding"
)

// EncodeBlockTo encode the series data points and write to underlying writer
// It does not return bytes to avoid need less copying when concat encoded time and values
func EncodeBlockTo(series common.Series, w io.Writer) (int, error) {
	N := 0
	n := 0
	var (
		tenc           encoding.TimeEncoder
		venc           encoding.ValueEncoder
		tBytes, vBytes []byte
		err            error
	)
	blockHeader := make([]byte, 4)

	// encode time and value separately
	// TODO: only use RawBigEndianTime/IntEncoder for now, should allow option or adaptive
	tenc = encoding.NewBigEndianBinaryEncoder()
	venc = encoding.NewBigEndianBinaryEncoder()

	// TODO: deal with columnar format
	switch series.GetSeriesType() {
	case common.TypeIntSeries:
		intSeries, ok := series.(*common.IntSeries)
		if !ok {
			return N, errors.Errorf("%s %v is marked as int but actually %s",
				series.GetName(), series.GetTags(), reflect.TypeOf(series))
		}
		for i := 0; i < len(intSeries.Points); i++ {
			tenc.WriteTime(intSeries.Points[i].T)
			venc.WriteInt(intSeries.Points[i].V)
		}
	default:
		return N, errors.Errorf("unsupported series type %d", series.GetSeriesType())
	}
	// NOTE: the encoder write encoding information at start of each block
	if tBytes, err = tenc.Bytes(); err != nil {
		return N, errors.Wrap(err, "can't get encoded time as bytes")
	}
	if vBytes, err = venc.Bytes(); err != nil {
		return N, errors.Wrap(err, "can't get encoded value as bytes")
	}

	// write block header
	binary.BigEndian.PutUint32(blockHeader, uint32(len(tBytes)))
	if n, err = w.Write(blockHeader); err != nil {
		return N, errors.Wrap(err, "can't write block header to buffer")
	}
	N += n
	// write encoded time and values, the encoding is in the bytes already, we don't need to prefix them
	if n, err = w.Write(tBytes); err != nil {
		return N, errors.Wrap(err, "cant write encoded time to buffer")
	}
	N += n
	if n, err = w.Write(vBytes); err != nil {
		return N, errors.Wrap(err, "can't write encoded value to buffer")
	}
	N += n
	return N, nil
}
