// Package commands provides command-line flag configuration
// and parsing for determining version tag formats and bump rules.
package commands

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

// Flags defines the command-line flags for the application.
func Flags() {
	// General flags
	pflag.BoolP("version", "v", false, "Show the version information and exit")
	pflag.BoolP("help", "h", false, "Show the help information and exit")
	pflag.BoolP("show", "s", false, "Show the configuration and exit")

	// Format flags
	pflag.StringP(
		"bump",
		"b",
		"minor",
		`Bump the next tag. Possible values: 'patch', 'minor', 'major', 'none'.
If the format is 'majorminor', selecting patch will be analogous to 'minor'`,
	)
	pflag.StringP(
		"format",
		"f",
		"auto",
		`The format of the tag. Possible values: 'majorminor', 'semver' or 'auto'.
With auto, it will be inferred from the input.
If no input is given, it will default to semver`,
	)

	pflag.CommandLine.SortFlags = false

	pflag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: echo [version] | %s [flags] [version]\n\n", "go-next-tag")
		fmt.Fprintf(os.Stderr, "Generate the next tag based on the current tag and the specified bump rule.\n\n")
		fmt.Fprintf(
			os.Stderr,
			"If no version is provided as input, the tag will be generated as <prefix>0.0.0 or <prefix>0.0, "+
				"depending on the format.\n\n",
		)
		fmt.Fprintf(os.Stderr, "The version can be provided as a positional argument or via stdin.\n\n")
		fmt.Fprintf(os.Stderr, "The next tag will be printed to stdout.\n\n")
		fmt.Fprintf(os.Stderr, "Flags:\n")
		pflag.PrintDefaults()
	}
}
