package utils

import (
	"strconv"
)

// Atoi converts string to int, returns 0 on error.
func Atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

// Atoi64 converts string to int64, returns 0 on error.
func Atoi64(s string) int64 {
	n, _ := strconv.ParseInt(s, 10, 64)
	return n
}

// Atob converts string to bool, returns false on error.
func Atob(s string) bool {
	b, _ := strconv.ParseBool(s)
	return b
}
