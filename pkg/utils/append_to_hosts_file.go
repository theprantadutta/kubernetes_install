package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// AppendToHostsFile appends an IP and hostname entry to /etc/hosts
func AppendToHostsFile(ip, hostname string) error {
	hostsFilePath := "/etc/hosts"

	// Open the /etc/hosts file for reading
	file, err := os.OpenFile(hostsFilePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("failed to open %s: %w", hostsFilePath, err)
	}
	defer file.Close()

	// Read existing contents
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to read %s: %w", hostsFilePath, err)
	}

	// Check if the entry already exists
	entry := fmt.Sprintf("%s\t%s", ip, hostname)
	for _, line := range lines {
		if strings.HasPrefix(line, ip) || strings.Contains(line, hostname) {
			// Entry already exists
			return nil
		}
	}

	// Append the new entry
	lines = append(lines, entry)

	// Rewrite the file with the new contents
	file.Seek(0, 0) // Rewind to the start of the file
	file.Truncate(0) // Clear the file contents
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("failed to write to %s: %w", hostsFilePath, err)
		}
	}
	if err := writer.Flush(); err != nil {
		return fmt.Errorf("failed to flush to %s: %w", hostsFilePath, err)
	}

	return nil
}
