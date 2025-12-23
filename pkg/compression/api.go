package compression

import (
	"github.com/superiden3/go_compress/internal/core"
)

// Custom error for unsupported algorithms
type ErrUnsupportedAlgorithmType = core.ErrUnsupportedAlgorithmType

// Common interface for compressors
type Compressor = core.Compressor

// Common interface for decompressors
type Decompressor = core.Decompressor

// Interfaces for file-to-file operations
type FileToFileCompressor = core.FileToFileCompressor

// Interfaces for file-to-file decompression operations
type FileToFileDecompressor = core.FileToFileDecompressor

// NewCompressor creates a new Compressor based on the specified algorithm type.
func NewCompressor(algorithm int) (Compressor, error) {
	return core.NewCompressor(algorithm)
}

// NewDecompressor creates a new Decompressor based on the specified algorithm type.
func NewDecompressor(algorithm int) (Decompressor, error) {
	return core.NewDecompressor(algorithm)
}

// NewFileToFileCompressor creates a new FileToFileCompressor based on the specified algorithm type.
func NewFileToFileCompressor(algorithm int) (FileToFileCompressor, error) {
	return core.NewFileToFileCompressor(algorithm)
}

// NewFileToFileDecompressor creates a new FileToFileDecompressor based on the specified algorithm type.
func NewFileToFileDecompressor(algorithm int) (FileToFileDecompressor, error) {
	return core.NewFileToFileDecompressor(algorithm)
}