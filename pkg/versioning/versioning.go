// Package versioning provides utilities for calculating the next, converting and manipulating package versions.
package versioning

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/Masterminds/semver/v3"
)

// Next generates the next version based on the current version and the specified bump rule.
// Supports semantic versioning and a simplified <major.minor> format.
func Next(version, bump string) (semver.Version, error) {
	if version == "" {
		defaultVersion, err := ToSemVer("0.0.0")
		if err != nil {
			return semver.Version{}, fmt.Errorf("constructing default version: %w", err)
		}

		return *defaultVersion, nil
	}

	semverVersion, err := semver.NewVersion(version)
	if err != nil {
		return semver.Version{}, fmt.Errorf("parsing version: %w", err)
	}

	// Generate the next tag based on the bump rule and format.
	switch bump {
	case "none":
		return *semverVersion, nil
	case "patch":
		return semverVersion.IncPatch(), nil
	case "minor":
		return semverVersion.IncMinor(), nil
	case "major":
		return semverVersion.IncMajor(), nil
	default:
		return semver.Version{}, fmt.Errorf("%w: unexpected bump type %q", semver.ErrInvalidSemVer, bump)
	}
}

// ToFormat converts the version to the specified format.
// Currently only supporting either default or <major.minor> format.
func ToFormat(version semver.Version, format string) string {
	switch format {
	case "majorminor":
		return fmt.Sprintf("%d.%d", version.Major(), version.Minor())
	default:
		return version.String()
	}
}

// ToSemVer converts the version string to a semantic version.
func ToSemVer(version string) (*semver.Version, error) {
	semVer, err := semver.NewVersion(version)
	if err != nil {
		return nil, fmt.Errorf("parsing version: %w", err)
	}

	return semVer, nil
}

// getFirstNonDigits returns all leading non-digit characters from the string.
func getFirstNonDigits(s string) string {
	for i, char := range s {
		if unicode.IsDigit(char) {
			return s[:i]
		}
	}

	if len(s) > 0 {
		return s
	}

	return ""
}

// GetPrefix returns all leading non-digit characters of the version string.
func GetPrefix(s string) string {
	return getFirstNonDigits(s)
}

// StripPrefix removes all leading non-digit characters from the string.
func StripPrefix(s string) string {
	return strings.TrimPrefix(s, getFirstNonDigits(s))
}

// IsSemVerish checks if the version string is semver-like.
func IsSemVerish(version string) bool {
	_, err := semver.StrictNewVersion(version)

	return err == nil
}
