package main

import (
	"flag"
	"fmt"
)

type Options struct {
	Domain        string
	Wordlist      string
	Resolvers     string
	Output        string
	Quiet         bool
	WildcardCheck bool
	Help          bool
}

func ParseOptions() (*Options, error) {
	options := &Options{}

	flag.StringVar(&options.Domain, "d", "", "Target domain")
	flag.StringVar(&options.Wordlist, "w", "", "Path to wordlist file")
	flag.StringVar(&options.Resolvers, "r", "configs/resolvers.txt", "Path to resolvers file")
	flag.StringVar(&options.Output, "o", "", "Output file path")
	flag.BoolVar(&options.Quiet, "q", false, "Quiet mode")
	flag.BoolVar(&options.WildcardCheck, "wildcard", true, "Enable wildcard detection")
	flag.BoolVar(&options.Help, "h", false, "Show help message")

	flag.Parse()

	if options.Domain == "" {
		return nil, fmt.Errorf("target domain not specified")
	}

	if options.Wordlist == "" {
		return nil, fmt.Errorf("wordlist not specified")
	}

	return options, nil
}

func PrintUsage() {
	fmt.Println("Usage: puredns -d <domain> -w <wordlist> [options]")
	flag.PrintDefaults()
}
