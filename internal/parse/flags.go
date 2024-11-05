// Package parse handles configuration parsing from command-line flags
// and environment variables, with input validation and format detection.
package parse

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	flags "github.com/idelchi/go-next-tag/internal/commands"
	"github.com/idelchi/go-next-tag/internal/config"
	"github.com/idelchi/go-next-tag/internal/versioning"
	"github.com/idelchi/godyl/pkg/pretty"
	"github.com/idelchi/gogen/pkg/stdin"
)

// Parse parses the application configuration (in order of precedence) from:
//   - command-line flags
//   - environment variables
func Parse(version string) (cfg config.Config, err error) {
	flags.Flags()

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
	handleExitFlags(cfg, version)

	if err != nil {
		return cfg, fmt.Errorf("validating input: %w", err)
	}

	return cfg, nil
}

//nolint:forbidigo // Function will print & exit for various help messages.
func handleExitFlags(cfg config.Config, version string) {
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
		pretty.PrintJSON(cfg)

		os.Exit(0)
	}
}

// validateInput validates the input configuration, selecting the tag from the command-line arguments or stdin.
func validateInput(cfg *config.Config) error {
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
