package main

import (
	"fmt"
	"os"
	"github.com/d3mondev/puredns/internal/brute"
	"github.com/d3mondev/puredns/internal/output"
)

func main() {
	options, err := ParseOptions()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if options.Help {
		PrintUsage()
		os.Exit(0)
	}

	resultsChan, err := brute.Run(options.Domain, options.Wordlist, options.Resolvers, 50)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	var results []brute.Result
	for result := range resultsChan {
		if !options.Quiet {
			output.PrintResult(result)
		}
		results = append(results, result)
	}

	if options.Output != "" {
		err := output.WriteResults(results, options.Output, "txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing results: %v\n", err)
		}
	}
}
