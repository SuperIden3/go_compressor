package core

import "github.com/superiden3/go_compress/internal/core/algorithms"

type ErrUnsupportedAlgorithmType struct { // Custom error for unsupported algorithms
	Algorithm string
}

func (e *ErrUnsupportedAlgorithmType) Error() string { // Format error message
	return "unsupported compression algorithm: " + e.Algorithm
}

type Compressor interface { // Common interface for compressors
	Compress(data []byte) ([]byte, error)
	Decompress(data []byte) ([]byte, error)
}

func NewCompressor(algorithm string) (Compressor, error) { // Factory function for compressors
	switch algorithm {
	case algorithms.RLEAlgorithm:
		return algorithms.NewRLECompressor(), nil
	default:
		return nil, &ErrUnsupportedAlgorithmType{Algorithm: algorithm}
	}
}