package algorithms

import (
	"bytes"
	"fmt"
	"os"
)

// Functions for logging

var RleQuiet = false
func generalPrintf(format string, v ...interface{}) {
	if !RleQuiet {
		fmt.Printf("Rle: " + format, v...)
	}
}

var RleVerbose = false
func verbosePrintf(format string, v ...interface{}) {
	if RleVerbose {
		generalPrintf("RleVerbose: " + format, v...)
	}
}

// --- // RLE Encoding

// Helper function to handle the writing logic and centralize error checking.
// FIX: Now accepts a pointer to bytes.Buffer.
func writePair(buffer *bytes.Buffer, c byte, char byte) error {
	// We ignore 'n' (number of bytes written) because we know WriteByte always writes 1 byte on success.
	if err := buffer.WriteByte(c); err != nil {
		generalPrintf("writePair: err: %v\n", err)
		return fmt.Errorf("failed to write count byte: %w", err)
	}
	if err := buffer.WriteByte(char); err != nil {
		generalPrintf("writePair: err: %v\n", err)
		return fmt.Errorf("failed to write data byte: %w", err)
	}
	verbosePrintf("writePair: count: %v, char: %v\n", c, char)
	return nil
}

// Encodes data using the RLE compression method, returning a byte slice.
// FIX: Changed return type from (string, error) to ([]byte, error) for correct handling of binary data.
func Rle(data []byte) ([]byte, error) {
	DATA_LEN := len(data)
	
	verbosePrintf("Rle: DATA_LEN: %v\n", DATA_LEN)

	if DATA_LEN == 0 {
		verbosePrintf("Rle: DATA_LEN == 0\n")
		return nil, nil // Return nil slice for empty input
	}

	var buffer bytes.Buffer // Initialize the empty buffer for storing the compressed data
	count := 1              // Store the count for repeated characters
	
	// Note: The maximum run length this implementation can correctly encode is 255
	// due to the use of byte(count).

	for i := 1; i < DATA_LEN; i++ { // Head start at index 1 to compare the byte behind the current byte
		verbosePrintf("Rle: i: %v\n", i)
		if data[i] == data[i - 1] && count < 255 { // If the current byte is the same as the previous one and the count is less than 255
			count++ // Pattern found
			verbosePrintf("Rle: count: %v\n", count)
		} else {
			// FIX: Pass the address of the buffer (&buffer) to the helper function.
			if err := writePair(&buffer, byte(count), data[i - 1]); err != nil { // Write the count as a byte and the repeated byte character
				generalPrintf("writePair: err: %v\n", err)
				return nil, err // Return nil slice if an error occurs
			}
			count = 1 // Reset count
			verbosePrintf("Rle: count reset to 1\n")
		}
	}

	// Write the last character and its count
	// FIX: Pass the address of the buffer (&buffer) to the helper function.
	if err := writePair(&buffer, byte(count), data[DATA_LEN - 1]); err != nil {
		generalPrintf("Rle: writePair: err: %v\n", err)
		return nil, err
	}

	return buffer.Bytes(), nil
}

// If you MUST return a string, you can convert the byte slice to a string:
func RleAsString(data string) (string, error) {
	verbosePrintf("RleAsString: data: %v\n", data)
	compressedBytes, err := Rle([]byte(data)) // Use the fixed RLE function
	if err != nil {
		return "", err
	}
	// WARNING: This conversion is dangerous for binary data.
	return string(compressedBytes), nil
}

// --- // RLE Decoding

func RleDecode(data []byte) ([]byte, error) {
	DATA_LEN := len(data)

	verbosePrintf("RleDecode: DATA_LEN: %v\n", DATA_LEN)

	if DATA_LEN == 0 {
		verbosePrintf("RleDecode: DATA_LEN == 0\n")
		return nil, nil // Return nil slice for empty input
	}

	var buffer bytes.Buffer // Initialize the empty buffer for storing the decompressed data

	for i := 0; i < DATA_LEN; i += 2 { // Increment by 2 to read count-character pairs
		verbosePrintf("i: %v", i)
		if i + 1 >= DATA_LEN {
			generalPrintf("RleDecode: err: %v\n", "malformed RLE data: incomplete pair at index %d", i)
			return nil, fmt.Errorf("malformed RLE data: incomplete pair at index %d", i)
		}

		count := int(data[i]) // Read the count byte
		char := data[i + 1]   // Read the character byte
		verbosePrintf("RleDecode: count: %v, char: %v\n", count, char)

		for j := 0; j < count; j++ { // Write the character 'count' times
			verbosePrintf("RleDecode: j: %v\n", j)
			if err := buffer.WriteByte(char); err != nil { // Try to write the byte to the buffer
				generalPrintf("RleDecode: err: %v\n", err)
				return nil, fmt.Errorf("failed to write decompressed byte: %w", err)
			}
		}
	}

	return buffer.Bytes(), nil
}

