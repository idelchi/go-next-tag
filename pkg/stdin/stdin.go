// Package stdin provides simple utilities for reading from stdin.
package stdin

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// IsPiped checks if something has been piped to stdin.
func IsPiped() bool {
	fi, err := os.Stdin.Stat()

	return (fi.Mode()&os.ModeCharDevice) == 0 && err == nil
}

// MaybePiped checks if something has been piped to stdin.
func MaybePiped() (bool, error) {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false, fmt.Errorf("getting stdin stat: %w", err)
	}

	isPipe := (stat.Mode()&os.ModeNamedPipe) != 0 ||
		(stat.Mode()&(os.ModeCharDevice|os.ModeDir|os.ModeSymlink)) == 0

	return isPipe, nil
}

// Read returns stdin as a string, trimming the trailing newline.
func Read() (string, error) {
	bytes, err := io.ReadAll(os.Stdin)

	return strings.TrimSuffix(string(bytes), "\n"), err
}
