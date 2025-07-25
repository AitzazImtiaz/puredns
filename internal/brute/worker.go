package brute

import (
	"fmt"
	"github.com/d3mondev/puredns/internal/dns"
)

func BruteForceWorker(
	jobs <-chan string,
	results chan<- Result,
	domain string,
	resolvers []string,
	wildcardIPs map[string]struct{},
) {
	for subdomain := range jobs {
		fqdn := fmt.Sprintf("%s.%s", subdomain, domain)
		ips, err := dns.Resolve(fqdn, resolvers)

		if err != nil {
			continue // Skip if resolution fails
		}

		// Filter out wildcard IPs
		var validIPs []string
		for _, ip := range ips {
			if _, isWildcard := wildcardIPs[ip]; !isWildcard {
				validIPs = append(validIPs, ip)
			}
		}

		if len(validIPs) > 0 {
			results <- Result{
				Subdomain: fqdn,
				IPs:       validIPs,
			}
		}
	}
}
