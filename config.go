package main

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Config represents the configuration for the go-next-tag application.
type Config struct {
	User struct {
		Name  string `validate:"required"`       // The name to use for git operations
		Email string `validate:"required,email"` // The email to use for git operations
	}
	Token string `validate:"required" mask:"fixed"` // Access token to authenticate to the git server

	Dry bool // Dry run mode

	Action struct {
		// The type of bump to perform
		Bump string `validate:"required,oneof=patch minor major none"`
		// The format of the tag
		Format string `validate:"required,oneof=majorminor semver"`
		// The prefix to use for the tag
		Prefix string
		// Whether to checkout, and if so, the branch name and the commit hash, separated by a space
		Checkout string `validate:"omitempty,checkoutformat"`
	}

	Verbose bool   // Verbose mode
	Output  string `validate:"required,oneof=json text"` // Output mode (json or text)
}

// Validate validates the configuration against the struct tags.
func (c Config) Validate() error {
	validate := validator.New()

	// Register the custom rule "checkoutformat"
	err := validate.RegisterValidation("checkoutformat", validateCheckoutFormat)
	if err != nil {
		return fmt.Errorf("registering custom validation: %w", err)
	}

	// Validate the configuration
	if err := validate.Struct(c); err != nil {
		return fmt.Errorf("validating configuration: %w", err)
	}

	return nil
}

// validateCheckoutFormat validates the format of the checkout field.
// It expects the field to be either empty or in the format "branch commit".
func validateCheckoutFormat(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	if value == "" {
		return true
	}

	parts := strings.Split(value, " ")

	//nolint:mnd // We expect exactly two parts
	return len(parts) == 2
}
