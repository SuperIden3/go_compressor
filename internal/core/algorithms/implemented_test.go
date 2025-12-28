package algorithms

import (
	"testing"
)

func TestOrderOfAlgorithms(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	} {
		{
			name:  "Run-Length Encryption",
			input: RLEAlgorithm,
			expected: "rle",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetAlgorithmName(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %s, but got %s", tt.expected, result)
			}
		})
	}
}