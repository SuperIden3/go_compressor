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
	// IsLeaf bool // Whether this node is a leaf
}

type HuffmanArray struct {
	Array []*HuffmanNode
	Length uint
}

// Helper Functions

// Swap two nodes in the array
func (h *HuffmanArray) HuffmanSwap(i, j uint) {
	h.Array[i], h.Array[j] = h.Array[j], h.Array[i]
	huffmanVerbosePrintf("Swapped nodes %v and %v\n", i, j)
}

// Get the parent of a node
func HuffmanParent(i uint) uint {
	return (i - 1) / 2
}

// Get the left child of a node
func HuffmanLeft(i uint) uint {
	return 2 * i + 1
}

// Get the right child of a node
func HuffmanRight(i uint) uint {
	return 2 * i + 2
}

func (n *HuffmanNode) IsLeaf() bool {
	return n.Left == nil && n.Right == nil
}

// MinHeapify ensures that the smallest frequency node is at the root of the heap
func (h *HuffmanArray) MinHeapify(i uint) {
	// Find the smallest frequency node in the subtree rooted at i
	smallest := i
	left := HuffmanLeft(i)
	right := HuffmanRight(i)

	if left < h.Length && h.Array[left].Weight < h.Array[smallest].Weight {
		// If the left child exists and its frequency is smaller than the root, update smallest
		smallest = left
	}
	if right < h.Length && h.Array[right].Weight < h.Array[smallest].Weight {
		// If the right child exists and its frequency is smaller than the root, update smallest
		smallest = right
	}
	if smallest != i {
		// If the smallest frequency node is not the root, HuffmanSwap it with the root
		h.HuffmanSwap(i, smallest)
		// Recursively call MinHeapify to ensure that the subtree rooted at smallest is a heap
		h.MinHeapify(smallest)
	}
}

// PopSmallest removes the smallest frequency node from the heap and returns it.
// The heap is reordered to maintain the heap property after removal.
func (h *HuffmanArray) PopSmallest() *HuffmanNode {
	// The smallest frequency node is always at the root of the heap
	smallest := h.Array[0]

	// Replace the root node with the last node in the heap
	h.Array[0] = h.Array[h.Length-1]

	// Decrement the length of the heap
	h.Length--

	// Reorder the heap to maintain the heap property
	h.MinHeapify(0)

	// Return the smallest node
	huffmanVerbosePrintf("PopSmallest: Returned node with value %v and frequency %v\n", smallest.Value, smallest.Weight)
	return smallest
}

// --- // Huffman encoding

// HuffmanMappify maps each unique byte in the data to the number of times it occurs.
// This function is used in the Huffman encoding algorithm to create a frequency map of the input data.
func HuffmanMappify(data []byte) HuffmanFrequencyMap {
	// Initialize the frequency map
	frequencyMap := make(HuffmanFrequencyMap)

	// Iterate through the data and count the occurrences of each byte
	for _, b := range data {
		// Increment the count of the byte in the frequency map
		frequencyMap[b]++
		huffmanVerbosePrintf("HuffmanMappify: frequencyMap: %v\n", frequencyMap)
	}
	return frequencyMap
}

// Build the initial heap from the frequency map
// This function takes a frequency map and returns a min-heap containing the same data.
// The heap is a binary tree where each node is smaller than its children.
func HuffmanBuildHeap(frequencyMap HuffmanFrequencyMap) *HuffmanArray {
	// Initialize the heap with the same number of elements as the frequency map
	map_len := len(frequencyMap)
	huffmanHeap := &HuffmanArray{
		Array: make([]*HuffmanNode, 0, map_len),
		Length: uint(map_len),
	}

	// Iterate through the frequency map and add each byte and its frequency to the heap
	for b, freq := range frequencyMap {
		huffmanHeap.Array = append(huffmanHeap.Array, &HuffmanNode{
			Value: b,
			Weight: freq,
		})
		// Log the creation of each node
		huffmanVerbosePrintf("HuffmanBuildHeap: Added node for byte %v with frequency %v\n", b, freq)
	}

	// Return the heap
	return huffmanHeap
}

// HuffmanEncodeHelper is a helper function for HuffmanEncode.
// It takes a node from the Huffman tree and the input data, and
// appends the encoded data to the encodedData slice.
func HuffmanEncodeHelper(node *HuffmanNode, data []byte, encodedData *[]byte) {
	// If the node is a leaf node, then it represents a single byte
	// of data. Iterate through the data and append the byte to the
	// encodedData slice if it matches the value of the node.
	if node.IsLeaf() {
		for _, b := range data {
			if b == node.Value {
				*encodedData = append(*encodedData, node.Value)
			}
		}
		return
	}
	// Recursively call HuffmanEncodeHelper on the left and right children of the node.
	HuffmanEncodeHelper(node.Left, data, encodedData)
	HuffmanEncodeHelper(node.Right, data, encodedData)
}

