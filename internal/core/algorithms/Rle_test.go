package algorithms

import (
	"bytes"
	"testing"
)

func TestRle(t *testing.T) {
	// 1. Define our test cases
	tests := []struct {
		name     string
		input    string
		expected []byte // Since rle returns []byte
	} {
		{
			name:     "Empty string",
			input:    "",
			expected: nil,
		},
		{
			name:     "Single character",
			input:    "A",
			expected: []byte{1, 'A'},
		},
		{
			name:     "Repeated characters",
			input:    "AAABBC",
			expected: []byte{3, 'A', 2, 'B', 1, 'C'},
		},
		{
			name:     "Long run (boundary check)",
			input:    "AAAAAAAAAA", // 10 A's
			expected: []byte{10, 'A'},
		},
	}

	// 2. Run the tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Rle(tt.input)

			// Check for unexpected errors
			if err != nil {
				t.Fatalf("Rle(%q) returned unexpected error: %v", tt.input, err)
			}

			// Compare slices using bytes.Equal
			if !bytes.Equal(got, tt.expected) {
				t.Errorf("Rle(%q) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
