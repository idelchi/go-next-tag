// Package config defines and validates the configuration structure for version tag generation and formatting.
package config

import (
	"errors"
	"fmt"

	"github.com/idelchi/gogen/pkg/validator"
)

// ErrUsage indicates an error in command-line usage or configuration.
var ErrUsage = errors.New("usage error")

// Config represents the configuration for the go-next-tag application.
type Config struct {
	// The type of bump to perform
	Bump string `validate:"required,oneof=patch minor major none"`

	// The format of the tag
	Format string `validate:"required,oneof=majorminor semver auto"`

	// The prefix of the tag, if any
	Prefix string

	// The tag to compare to
	Tag string `validate:"version"`

	// Show the configuration and exit
	Show bool
}

// Validate performs configuration validation using the validator package.
// It returns a wrapped ErrUsage if any validation rules are violated.
func (c Config) Validate() error {
	validator := validator.NewValidator()

	if err := registerVersion(validator); err != nil {
		return fmt.Errorf("registering version: %w", err)
	}

	errs := validator.Validate(c)

	switch {
	case errs == nil:
		return nil
	case len(errs) == 1:
		return fmt.Errorf("%w: %w", ErrUsage, errs[0])
	case len(errs) > 1:
		return fmt.Errorf("%ws:\n%w", ErrUsage, errors.Join(errs...))
	}

	return nil
}
