package algorithms

import "fmt"

var Algorithms = []string{ "rle" } // List of names of available (implemented) compression algorithms
var ImplementedAlgorithms = len(Algorithms) // Number of implemented algorithms

const ( // Constant integers for each algorithm; each one is aligned with its name in the Algorithms array
	RLEAlgorithm = iota // Run-Length Encoding
)

// Print the names of all available compression algorithms
func PrintAlgorithms() {
	fmt.Println("Available compression algorithms:")
	for _, alg := range Algorithms {
		fmt.Printf("- %s\n", alg)
	}
}

func GetAlgorithmName(alg int) string { // Get the name of an algorithm by its integer identifier
	if alg < 0 || alg >= len(Algorithms) {
		return "Unknown algorithm"
	}
	return Algorithms[alg]
}

func GetAlgorithmID(name string) int { // Get the integer identifier of an algorithm by its name
	for i, alg := range Algorithms {
		if alg == name {
			return i
		}
	}
	return -1 // Return -1 if the algorithm name is not found
}
