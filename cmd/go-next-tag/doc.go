/*
go-next-tag is a simple command-line application for incrementing version tags.
It returns the next version based on the given version, format and bump strategy.

Usage:

	echo [version] | go-next-tag [flags] [version]

For full usage information, run `go-next-tag -h`.

`go-next-tag supports` configuration through environment variables starting with the prefix `GO_NEXT_TAG_`.
For example, to set the bump strategy, use `GO_NEXT_TAG_BUMP`.
*/
package main
