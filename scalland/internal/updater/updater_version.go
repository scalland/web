package updater

import (
	"os"
	"strings"
)

// ReadVersionFile reads the VERSION file from the project root.
func ReadVersionFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}