func RleDecodeAsString(data []byte) (string, error) {
	decompressedBytes, err := RleDecode(data) // Use the fixed RLE decode function
	verbosePrintf("RleDecodeAsString: decompressedBytes: %v\n", decompressedBytes)
	if err != nil {
		generalPrintf("RleDecodeAsString: err: %v\n", err)
		return "", err
	}
	return string(decompressedBytes), nil
}

// --- // RLE Compressor Interface

type RLECompressor struct {}

// Compress implements core.Compressor.
func (r *RLECompressor) Compress(data []byte) ([]byte, error) {
	// Convert input bytes to string for the existing Rle function.
	if len(data) == 0 {
		verbosePrintf("RLECompressor: len(data) == 0")
		return nil, nil
	}

	compressedData, err := Rle(data)
	verbosePrintf("RLECompressor: compressedData: %v\n", compressedData)
	if err != nil {
		generalPrintf("RLECompressor: err: %v\n", err)
		return nil, err
	}

	return compressedData, nil
}

// Decompress implements core.Compressor.
func (r *RLECompressor) Decompress(data []byte) ([]byte, error) {
	if len(data) == 0 {
		verbosePrintf("RLECompressor: len(data) == 0\n")
		return nil, nil
	}

	decompressedData, err := RleDecode(data)
	verbosePrintf("RLECompressor: decompressedData: %v\n", decompressedData)
	if err != nil {
		generalPrintf("RLECompressor: err: %v\n", err)
		return nil, err
	}

	return decompressedData, nil
}

// Factory functions for creating instances of RLECompressor.
func NewRLECompressor() *RLECompressor {
	return &RLECompressor {}
}

// Factory function for creating a decompressor instance.
func NewRLEDecompressor() *RLECompressor {
	return &RLECompressor {}
}

// --- // File To File Compressing and Decompressing

// RleCompressFile compresses a file and writes the result to another file.
func RleCompressFile(inputFilePath string, outputFilePath string) error {
	// Read the input file content.
	inputData, err := os.ReadFile(inputFilePath)
	generalPrintf("RleCompressFile: Reading from \"%v\" and writing to \"%v\"\n", inputFilePath, outputFilePath)
	verbosePrintf("RleCompressFile: inputData: %v\n", inputData)
	if err != nil {
		generalPrintf("RleCompressFile: err: %v\n", err)
		return fmt.Errorf("failed to read input file: %w", err)
	}

	// Compress the data.
	compressedData, err := Rle(inputData)
	verbosePrintf("RleCompressFile: compressedData: %v\n", compressedData)
	if err != nil {
		verbosePrintf("RleCompressFile: err: %v\n", err)
		return fmt.Errorf("failed to compress data: %w", err)
	}

	// Write the compressed data to the output file.
	if err := os.WriteFile(outputFilePath, compressedData, 0644); err != nil {
		generalPrintf("RleCompressFile: err: %v\n", err)
		return fmt.Errorf("failed to write output file: %w", err)
	}

	return nil
}

// RleDecompressFile decompresses a file and writes the result to another file.
func RleDecompressFile(inputFilePath string, outputFilePath string) error {
	// Read the input file content.
	inputData, err := os.ReadFile(inputFilePath)
	generalPrintf("RleDecompressFile: Reading from \"%v\" and writing to \"%v\"\n", inputFilePath, outputFilePath)
	verbosePrintf("RleDecompressFile: inputData: %v\n", inputData)
	if err != nil {
		generalPrintf("RleDecompressFile: err: %v\n", err)
		return fmt.Errorf("failed to read input file: %w", err)
	}

	// Decompress the data.
	decompressedData, err := RleDecode(inputData)
	verbosePrintf("RleDecompressFile: decompressedData: %v\n", decompressedData)
	if err != nil {
		generalPrintf("RleDecompressFile: err: %v\n", err)
		return fmt.Errorf("failed to decompress data: %w", err)
	}

	// Write the decompressed data to the output file.
	if err := os.WriteFile(outputFilePath, decompressedData, 0644); err != nil {
		verbosePrintf("RleDecompressFile: err: %v\n", err)
		return fmt.Errorf("failed to write output file: %w", err)
	}

	return nil
}

type RLEFileToFileCompressor struct {}
type RLEFileToFileDecompressor struct {}

// CompressFile implements core.FileToFileCompressor.
func (r *RLEFileToFileCompressor) CompressFileToFile(inputFilePath string, outputFilePath string) error {
	return RleCompressFile(inputFilePath, outputFilePath)
}

// DecompressFile implements core.FileToFileDecompressor.
func (r *RLEFileToFileDecompressor) DecompressFileToFile(inputFilePath string, outputFilePath string) error {
	return RleDecompressFile(inputFilePath, outputFilePath)
}

// Factory functions for creating instances of RLEFileToFileCompressor.
func NewRLEFileToFileCompressor() *RLEFileToFileCompressor {
	return &RLEFileToFileCompressor {}
}

// Factory function for creating a decompressor instance.
func NewRLEFileToFileDecompressor() *RLEFileToFileDecompressor {
	return &RLEFileToFileDecompressor {}
}
