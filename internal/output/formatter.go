package output

import (
	"encoding/json"
	"fmt"
	"github.com/d3mondev/puredns/internal/brute"
)

// Format formats the results based on the specified format.
func Format(results []brute.Result, format string) (string, error) {
	switch format {
	case "json":
		data, err := json.MarshalIndent(results, "", "  ")
		if err != nil {
			return "", err
		}
		return string(data), nil
	default: // txt
		var output string
		for _, result := range results {
			output += fmt.Sprintf("%s %v\n", result.Subdomain, result.IPs)
		}
		return output, nil
	}
}
