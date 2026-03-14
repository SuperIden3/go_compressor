package algorithms;

import (
	"fmt"
)

// Logging

var HuffmanQuiet = false
func huffmanGeneralPrintf(format string, v ...interface{}) {
	if !HuffmanQuiet {
		fmt.Printf("Huffman: " + format, v...)
	}
}

var HuffmanVerbose = false
func huffmanVerbosePrintf(format string, v ...interface{}) {
	if HuffmanVerbose {
		huffmanGeneralPrintf("HuffmanVerbose: " + format, v...)
	}
}

// --- // Huffman Functions

type HuffmanFrequencyMap map[byte]uint

type HuffmanNode struct {
	Left  *HuffmanNode
	Right *HuffmanNode
	Value byte
}

type HuffmanTree struct {
	Root *HuffmanNode
}


