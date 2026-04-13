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

// HuffmanByteCount is used for initial heap elements.
type HuffmanByteCount struct {
	Char  byte
	Count Huffman_PositiveCount
}

// HuffmanCodeMap is a map of bytes to Huffman codes.
type HuffmanCodeMap map[byte][]bool

// HuffmanHeap is a min-heap of HuffmanNodes.
type HuffmanHeap []*HuffmanNode

// --- // Huffman Heap Functions

// Length of a HuffmanHeap
func (h HuffmanHeap) Len() int {
	return len(h)
}

// Min-heap: smaller counts have higher priority for merging
func (h HuffmanHeap) Less(i, j int) bool {
	return h[i].Count < h[j].Count
}

// Swap two HuffmanNodes
func (h HuffmanHeap) Swap(i, j int) {
	HuffmanVerbosePrintf("Swap - i: %v, j: %v\n", i, j)
	h[i], h[j] = h[j], h[i]
}

// Push a HuffmanNode
func (h *HuffmanHeap) Push(x interface{}) {
	*h = append(*h, x.(*HuffmanNode))
}

// Pop and get a HuffmanNode
func (h *HuffmanHeap) Pop() interface{} {
	last_index := (*h).Len() - 1
	x := (*h)[last_index]
	*h = (*h)[0 : last_index]
	return x
}

// --- // Huffman Helper Functions

// MakeFrequencyMap returns a map of bytes to their frequency
func MakeFrequencyMap(data []byte) map[byte]Huffman_PositiveCount {
	ret := make(map[byte]Huffman_PositiveCount)  // Initialize empty frequency map

	if nil == data {
		return ret  // Return empty map if data is nil
	}

	for _, c := range data {  // Iterate over each byte in data
		ret[c]++  // Increment count for this byte
		HuffmanVerbosePrintf("MakeFrequencyMap: c: %v with count %v\n", c, ret[c])  // Verbose logging
	}
	HuffmanVerbosePrintf("MakeFrequencyMap: ret: %v\n", ret)  // Verbose logging of result

	return ret  // Return the frequency map
}

func MergeHuffmanNodes(left, right *HuffmanNode) *HuffmanNode {
	return &HuffmanNode{
		Count: left.Count + right.Count,
		Left:  nil,
		Right: nil,
	}
}

func MergeHuffmanTree(root *HuffmanNode) HuffmanCodeMap {
	ret := make(HuffmanCodeMap)  // Initialize empty code map
	
	if nil == root {
		return ret  // Return empty map if root is nil
	}

	buildCodes(root, []bool{}, ret)  // Build codes starting from root with empty code
	return ret  // Return the code map
}

// buildCodes recursively builds the Huffman codes for each character in the tree.
func buildCodes(node *HuffmanNode, code []bool, codes HuffmanCodeMap) {
	if node.Left == nil && node.Right == nil {  // If leaf node
		codes[node.Char] = make([]bool, len(code))  // Create copy of current code
		copy(codes[node.Char], code)  // Copy the code bits
		return  // Done for this leaf
	}
	if node.Left != nil {  // If left child exists
		buildCodes(node.Left, append(code, false), codes)  // Recurse left with 0 bit
	}
	if node.Right != nil {  // If right child exists
		buildCodes(node.Right, append(code, true), codes)  // Recurse right with 1 bit
	}
}

// --- // Huffman Encode and Decode

// HuffmanEncode encodes the input data using Huffman coding and returns the encoded byte slice.
func HuffmanEncode(data []byte) ([]byte, error) {
	if nil == data {  // Check if data is nil
		return nil, fmt.Errorf("HuffmanEncode: data is nil")  // Error for nil data
	}
	if len(data) == 0 {  // Check if data is empty
		return []byte{}, nil  // Return empty encoded data
	}

	freq := MakeFrequencyMap(data)  // Build frequency map of bytes
	if len(freq) == 0 {  // If no frequencies (shouldn't happen)
		return []byte{}, nil  // Return empty
	}

	h := &HuffmanHeap{}  // Initialize min-heap for Huffman tree building
	heap.Init(h)  // Initialize heap

	for char, count := range freq {  // For each unique byte
		node := &HuffmanNode{Char: char, Count: count}  // Create leaf node
		heap.Push(h, node)  // Push to heap
	}

	for h.Len() > 1 {  // While more than one node in heap
		left := heap.Pop(h).(*HuffmanNode)  // Pop smallest node
		right := heap.Pop(h).(*HuffmanNode)  // Pop next smallest
		merged := &HuffmanNode{  // Create merged node
			Count: left.Count + right.Count,  // Sum counts
			Left:  left,  // Left child
			Right: right,  // Right child
		}
		heap.Push(h, merged)  // Push merged back to heap
	}

	root := heap.Pop(h).(*HuffmanNode)  // The root of the Huffman tree
	codes := MergeHuffmanTree(root)  // Generate Huffman codes from tree

	var bits []bool  // Slice to collect all bits
	for _, b := range data {  // For each byte in data
		if code, ok := codes[b]; ok {  // Get code for this byte
			bits = append(bits, code...)  // Append code bits
		} else {
			return nil, fmt.Errorf("HuffmanEncode: no code for byte %v", b)  // Error if no code
		}
	}

	encoded := packBits(bits)  // Pack bits into bytes
	return encoded, nil  // Return encoded data
}

func packBits(bits []bool) []byte {
	var result []byte  // Result byte slice
	for i := 0; i < len(bits); i += 8 {  // Process bits in groups of 8
		var b byte  // Current byte to build
		for j := 0; j < 8 && i+j < len(bits); j++ {  // For each bit in byte
			if bits[i+j] {  // If bit is true (1)
				b |= 1 << (7 - j)  // Set bit in byte (MSB first)
			}
		}
		result = append(result, b)  // Append byte to result
	}
	return result  // Return packed bytes
}

