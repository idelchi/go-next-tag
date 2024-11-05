package config

import (
	"fmt"

	versioning "github.com/idelchi/go-next-tag/internal/versioning"
	"github.com/idelchi/gogen/pkg/validator"
)

// registerExclusive adds a custom validator ensuring two fields are mutually exclusive.
// It registers both the validation logic and a human-readable error message.
func registerVersion(validator *validator.Validator) error {
	// Register the exclusive validation
	if err := validator.RegisterValidationAndTranslation(
		"version",
		validateSemVer,
		"{0} is not a valid semver-compatible version",
	); err != nil {
		return fmt.Errorf("registering version validation: %w", err)
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
