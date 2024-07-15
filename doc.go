/*
go-next-tag is a simple command-line application for managing version tagging in a git repository.
It allows for the automatic calculation of the next tag based on the current repository state and the bump strategy.

Usage:

	go-next-tag [flags]

For full usage information, run `go-next-tag -h`.

Environment Variables:

	go-next-tag supports configuration through environment variables.
	Prefix "NEXT_TAG_" is used to avoid conflicts. For example, to set the user name, use NEXT_TAG_USER_NAME.
*/
package main
