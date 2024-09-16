package utils

import "os"

// HasRootPrivileges checks if the current process is running as root (UID 0).
func HasRootPrivileges() bool {
	return os.Geteuid() == 0
}
