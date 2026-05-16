package utils

import (
	"fmt"
	"strings"
)

// JoinInts joins a slice of ints with a separator.
func JoinInts(ints []int, sep string) string {
	s := make([]string, len(ints))
	for i, v := range ints {
		s[i] = fmt.Sprintf("%d", v)
	}
	return strings.Join(s, sep)
}

// JoinInt64s joins a slice of int64s with a separator.
func JoinInt64s(ints []int64, sep string) string {
	s := make([]string, len(ints))
	for i, v := range ints {
		s[i] = fmt.Sprintf("%d", v)
	}
	return strings.Join(s, sep)
}
