package algorithms

import (
	"bytes"
	"fmt"
)

// Helper function to handle the writing logic and centralize error checking.
// FIX: Now accepts a pointer to bytes.Buffer.
func writePair(buffer *bytes.Buffer, c byte, char byte) error {
	// We ignore 'n' (number of bytes written) because we know WriteByte always writes 1 byte on success.
	if _, err := buffer.WriteByte(c); err != nil {
		return fmt.Errorf("failed to write count byte: %w", err)
	}
	if _, err := buffer.WriteByte(char); err != nil {
		return fmt.Errorf("failed to write data byte: %w", err)
	}
	return nil
}

// Encodes data using the RLE compression method, returning a byte slice.
// FIX: Changed return type from (string, error) to ([]byte, error) for correct handling of binary data.
func rle(data string) ([]byte, error) {
	DATA_LEN := len(data)

	if DATA_LEN == 0 {
		return nil, nil // Return nil slice for empty input
	}

	var buffer bytes.Buffer // Initialize the empty buffer for storing the compressed data
	count := 1              // Store the count for repeated characters

	// Note: The maximum run length this implementation can correctly encode is 255
	// due to the use of byte(count).

	for i := 1; i < DATA_LEN; i++ { // Head start at index 1 to compare the byte behind the current byte
		if data[i] == data[i-1] {
			count++ // Pattern found
		} else {
			// FIX: Pass the address of the buffer (&buffer) to the helper function.
			if err := writePair(&buffer, byte(count), data[i-1]); err != nil { // Write the count as a byte and the repeated byte character
				return nil, err // Return nil slice if an error occurs
			}
			count = 1 // Reset count
		}
	}

	// Write the last character and its count
	// FIX: Pass the address of the buffer (&buffer) to the helper function.
	if err := writePair(&buffer, byte(count), data[DATA_LEN - 1]); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

// If you MUST return a string, you can convert the byte slice to a string:
func rleAsString(data string) (string, error) {
    compressedBytes, err := rle(data) // Use the fixed RLE function
    if err != nil {
        return "", err
    }
    // WARNING: This conversion is dangerous for binary data.
    return string(compressedBytes), nil
}
