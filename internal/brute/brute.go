package brute

import (
	"fmt"
	"github.com/d3mondev/puredns/internal/dns"
	"github.com/d3mondev/puredns/internal/input"
)

func Run(domain, wordlistPath, resolversPath string, concurrency int) (<-chan Result, error) {
	wordlist, err := input.ReadWordlist(wordlistPath)
	if err != nil {
		return nil, fmt.Errorf("could not read wordlist: %w", err)
	}

	resolvers, err := input.ReadResolvers(resolversPath)
	if err != nil {
		return nil, fmt.Errorf("could not read resolvers: %w", err)
	}

	wildcardIPs, err := dns.DetectWildcard(domain, resolvers)
	if err != nil {
		return nil, fmt.Errorf("could not detect wildcard: %w", err)
	}

	results := make(chan Result)
	jobs := make(chan string)

	for i := 0; i < concurrency; i++ {
		go BruteForceWorker(jobs, results, domain, resolvers, wildcardIPs)
	}

	go func() {
		defer close(jobs)
		for _, subdomain := range wordlist {
			jobs <- subdomain
		}
	}()

	return results, nil
}
