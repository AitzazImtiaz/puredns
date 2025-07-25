package dns

import (
	"net"
)

// IsA returns true if the record is an A record.
func IsA(record net.IP) bool {
	return record.To4() != nil
}

// IsCNAME returns true if the record is a CNAME record.
func IsCNAME(record string) bool {
	// Basic check, can be improved
	_, err := net.LookupHost(record)
	return err == nil
}
