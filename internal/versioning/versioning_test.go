package versioning_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/idelchi/go-next-tag/internal/versioning"
	"github.com/stretchr/testify/assert"
)

func TestNextMajorMinor(t *testing.T) {
	tests := []struct {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := versioning.Next(tt.version, tt.bump)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result.String())
			}
		})
	}
}

func TestToFormat(t *testing.T) {
	tests := []struct {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := semver.NewVersion(tt.version)
			assert.NoError(t, err)

			result := versioning.ToFormat(*v, tt.format)
			assert.Equal(t, tt.expected, result)
		})
	}
}
