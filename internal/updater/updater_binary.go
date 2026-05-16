package updater

import (
	"fmt"
	"os"
	"runtime"
)

// BinaryName returns the expected binary name for the current platform.
func BinaryName(appName, version string) string {
	return fmt.Sprintf("%s_%s_%s_%s", appName, version, runtime.GOOS, runtime.GOARCH)
}

// ReplaceBinary replaces the current binary with a new one.
func ReplaceBinary(newPath string) error {
	current, err := os.Executable()
	if err != nil {
		return fmt.Errorf("get current binary: %w", err)
	}
	return os.Rename(newPath, current)
}
