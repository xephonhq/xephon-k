package disk

import (
	"encoding/binary"
	"io"
	"reflect"

	"github.com/pkg/errors"
	"github.com/xephonhq/xephon-k/pkg/common"
	"github.com/xephonhq/xephon-k/pkg/encoding"
)

// TODO: allow adaptive encoding
type EncodingOption struct {
	TimeCodec        byte
	IntValueCodec    byte
	DoubleValueCodec byte
}

// https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
func NewEncodingOption(options ...func(*EncodingOption)) (EncodingOption, error) {
	opt := EncodingOption{
		TimeCodec:        encoding.CodecRawBigEndian,
		IntValueCodec:    encoding.CodecVarInt,
		DoubleValueCodec: encoding.CodecVarInt,
	}
	for _, option := range options {
		option(&opt)
	}
	// valid if all the codec we use are already registered
	// TODO: also check if this codec supports encoding time, int, double value etc.
	if !encoding.IsRegisteredCodec(opt.TimeCodec) {
		return opt, errors.Errorf("codec %v is not registered, can't use it for time", opt.TimeCodec)
	}
	if !encoding.IsRegisteredCodec(opt.IntValueCodec) {
		return opt, errors.Errorf("codec %v is not registered, can't use it for int value", opt.IntValueCodec)
	}
	if !encoding.IsRegisteredCodec(opt.DoubleValueCodec) {
		return opt, errors.Errorf("codec %v is not registered, can't use it for double value", opt.DoubleValueCodec)
	}
	return opt, nil
}

// EncodeBlockTo encode the series data points and write to underlying writer
// It does not return bytes to avoid need less copying when concat encoded time and values
func EncodeBlockTo(series common.Series, opt EncodingOption, w io.Writer) (int, error) {
	N := 0
	n := 0
	var (
		vCodec         byte
		tenc           encoding.TimeEncoder
		venc           encoding.ValueEncoder
		tBytes, vBytes []byte
		err            error
	)
	blockHeader := make([]byte, 4)

	// encode time and value separately
	// TODO: should put this logic in the encoding package, like have a object called codec
	switch opt.TimeCodec {
	case encoding.CodecRawBigEndian:
		tenc = encoding.NewBigEndianBinaryEncoder()
	case encoding.CodecRawLittleEndian:
		tenc = encoding.NewLittleEndianBinaryEncoder()
	case encoding.CodecVarInt:
		tenc = encoding.NewVarIntEncoder()
	default:
		return 0, errors.Errorf("unsupported codec %s for time encoder", encoding.CodecString(opt.TimeCodec))
	}

	// determine which kind of value are we encoding and read the correspond config
	switch series.GetSeriesType() {
	case common.TypeIntSeries:
		vCodec = opt.IntValueCodec
	case common.TypeDoubleSeries:
		vCodec = opt.DoubleValueCodec
	default:
		return 0, errors.Errorf("unsupported series type %s, no available codec in option",
			common.SeriesTypeString(series.GetSeriesType()))
	}

	switch vCodec {
	case encoding.CodecRawBigEndian:
		venc = encoding.NewBigEndianBinaryEncoder()
	case encoding.CodecRawLittleEndian:
		venc = encoding.NewLittleEndianBinaryEncoder()
	case encoding.CodecVarInt:
		venc = encoding.NewVarIntEncoder()
	default:
		return 0, errors.Errorf("unsupported codec %s for value encoder", encoding.CodecString(opt.TimeCodec))
	}

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

	var (
		s    common.Series
		tdec encoding.TimeDecoder
		vdec encoding.ValueDecoder
	)
	switch tBytes[0] {
	case encoding.CodecRawBigEndian, encoding.CodecRawLittleEndian:
		tdec = encoding.NewRawBinaryDecoder()
	case encoding.CodecVarInt:
		tdec = encoding.NewVarIntDecoder()
	default:
		return nil, errors.Wrapf(encoding.ErrCodecNotSupported, "unknown codec %s", encoding.CodecString(tBytes[0]))
	}
	switch vBytes[0] {
	case encoding.CodecRawBigEndian, encoding.CodecRawLittleEndian:
		vdec = encoding.NewRawBinaryDecoder()
	case encoding.CodecVarInt:
		vdec = encoding.NewVarIntDecoder()
	default:
		return nil, errors.Wrapf(encoding.ErrCodecNotSupported, "unknown codec %s", encoding.CodecString(vBytes[0]))
	}

	if err := tdec.Init(tBytes); err != nil {
		return nil, errors.Wrap(err, "can't initial time decoder")
	}
	if err := vdec.Init(vBytes); err != nil {
		return nil, errors.Wrap(err, "can't initial value decoder")
	}

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
