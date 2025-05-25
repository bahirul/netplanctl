package validation

import (
	"fmt"
	"net"
	"strings"
)

func ValidateIP(cidr string) error {
	if cidr == "" {
		return fmt.Errorf("ip address cannot be empty")
	}

	if !strings.Contains(cidr, "/") {
		return fmt.Errorf("must be in CIDR format (e.g. 192.168.1.1/24)")
	}

	// valid ipv4 or ipv6 address
	ip := strings.Split(cidr, "/")[0]
	if net.ParseIP(ip) == nil {
		return fmt.Errorf("invalid IP address: %s", ip)
	}

	// valid CIDR notation
	_, _, err := net.ParseCIDR(cidr)
	if err != nil {
		return fmt.Errorf("invalid CIDR format: %w", err)
	}

	return nil
}
