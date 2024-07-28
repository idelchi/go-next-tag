package versioning

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
)

// Next generates the next version based on the current version and the specified bump rule.
// Supports semantic versioning and a simplified <major.minor> format.
func Next(version string, bump string) (semver.Version, error) {
	if version == "" {
		defaultVersion, err := semver.NewVersion("0.0.0")
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

func ToFormat(version semver.Version, format string) string {
	switch format {
	case "majorminor":
		return fmt.Sprintf("%d.%d", version.Major(), version.Minor())
	default:
		return version.String()
	}
}