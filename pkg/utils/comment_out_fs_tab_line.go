package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// CommentOutFstabLine comments out a line in /etc/fstab that matches the given pattern
func CommentOutFstabLine(pattern string) error {
	fstabFilePath := "/etc/fstab"

	// Open the /etc/fstab file for reading and writing
	file, err := os.OpenFile(fstabFilePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("failed to open %s: %w", fstabFilePath, err)
	}
	defer file.Close()

	// Read existing contents
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to read %s: %w", fstabFilePath, err)
	}

	// Check and comment out the matching line
	var updatedLines []string
	for _, line := range lines {
		if strings.Contains(line, pattern) {
			// Comment out the line
			line = "# " + line
		}
		updatedLines = append(updatedLines, line)
	}

	// Rewrite the file with the updated contents
	file.Seek(0, 0)  // Rewind to the start of the file
	file.Truncate(0) // Clear the file contents
	writer := bufio.NewWriter(file)
	for _, line := range updatedLines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("failed to write to %s: %w", fstabFilePath, err)
		}
	}
	if err := writer.Flush(); err != nil {
		return fmt.Errorf("failed to flush to %s: %w", fstabFilePath, err)
	}

	return nil
}
