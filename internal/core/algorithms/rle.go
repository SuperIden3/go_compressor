package algorithms;

import (
	"bytes"
	"fmt"
)

// Helper function to handle the writing logic and centralize error checking
func writePair(c byte, char byte) error {
	// We ignore 'n' (number of bytes written) because we know WriteByte always writes 1 byte on success.
	if _, err := buffer.WriteByte(c); err != nil {
		return fmt.Errorf("failed to write count byte: %w", err)
	}
	if _, err := buffer.WriteByte(char); err != nil {
		return fmt.Errorf("failed to write data byte: %w", err)
	}
	return nil
}

// Encrypts data using the RLE compression method.
// RLE scans the data and replaces repeating consecutive characters with two characters, them being a binary character that has the hexadecimal value of how many of those characters have repeated, followed by a single one of that character that has been repeated.
// Example: "aaaaaaaaaabbbbbbbbb" -> "(NEWLINE)a(TAB)b", where "a" repeats ten times and "b" repeats nine times, represented by a newline and a tab.
func rle(data string) (string, error) {
	if len(data) == 0 {
		return "", nil
	} // Handle empty data

	var buffer bytes.Buffer // Initialize the empty buffer for storing the compressed data
	count := 1              // Store the count for repeated characters

	for i := 1; i < len(data); i++ { // Start at 1 to compare the previous to see if they're the same
		if data[i] == data[i-1] { // Compare characters at current position and previous position
			count++ // Increment count
		} else { // Write the count as a character and then the character repeated after that
			if err := writePair(byte(count), data[i-1]); err != nil {
				return "", err // Return immediately if an error occurs
			}
			count = 1 // Reset count
		}
	}

	// Write the last character and its count
	if err := writePair(byte(count), data[len(data)-1]); err != nil {
		return "", err
	}

	return buffer.String(), nil // Return string with no error
}
