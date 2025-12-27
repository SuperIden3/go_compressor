package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/superiden3/go_compress/internal/core"
	"github.com/superiden3/go_compress/internal/core/algorithms"
)

// Printing usage info
func usage() {
	fmt.Println("Usage: go run main.go [options] <input_file> <output_file> [input_file2] [output_file2] ...")
	fmt.Println("Options:")
	flag.PrintDefaults()
}

// Main compressing function for `main` to use.
func mainCompress(alg_int int, wg *sync.WaitGroup) {
	// Checking for invalid arguments
	if alg_int < 0 || alg_int >= len(algorithms.Algorithms)  {
		fmt.Fprintf(os.Stderr, "Error: Unknown algorithm number %d\n", alg_int)
		return
	}

	// Create a new compressor with the selected algorithm
	compressor, err := core.NewFileToFileCompressor(alg_int) // Create a new compressor with the selected algorithm
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Failed to create compressor: %v\n", err)
		return
	}

	// Loop over the remaining arguments in pairs
	for i := 0; i < flag.NArg(); i += 2 {
		// Get input and output files
		inputFile := flag.Arg(i)
		outputFile := flag.Arg(i + 1)

		// Check if the input file exists
		_, err := os.Stat(inputFile)
		if os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "Error: Input file \"%s\" does not exist\n", inputFile)
			continue
		}

		// Increment the WaitGroup counter
		wg.Add(1)

		// Start a new goroutine to compress the file
		go func(inputFile, outputFile string) {
			defer wg.Done()

			// Compress the file
			err := compressor.CompressFileToFile(inputFile, outputFile)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: Failed to compress file \"%s\": %v\n", inputFile, err)
			}
		}(inputFile, outputFile)
	}

	// Wait for all goroutines to finish
	wg.Wait()
}

// Main decompressing function for `main` to use.
func mainDecompress(alg_int int, wg *sync.WaitGroup) {
	// Checking for invalid arguments
	if alg_int < 0 || alg_int >= len(algorithms.Algorithms)  {
		fmt.Fprintf(os.Stderr, "Error: Unknown algorithm number %d\n", alg_int)
		return
	}

	// Create a new decompressor with the selected algorithm
	decompressor, err := core.NewFileToFileDecompressor(alg_int)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Failed to create decompressor: %v\n", err)
		return
	}

	// Loop over the remaining arguments in pairs
	for i := 0; i < flag.NArg(); i += 2 {
		// Get input and output files
		inputFile := flag.Arg(i)
		outputFile := flag.Arg(i + 1)

		// Check if the input file exists
		_, err := os.Stat(inputFile)
		if os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "Error: Input file '%s' does not exist\n", inputFile)
			return
		}

		// Increment the WaitGroup counter
		wg.Add(1)

		// Start a new goroutine to decompress the file
		go func(inputFile, outputFile string) {
			defer wg.Done()

			// Decompress the file
			err := decompressor.DecompressFileToFile(inputFile, outputFile)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: Failed to decompress file '%s': %v\n", inputFile, err)
			}
		}(inputFile, outputFile)
	}

	// Wait for all goroutines to finish
	wg.Wait()
}

// Enable verbose logging
func verbosify(makeVerbose bool) {
	algorithms.RleVerbose = makeVerbose
}

// Make logging quiet
func quietify(makeQuiet bool) {
	algorithms.RleQuiet = makeQuiet
}

func main() {
	// Parse command-line arguments
	print_algs := flag.Bool("print-algorithms", false, "Print available compression algorithms and exit")
	alg := flag.String("algorithm", "rle", "Compression algorithm to use (default: rle)")
	decompress := flag.Bool("decompress", false, "Decompress the input file instead of compressing it")
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	quiet := flag.Bool("quiet", false, "Disable logging (overrides verbose)")
	flag.Usage = usage
	flag.Parse()

	// Print available algorithms if requested
	if *print_algs {
		algorithms.PrintAlgorithms()
		return
	}

	// Check for the correct number of arguments
	if flag.NArg() < 2 {
		usage()
		return
	}

	// Validate the selected algorithm
	alg_int := -1
	for i := 0; i < len(algorithms.Algorithms); i++ {
		if algorithms.Algorithms[i] == *alg {
			alg_int = i
			break
		}
		if i == len(algorithms.Algorithms)-1 {
			fmt.Fprintf(os.Stderr, "Error: Unknown algorithm '%s'\n", *alg)
			return
		}
	}

	// Enable quiet logging if requested (overrides verbose)
	if *quiet {
		quietify(true)
	} else {
		// Enable verbose logging if requested
		verbosify(*verbose)
	}

	// Create a new WaitGroup to manage goroutines
	wg := &sync.WaitGroup{}

	if !*decompress {
		// Compress the files
		mainCompress(alg_int, wg)
	} else {
		// Decompress the files
		mainDecompress(alg_int, wg)
	}
}
