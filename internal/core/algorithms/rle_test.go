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
		{
			name:     "Mixed characters",
			input:    "AABBBCCCCDDDDDE",
			expected: []byte{2, 'A', 3, 'B', 4, 'C', 5, 'D', 1, 'E'},
		},
		{
			name:     "No repeated characters",
			input:    "ABCDEFG",
			expected: []byte{1, 'A', 1, 'B', 1, 'C', 1, 'D', 1, 'E', 1, 'F', 1, 'G'},
		},
		{
			name:     "Overflow",
			input:    string(bytes.Repeat([]byte{'A'}, 300)), // 300 A's
			expected: []byte{255, 'A', 45, 'A'}, // 255 + 45
		},
	}

	// 2. Run the tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RleAsString(tt.input)

			// Check for unexpected errors
			if err != nil {
				t.Fatalf("Rle(%q) returned unexpected error: %v", tt.input, err)
			}

			// Compare slices using bytes.Equal
			if !bytes.Equal([]byte(got), tt.expected) {
				t.Errorf("Rle(%q) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