// HuffmanEncode takes a slice of bytes and returns a slice of bytes representing the Huffman encoded data.
func HuffmanEncode(data []byte) ([]byte, HuffmanFrequencyMap, error) {
	// Check if the data is nil or empty
	if data == nil {
		return []byte{}, nil, fmt.Errorf("data is nil")
	}
	data_len := len(data)
	if data_len == 0 {
		return []byte{}, nil, fmt.Errorf("data is empty")
	}

	// Create a frequency map of the input data
	frequencyMap := HuffmanMappify(data)
	huffmanVerbosePrintf("HuffmanEncode: frequencyMap: %v\n", frequencyMap)

	// Build the initial heap from the frequency map
	huffmanHeap := HuffmanBuildHeap(frequencyMap)
	huffmanVerbosePrintf("HuffmanEncode: huffmanHeap: %v\n", huffmanHeap)

	// Merge the nodes in the heap until only one node remains
	for huffmanHeap.Length > 1 {
		left, right := huffmanHeap.PopSmallest(), huffmanHeap.PopSmallest()

		// Create a new node with the left and right nodes as children
		mergedNode := &HuffmanNode{
			Left:  left,
			Right: right,
			Weight: left.Weight + right.Weight,
		}
		huffmanVerbosePrintf("HuffmanEncode: Merged nodes with frequencies %v and %v into new node with frequency %v\n", left.Weight, right.Weight, mergedNode.Weight)

		// Add the merged node back to the heap
		huffmanHeap.Array = append(huffmanHeap.Array, mergedNode)
		huffmanHeap.Length++
		huffmanHeap.MinHeapify(huffmanHeap.Length - 1)
		huffmanVerbosePrintf("HuffmanEncode: huffmanHeap: %v\n", huffmanHeap)
	}

	// Pop the only node remaining in the heap
	root := huffmanHeap.PopSmallest()
	huffmanVerbosePrintf("HuffmanEncode: root: %v\n", root)

	// Initialize the encoded data slice
	encodedData := make([]byte, 0, data_len)
	HuffmanEncodeHelper(root, data, &encodedData)
	huffmanVerbosePrintf("HuffmanEncode: encodedData: %v\n", encodedData)

	return encodedData, frequencyMap, nil
}

// --- // Huffman decoding

func HuffmanDecode(encodedData []byte, frequencyMap HuffmanFrequencyMap) ([]byte, error) {
	// Check if the encoded data is nil or empty
	if encodedData == nil {
		return []byte{}, fmt.Errorf("encoded data is nil")
	}
	encodedData_len := len(encodedData)
	if encodedData_len == 0 {
		return []byte{}, fmt.Errorf("encoded data is empty")
	}

	// Check if the frequency map is nil or empty
	if frequencyMap == nil {
		return []byte{}, fmt.Errorf("frequency map is nil")
	}
	frequencyMap_len := len(frequencyMap)
	if frequencyMap_len == 0 {
		return []byte{}, fmt.Errorf("frequency map is empty")
	}

	// Build the Huffman tree from the frequency map
	huffmanHeap := HuffmanBuildHeap(frequencyMap)
	huffmanVerbosePrintf("HuffmanDecode: huffmanHeap: %v\n", huffmanHeap)

	// Merge the nodes in the heap until only one node remains
	for huffmanHeap.Length > 1 {
		left, right := huffmanHeap.PopSmallest(), huffmanHeap.PopSmallest()

		// Create a new node with the left and right nodes as children
		mergedNode := &HuffmanNode{
			Left:  left,
			Right: right,
			Weight: left.Weight + right.Weight,
		}
		huffmanVerbosePrintf("HuffmanDecode: Merged nodes with frequencies %v and %v into new node with frequency %v\n", left.Weight, right.Weight, mergedNode.Weight)

		// Add the merged node back to the heap
		huffmanHeap.Array = append(huffmanHeap.Array, mergedNode)
		huffmanHeap.Length++
		huffmanHeap.MinHeapify(huffmanHeap.Length - 1)
		huffmanVerbosePrintf("HuffmanDecode: huffmanHeap: %v\n", huffmanHeap)
	}

	// Pop the only node remaining in the heap
	root := huffmanHeap.PopSmallest()
	huffmanVerbosePrintf("HuffmanDecode: root: %v\n", root)

	// Initialize the decoded data slice
	decodedData := make([]byte, 0, encodedData_len)

	// Iterate through the encoded data and traverse the Huffman tree to decode it
	for _, b := range encodedData {
		currentNode := root

		for !currentNode.IsLeaf() {
			if b == currentNode.Left.Value {
				currentNode = currentNode.Left
			} else if b == currentNode.Right.Value {
				currentNode = currentNode.Right
			} else {
				return []byte{}, fmt.Errorf("invalid encoded data: byte %v not found in Huffman tree", b)
			}
		}
		decodedData = append(decodedData, currentNode.Value)
	}

	return decodedData, nil
}
