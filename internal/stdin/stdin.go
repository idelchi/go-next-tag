package stdin

import (
	"io"
	"os"
	"strings"
)

// IsPiped checks if stdin is piped.
func IsPiped() bool {
	fi, err := os.Stdin.Stat()

	return (fi.Mode()&os.ModeCharDevice) == 0 && err == nil
}

// Read returns stdin as a string, trimming the trailing newline.
func Read() (string, error) {
	bytes, err := io.ReadAll(os.Stdin)

	return strings.TrimSuffix(string(bytes), "\n"), err
}
