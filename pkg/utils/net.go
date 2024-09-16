package utils

import "net"

func IsValidIP(ip string) bool {
	// ParseIP returns nil if the ip string is invalid
	parsedIP := net.ParseIP(ip)
	return parsedIP != nil
}
