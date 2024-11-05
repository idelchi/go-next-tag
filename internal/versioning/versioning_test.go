package versioning_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/idelchi/go-next-tag/pkg/versioning"
)

func TestNextMajorMinor(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name        string
		version     string
		bump        string
		expected    string
		expectError bool
	}{
		{"Empty version", "", "patch", "0.0.0", false},
		{"Zero version", "0.0", "patch", "0.0.1", false},
		{"Patch bump", "1.2.3", "patch", "1.2.4", false},
		{"Minor bump", "1.2.3", "minor", "1.3.0", false},
		{"Major bump", "1.2.3", "major", "2.0.0", false},
		{"No bump", "1.2.3", "none", "1.2.3", false},
		{"Invalid version", "invalid", "patch", "", true},
		{"Invalid bump", "1.2.3", "invalid", "", true},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			result, err := versioning.Next(testCase.version, testCase.bump)

			if testCase.expectError {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, testCase.expected, result.String())
			}
		})
	}
}

func TestToFormat(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		version  string
		format   string
		expected string
	}{
		{"Default format", "1.2.3", "", "1.2.3"},
		{"Major.Minor format", "1.2.3", "majorminor", "1.2"},
		{"Default format with pre-release", "1.2.3-alpha", "", "1.2.3-alpha"},
		{"Major.Minor format with pre-release", "1.2.3-alpha", "majorminor", "1.2"},
		{"Default format with build metadata", "1.2.3+build.1", "", "1.2.3+build.1"},
		{"Major.Minor format with build metadata", "1.2.3+build.1", "majorminor", "1.2"},
		{"Unknown format", "1.2.3", "unknown", "1.2.3"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			v, err := semver.NewVersion(testCase.version)
			require.NoError(t, err)

			result := versioning.ToFormat(*v, testCase.format)
			assert.Equal(t, testCase.expected, result)
		})
	}
}
