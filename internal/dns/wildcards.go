package dns

import (
	"fmt"
	"math/rand"
	"time"
)

// DetectWildcard performs wildcard detection on a domain.
func DetectWildcard(domain string, resolvers []string) (map[string]struct{}, error) {
	wildcardIPs := make(map[string]struct{})
	for i := 0; i < 5; i++ {
		// Generate a random subdomain that is unlikely to exist
		subdomain := fmt.Sprintf("%d-puredns-wildcard-check.%s", rand.Intn(1000000), domain)
		ips, err := Resolve(subdomain, resolvers)
		if err == nil {
			for _, ip := range ips {
				wildcardIPs[ip] = struct{}{}
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
	return wildcardIPs, nil
}
