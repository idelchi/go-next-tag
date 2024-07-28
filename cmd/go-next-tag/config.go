package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// Config represents the configuration for the go-next-tag application.
type Config struct {
	// The type of bump to perform
	Bump string `validate:"required,oneof=patch minor major none"`
	// The format of the tag
	Format string `validate:"required,oneof=majorminor semver"`
	// The prefix to use for the tag
	Prefix string
	// The tag to compare to
	Tag string
}

// Validate validates the configuration against the struct tags.
func (c Config) Validate() error {
	validate := validator.New()

	// Validate the configuration
	if err := validate.Struct(c); err != nil {
		return fmt.Errorf("validating configuration: %w", err)
	}

	return nil
}
