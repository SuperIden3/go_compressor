package algorithms;

import (
	"fmt"
	"container/heap"
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

// --- // Huffman Tree Type and Functions

// Huffman_PositiveCount is a positive integer type for counting the number of occurrences of each byte.
type Huffman_PositiveCount uint

// HuffmanNode is a node in the Huffman tree.
type HuffmanNode struct {
	Char  byte
	Count Huffman_PositiveCount
	Left  *HuffmanNode
	Right *HuffmanNode
}

// HuffmanCodeMap is a map of bytes to Huffman codes. Measures frequency of each byte.
type HuffmanCodeMap map[byte][]bool

// HuffmanHeap is a min-heap of HuffmanByteCountt. Heap is implemented array-method.
type HuffmanHeap []HuffmanByteCount

// --- // Huffman Heap Functions

// Length of a HuffmanHeap
func (h HuffmanHeap) Len() int {
	return len(h)
}

// Max heap: for merging lower numbers at ends
func (h HuffmanHeap) Less(i, j int) bool {
	return h[i].Count > h[j].Count
}

// Swap two HuffmanByteCount
func (h HuffmanHeap) Swap(i, j int) {
	HuffmanVerbosePrintf("Swap - i: %v, j: %v\n", i, j)
	h[i], h[j] = h[j], h[i]
}

// Push a HuffmanByteCount
func (h *HuffmanHeap) Push(x interface{}) {
	*h = append(*h, x.(HuffmanByteCount))
}

// Pop and get a HuffmanByteCount
func (h *HuffmanHeap) Pop() interface{} {
	last_index := (*h).Len() - 1
	x := (*h)[last_index]
	*h = (*h)[0 : last_index]
	return x
}

// --- // Huffman Helper Functions

// ByteFrequencyMap returns a map of bytes to their frequency
func MakeByteFrequencyMap(data []byte) HuffmanCodeMap {
	ret := make(HuffmanCodeMap)

	if nil == data {
		return ret
	}

	for _, c := range data {
		ret[c]++
		HuffmanVerbosePrintf("ByteFrequencyMap: c: %v with count %v\n", c, ret[c])
	}
	HuffmanVerbosePrintf("ByteFrequencyMap: ret: %v\n", ret)

	return ret
}

func MergeHuffmanNodes(left, right *HuffmanNode) *HuffmanNode {
	return &HuffmanNode{
		Count: left.Count + right.Count,
		Left:  nil,
		Right: nil,
	}
}

func MergeHuffmanTree(root *HuffmanNode) HuffmanCodeMap {
	ret := make(HuffmanCodeMap)
	
	if nil == root {
		return ret
	}

	// TODO
}

// --- // Huffman Encode and Decode

func HuffmanEncode(data []byte) ([]byte, error) {
	if nil == data {
		return nil, fmt.Errorf("HuffmanEncode: data is nil")
	}
}

