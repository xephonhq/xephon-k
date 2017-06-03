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
			return N, errors.Errorf("%s %v is marked as IntSeries but actually %s",
				series.GetName(), series.GetTags(), reflect.TypeOf(series))
		}
		for i := 0; i < len(intSeries.Points); i++ {
			tenc.WriteTime(intSeries.Points[i].T)
			venc.WriteInt(intSeries.Points[i].V)
		}
	case common.TypeDoubleSeries:
		doubleSeries, ok := series.(*common.DoubleSeries)
		if !ok {
			return N, errors.Errorf("%s %v is marked as DoubleSeries but actually %s",
				series.GetName(), series.GetTags(), reflect.TypeOf(series))
		}
		for i := 0; i < len(doubleSeries.Points); i++ {
			tenc.WriteTime(doubleSeries.Points[i].T)
			venc.WriteDouble(doubleSeries.Points[i].V)
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

func DecodeBlock(p []byte, meta common.SeriesMeta) (common.Series, error) {
	// read header
	// NOTE: currently we can only deal with time + value block, can't deal with time + value1 + value 2 ...
	timeBlockLength := binary.BigEndian.Uint32(p[:4])
	tBytes := p[4 : 4+timeBlockLength]
	vBytes := p[4+timeBlockLength:]
	// TODO: currently we only use raw binary decoder since it's our only encoding
	tdec := encoding.NewRawBinaryDecoder()
	vdec := encoding.NewRawBinaryDecoder()
	if err := tdec.Init(tBytes); err != nil {
		return nil, errors.Wrap(err, "can't initial time decoder")
	}
	if err := vdec.Init(vBytes); err != nil {
		return nil, errors.Wrap(err, "can't initial value decoder")
	}

	var s common.Series
	switch meta.GetSeriesType() {
	case common.TypeIntSeries:
		intSeries := common.NewIntSeries(meta.GetName())
		intSeries.SeriesMeta = meta
		// TODO: we can allocate the space directly if index entry record length
		for tdec.Next() && vdec.Next() {
			intSeries.Points = append(intSeries.Points,
				common.IntPoint{T: tdec.ReadTime(), V: vdec.ReadInt()})
		}
		s = intSeries
	case common.TypeDoubleSeries:
		doubleSeries := common.NewDoubleSeries(meta.GetName())
		doubleSeries.SeriesMeta = meta
		for tdec.Next() && vdec.Next() {
			doubleSeries.Points = append(doubleSeries.Points,
				common.DoublePoint{T: tdec.ReadTime(), V: vdec.ReadDouble()})
		}
		s = doubleSeries
	default:
		return nil, errors.Errorf("unsupported series type %s", common.SeriesTypeString(meta.GetSeriesType()))
	}
	return s, nil
}

//func DecodeBlockAsColunmar(p []byte, meta common.SeriesMeta) (common.SeriesColumnar, error) {
//
//}
