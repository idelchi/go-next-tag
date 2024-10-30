package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/idelchi/go-next-tag/pkg/stdin"
	"github.com/idelchi/go-next-tag/pkg/versioning"
)

// flags defines the command-line flags for the application.
func flags() {
	// General flags
	pflag.BoolP("version", "v", false, "Show the version information and exit")
	pflag.BoolP("help", "h", false, "Show the help information and exit")
	pflag.BoolP("show", "s", false, "Show the configuration and exit")

	// Format flags
	pflag.StringP(
		"bump",
		"b",
		"patch",
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

// parseFlags parses the application configuration (in order of precedence) from:
//   - command-line flags
//   - environment variables
func parseFlags() (cfg Config, err error) {
	flags()

	// Parse the command-line flags
	pflag.Parse()

	// Bind pflag flags to viper
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		return cfg, fmt.Errorf("binding flags: %w", err)
	}

	// Set viper to automatically read from environment variables
	viper.SetEnvPrefix("go_next_tag")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// Unmarshal the configuration into the Config struct
	if err := viper.Unmarshal(&cfg); err != nil {
		return cfg, fmt.Errorf("unmarshalling config: %w", err)
	}

	// Validate the input
	err = validateInput(&cfg)

	// Handle the commandline flags that exit the application
	handleExitFlags(cfg)

	if err != nil {
		return cfg, fmt.Errorf("validating input: %w", err)
	}

	return cfg, nil
}

// validateInput validates the input configuration, selecting the tag from the command-line arguments or stdin.
func validateInput(cfg *Config) error {
	// Handle input source
	if pflag.NArg() > 0 {
		cfg.Tag = pflag.Arg(0)
	} else if stdin.IsPiped() {
		input, err := stdin.Read()
		if err != nil {
			return fmt.Errorf("reading input: %w", err)
		}

		cfg.Tag = input
	}

	cfg.Prefix = versioning.GetPrefix(cfg.Tag)
	cfg.Tag = versioning.StripPrefix(cfg.Tag)

	// Auto-detect format and set defaults
	if cfg.Format == "auto" {
		cfg.Format = "semver" // default
		if cfg.Tag != "" && !versioning.IsSemVerish(cfg.Tag) {
			cfg.Format = "majorminor"
		}
	}

	// Adjust bump type for majorminor format
	if cfg.Format == "majorminor" && cfg.Bump == "patch" {
		cfg.Bump = "minor"
	}

	return nil
}

//nolint:forbidigo // Function will print & exit for various help messages.
func handleExitFlags(cfg Config) {
	// Check if the version flag was provided
	if viper.GetBool("version") {
		fmt.Println(version)
		os.Exit(0)
	}

	// Check if the help flag was provided
	if viper.GetBool("help") {
		pflag.Usage()
		os.Exit(0)
	}

	if viper.GetBool("show") {
		fmt.Println(PrintJSON(cfg))

		os.Exit(0)
	}
}

// PrintJSON returns a pretty-printed JSON representation of the provided object.
func PrintJSON(obj any) string {
	bytes, err := json.MarshalIndent(obj, "  ", "    ")
	if err != nil {
		return err.Error()
	}

	return string(bytes)
}
