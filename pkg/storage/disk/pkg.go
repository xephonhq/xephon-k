package disk

const (
	CompressionNone = iota
	CompressionZip
)

// https://golang.org/pkg/compress/
// - https://github.com/klauspost/compress a drop in replace that claims to be faster
// https://github.com/golang/snappy

// uncompressed size (saw it from archive/zip)
// if the compressed data can be traverse

// NOTE: saw this when browse prometheus's code
// LSB means Least Significant Bits first, as used in the GIF file format.
// LSB Order = iota
// MSB means Most Significant Bits first, as used in the TIFF and PDF
// file formats.
// MSB
