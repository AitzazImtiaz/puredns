package dns

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"time"
)

// Resolve performs DNS resolution for a given domain using a list of resolvers.
func Resolve(domain string, resolvers []string) ([]string, error) {
	var ips []string
	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(1000),
			}
			// Round-robin through resolvers
			resolverAddr := resolvers[rand.Intn(len(resolvers))]
			return d.DialContext(ctx, "udp", resolverAddr)
		},
	}

	addrs, err := resolver.LookupHost(context.Background(), domain)
	if err != nil {
		return nil, err
	}

	if len(addrs) == 0 {
		return nil, fmt.Errorf("no IP addresses found for %s", domain)
	}

	ips = append(ips, addrs...)
	return ips, nil
}
