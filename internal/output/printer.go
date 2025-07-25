package output

import (
	"fmt"
	"github.com/d3mondev/puredns/internal/brute"
)

// PrintResult prints a single result to the console.
func PrintResult(result brute.Result) {
	fmt.Printf("\033[32mFound:\033[0m %s - %v\n", result.Subdomain, result.IPs)
}
