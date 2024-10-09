package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"

	versioning "github.com/idelchi/go-next-tag/pkg/versioning"
)

// Config represents the configuration for the go-next-tag application.
type Config struct {
	// The type of bump to perform
	Bump string `validate:"required,oneof=patch minor major none"`
	// The format of the tag
	Format string `validate:"required,oneof=majorminor semver"`
	// The prefix to use for the tag (e.g. "v")
	Prefix string
	// The tag to compare to
	Tag string `validate:"version"`
}

// Validate validates the configuration against the struct tags.
func (c Config) Validate() error {
	validate := validator.New()

	// Register the custom rule "version"
	err := validate.RegisterValidation("version", validateSemVer)
	if err != nil {
		return fmt.Errorf("registering custom validation: %w", err)
	}

	// Validate the configuration
	if err := validate.Struct(c); err != nil {
		return fmt.Errorf("validating configuration: %w", err)
	}

	return nil
}

// validateSemVer validates the format of the `Tag` field.
// It expects the field to be either semver-compatible or empty.
func validateSemVer(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	if value == "" {
		return true
	}

	_, err := versioning.ToSemVer(value)

	return err == nil
}
