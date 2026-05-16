package utils

import (
	"io"
	"os"
	"path/filepath"
)

// FileExists returns true if the file at path exists.
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// EnsureDir creates a directory and all parents if they don't exist.
func EnsureDir(path string) error {
	return os.MkdirAll(path, 0755)
}

// CopyFile copies src to dst.
func CopyFile(src, dst string) error {
	if err := EnsureDir(filepath.Dir(dst)); err != nil {
		return err
	}
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	return err
}

// ReadFileString reads a file and returns its content as a string.
func ReadFileString(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// WriteFileString writes a string to a file.
func WriteFileString(path, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}
