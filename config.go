package main

import (
	"errors"
	"fmt"
	"log"
	"slices"
	"strings"
)

// User represents the user to use for git operations.
type User struct {
	Name  string // The name to use for git operations
	Email string // The email to use for git operations
}

// Credentials represents the credentials to use for git operations.
type Credentials struct {
	User  string // The username to use for git operations
	Token string // The token to use for git operations
}

// Action represents the action to perform on the git repository.
type Action struct {
	Bump   string // The type of bump to perform
	Push   bool   // Whether to push the tag to the remote repository
	Format string // The format of the tag
	Prefix string // The prefix to use for the tag
}

// ErrActionValidation is the error returned when an action is not valid.
var ErrActionValidation = errors.New("action validation failed")

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
