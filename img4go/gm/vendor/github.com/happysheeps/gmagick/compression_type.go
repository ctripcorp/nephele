package gmagick

/*
#include <wand/wand_api.h>
*/
import "C"

type CompressionType int

const (
	COMPRESSION_UNDEFINED     CompressionType = C.UndefinedCompression
	COMPRESSION_NO            CompressionType = C.NoCompression
	COMPRESSION_BZIP          CompressionType = C.BZipCompression
	COMPRESSION_FAX           CompressionType = C.FaxCompression
	COMPRESSION_GROUP4        CompressionType = C.Group4Compression
	COMPRESSION_JPEG          CompressionType = C.JPEGCompression
	COMPRESSION_JPEG2000      CompressionType = C.JPEG2000Compression
	COMPRESSION_LOSSLESS_JPEG CompressionType = C.LosslessJPEGCompression
	COMPRESSION_LZW           CompressionType = C.LZWCompression
	COMPRESSION_RLE           CompressionType = C.RLECompression
	COMPRESSION_ZIP           CompressionType = C.ZipCompression
	COMPRESSION_LZMA          CompressionType = C.LZMACompression
	COMPRESSION_JBIG1         CompressionType = C.JBIG1Compression
	COMPRESSION_JBIG2         CompressionType = C.JBIG2Compression
)
