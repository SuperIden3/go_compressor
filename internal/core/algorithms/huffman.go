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

// --- // Huffman helper types and functions

type HuffmanFrequencyMap map[byte]uint // For counting how many each unique byte occurs

// Min Heap Structure for Huffman: Array
// Implement the min heap structure as an array for efficiency
// Where `i` is a node's index:
//   - Left child: `2 * i + 1`
//   - Right child: `2 * i + 2`
//   - Parent: `floor((i - 1) / 2)`

type HuffmanNode struct {
	Left  *HuffmanNode
	Right *HuffmanNode
	Weight uint // The number of times the byte occurs
	Value byte // The unique byte
	IsLeaf bool // Whether this node is a leaf
}

type HuffmanArray struct {
	Array []HuffmanNode
	Length uint
}

// Helper Functions

// Swap two nodes in the array
func (h *HuffmanArray) Swap(i, j uint) {
	h.Array[i], h.Array[j] = h.Array[j], h.Array[i]
	HuffmanVerbosePrintf("Swappped nodes %v and %v\n", i, j)
}

// Get the parent of a node
func (h *HuffmanArray) Parent(i uint) uint {
	return (i - 1) / 2
}

// Get the left child of a node
func (h *HuffmanArray) Left(i uint) uint {
	return 2 * i + 1
}

// Get the right child of a node
func (h *HuffmanArray) Right(i uint) uint {
	return 2 * i + 2
}

func HuffmanMappify(data []byte) HuffmanFrequencyMap {
	frequencyMap := make(HuffmanFrequencyMap)
	for _, b := range data {
		frequencyMap[b]++
		HuffmanVerbosePrintf("HuffmanMappify: frequencyMap: %v\n", frequencyMap)
	}
	return frequencyMap
}

// --- // Huffman encoding

func HuffmanEncode(data []byte) ([]byte, error) {
	if data == nil {
		return []byte{}, fmt.Errorf("data is nil")
	}
	data_len := len(data)
	if data_len == 0 {
		return []byte{}, fmt.Errorf("data is empty")
	}

	frequencyMap := HuffmanMappify(data)
	huffmanVerbosePrintf("HuffmanEncode: frequencyMap: %v\n", frequencyMap)
}

