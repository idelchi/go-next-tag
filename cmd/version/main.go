package main

import (
	"fmt"
	"log"
	"unicode"

	"github.com/Masterminds/semver/v3"
)

func IsSemVer(v string) bool {
	_, err := semver.StrictNewVersion(stripFirstNonDigit(v))

	return err == nil
}

func IsMajorMinor(v string) bool {
	_, err := semver.StrictNewVersion(stripFirstNonDigit(v))

	return err != nil
}

func main() {
	testVersions := []string{
		// SemVer
		"1.2.3",
		"1.2.3-beta",
		"1.2.3+build.123",
		"1.2.3-beta+build.123",
		// SemVer
		"v1.2.3",
		"v1.2.3-beta",
		"v1.2.3+build.123",
		"v1.2.3-beta+build.123",
		// MajorMinor
		"1.2",
		"1.2-beta",
		"1.2+build.123",
		"1.2-beta+build.123",
		// MajorMinor
		"v1.2",
		"v1.2-beta",
		"v1.2+build.123",
		"v1.2-beta+build.123",
	}

	for _, v := range testVersions {
		if IsMajorMinor(v) {
			fmt.Printf("Version %q is a major.minor\n", v)
		} else {
			fmt.Printf("Version %q is a semver\n", v)
		}
	}
}

func main2() {
	testVersions := []string{
		// "1",
		// "1.2",
		"1.2.3",
		// "1.2.3-beta",
		// "1.2.3+build.123",
		// "1.2.3-beta+build.123",
		// "invalid.version", // Should fail
		// "1",
		// "v1.2",
		"v1.2.3",
		// "v1.2.3+build.123",
		// "v1.2.3-beta+build.123",
		// "vinvalid.version", // Should fail
	}

	for _, v := range testVersions {
		v = stripFirstNonDigit(v)

		version, err := semver.StrictNewVersion(v)
		if err != nil {
			log.Printf("Failed to parse version %q: %v\n", v, err)
			fmt.Println()
			continue
		}
		fmt.Printf("Successfully parsed version: %v -> %v\n", v, version)
		fmt.Println()
	}
}

func stripFirstNonDigit(s string) string {
	if len(s) > 0 && !unicode.IsDigit(rune(s[0])) {
		return s[1:]
	}
	return s
}
