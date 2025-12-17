package core

import "github.com/superiden3/go_compress/internal/core/algorithms"

type ErrUnsupportedAlgorithmType struct {
	Algorithm string
}

func (e *ErrUnsupportedAlgorithmType) Error() string {
	return "unsupported compression algorithm: " + e.Algorithm
}

type Compressor interface {
	Compress(data []byte) ([]byte, error)
	Decompress(data []byte) ([]byte, error)
}

func NewCompressor(algorithm string) (Compressor, error) {
	switch algorithm {
		case "gzip":
			return algorithms.NewGzipCompressor(), nil
		case "zlib":
			return algorithms.NewZlibCompressor(), nil
		case "lz4":
			return algorithms.NewLz4Compressor(), nil
		default:
			return nil, &ErrUnsupportedAlgorithmType{Algorithm: algorithm}
	}
}