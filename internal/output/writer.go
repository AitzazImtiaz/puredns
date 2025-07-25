package output

import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/d3mondev/puredns/internal/brute"
)

// WriteResults writes results to a file in the specified format.
func WriteResults(results []brute.Result, path, format string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	switch format {
	case "json":
		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ")
		return encoder.Encode(results)
	default: // txt
		for _, result := range results {
			fmt.Fprintf(file, "%s %v\n", result.Subdomain, result.IPs)
		}
	}
	return nil
}
