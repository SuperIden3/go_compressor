package core

import "github.com/superiden3/go_compress/internal/core/algorithms"

// Custom error for unsupported algorithms
type ErrUnsupportedAlgorithmType struct {
	Algorithm string // Name of the unsupported algorithm
}

// Format the ErrUnsupportedAlgorithmType error message.
func (e *ErrUnsupportedAlgorithmType) Error() string {
	return "unsupported compression algorithm: " + e.Algorithm
}

// Common interface for compressors
type Compressor interface {
	Compress(data []byte) ([]byte, error)
}

// Common interface for decompressors
type Decompressor interface {
	Decompress(data []byte) ([]byte, error)
}

// Interfaces for file-to-file operations
type FileToFileCompressor interface {
	CompressFileToFile(inputPath, outputPath string) error
}

// Interfaces for file-to-file decompression operations
type FileToFileDecompressor interface {
	DecompressFileToFile(inputPath, outputPath string) error
}

// NewCompressor creates a new Compressor based on the specified algorithm type, which is an int meant for the `Algorithms` array in `implemented.go`.
func NewCompressor(algorithm int) (Compressor, error) { // Factory function for compressors
	switch algorithm {
	case algorithms.RLEAlgorithm:
		return algorithms.NewRLECompressor(), nil
	default:
		return nil, &ErrUnsupportedAlgorithmType { Algorithm: algorithms.Algorithms[algorithm] }
	}
}

// NewDecompressor creates a new Decompressor based on the specified algorithm type, which is an int meant for the `Algorithms` array in `implemented.go`.
func NewDecompressor(algorithm int) (Decompressor, error) { // Factory function for decompressors
	switch algorithm {
	case algorithms.RLEAlgorithm:
		return algorithms.NewRLEDecompressor(), nil
	default:
		return nil, &ErrUnsupportedAlgorithmType { Algorithm: algorithms.Algorithms[algorithm] }
	}
}

// NewFileToFileCompressor creates a new FileToFileCompressor based on the specified algorithm type, which is an int meant for the `Algorithms` array in `implemented.go`.
func NewFileToFileCompressor(algorithm int) (FileToFileCompressor, error) { // Factory function for file-to-file compressors
	switch algorithm {
	case algorithms.RLEAlgorithm:
		return algorithms.NewRLEFileToFileCompressor(), nil
	default:
		return nil, &ErrUnsupportedAlgorithmType { Algorithm: algorithms.Algorithms[algorithm] }
	}
}

// NewFileToFileDecompressor creates a new FileToFileDecompressor based on the specified algorithm type, which is an int meant for the `Algorithms` array in `implemented.go`.
func NewFileToFileDecompressor(algorithm int) (FileToFileDecompressor, error) { // Factory function for file-to-file decompressors
	switch algorithm {
	case algorithms.RLEAlgorithm:
		return algorithms.NewRLEFileToFileDecompressor(), nil
	default:
		return nil, &ErrUnsupportedAlgorithmType { Algorithm: algorithms.Algorithms[algorithm] }
	}
}