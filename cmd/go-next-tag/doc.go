/*
go-next-tag is a simple command-line application for incrementing version tags.
It allows for the automatic calculation of the next tag based on the given tag, format and bump strategy.

Usage:

	echo [version] | go-next-tag [flags] [version]

For full usage information, run `go-next-tag -h`.

Environment Variables:

	go-next-tag supports configuration through environment variables.
	Prefix "NEXT_TAG_" is used to avoid conflicts. For example, to set the bump strategy, use `NEXT_TAG_BUMP`.
*/
package main
