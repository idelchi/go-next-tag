package main

import (
	"errors"
	"fmt"
	"log"
	"slices"
	"strings"
)

// User represents the user to use for git operations.
// The details will be embedded in the git configuration when pushing the tag.
type User struct {
	Name  string // The name to use for git operations
	Email string // The email to use for git operations
}

// Credentials represents the credentials to use for git operations.
// Allows for the use token-based authentication.
type Credentials struct {
	User  string // The username to use for git operations
	Token string // The token to use for git operations
}

// Action represents the action to perform on the git repository, allowing specification of
// the type of bump to perform, whether to push the tag to the remote repository, the versioning format,
// and the prefix to use for the tag (e.g. "v").
type Action struct {
	Bump   string // The type of bump to perform
	Push   bool   // Whether to push the tag to the remote repository
	Format string // The format of the tag
	Prefix string // The prefix to use for the tag
}

// ErrActionValidation is the error returned when an invalid action is requested.
var ErrActionValidation = errors.New("action validation failed")

// validate that the bump and formats are within the allowed values.
// Automatically increments the minor version if the format is "majorminor" and the bump is "patch".
func (a *Action) validate(allowedBumps, allowedFormats []string) error {
	if !slices.Contains(allowedBumps, a.Bump) {
		return fmt.Errorf(
			"%w: invalid bump: %q, allowed values: %s",
			ErrActionValidation,
			a.Bump,
			strings.Join(allowedBumps, ", "),
		)
	}

	if !slices.Contains(allowedFormats, a.Format) {
		return fmt.Errorf(
			"%w: invalid format: %q, allowed values: %s",
			ErrActionValidation,
			a.Format,
			strings.Join(allowedFormats, ", "),
		)
	}

	if a.Format == "majorminor" && a.Bump == "patch" {
		log.Println("incrementing minor version as format is 'majorminor'")

		a.Bump = "minor"
	}

	return nil
}
