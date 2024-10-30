package main

import (
	"fmt"
	"unicode"

	"github.com/Masterminds/semver/v3"
)

// getFirstNonDigits returns all leading non-digit characters until we hit the actual version part.
func getFirstNonDigits(versionWithPrefix string) string {
	for i := range len(versionWithPrefix) {
		candidate := versionWithPrefix[i:]
		// First check if it starts with a digit
		if len(candidate) > 0 && !unicode.IsDigit(rune(candidate[0])) {
			continue
		}
		// Then check if it's valid semver
		if _, err := semver.NewVersion(candidate); err == nil {
			return versionWithPrefix[:i]
		}
	}

	return ""
}

func main() {
	strings := []string{
		"v1.2.3",
		"v1.2",
		"v1",
		"1.2.3",
		"1.2",
		"1",
		// More complex with long prefixes
		"v1.2.3-rc1",
		"v1.2-rc1",
		"v1-rc1",
		"prefix-v1.2.3",
		"prefix-v1.2",
		"prefix-v1",
		"prefix123withnumbers-v1.2.3",
		"prefix123withnumbers-v1.2",
		"prefix123withnumbers-v1",
		"prefix123withnumbers-v1.2.3-beta1",
		"prefix123withnumbers-v1.2-beta1",
		"prefix123withnumbers-v1-beta1",
		"invalidversionwith-v1xyu",
	}

	for _, s := range strings {
		prefix := getFirstNonDigits(s)
		fmt.Printf("Prefix of %q is %q\n", s, prefix)
	}
}
