package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// flags defines the command-line flags for the application.
func flags() {
	// General flags
	pflag.Bool("version", false, "Show the version information and exit")
	pflag.BoolP("help", "h", false, "Show the help information and exit")
	pflag.String("config", "", "Load configuration from `FILE`")
	pflag.BoolP("show", "s", false, "Show the configuration and exit")

	// User flags
	pflag.String("user-name", "", "Username to use for git operations")
	pflag.String("user-email", "", "Email to use for git operations")
	pflag.String("token", "", "Access token to authenticate to the git server")

	// Action flags
	pflag.String("bump", "patch", "Bump the next tag. Possible values: patch, minor, major, none")
	pflag.Bool("push", false, "Push the tag to the remote repository")
	pflag.String("format", "majorminor", "The format of the tag. Possible values: majorminor, semver")
	pflag.String("prefix", "v", "The prefix to use for the tag")
	pflag.String("checkout", "", "Checkout the branch name and the commit hash, separated by a space")

	// Global flags
	pflag.BoolP("verbose", "v", false, "Verbose mode")
	pflag.StringP("output", "o", "text", "Output mode (json or text)")

	pflag.CommandLine.SortFlags = false
}

// parseFlags parses the application configuration (in order of precedence) from:
//   - command-line flags
//   - environment variables
//   - configuration file
func parseFlags() (cfg Config, err error) {
	flags()

	// Parse the command-line flags
	pflag.Parse()

	// Bind pflag flags to viper
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		return cfg, fmt.Errorf("binding flags: %w", err)
	}

	// Set viper to automatically read from environment variables
	viper.SetEnvPrefix("next_tag")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// Only return errors if a configuration file was provided
	if viper.GetString("config") != "" {
		viper.SetConfigFile(viper.GetString("config"))

		if err := viper.ReadInConfig(); err != nil {
			return cfg, fmt.Errorf("reading config file: %w", err)
		}
	}

	// Unmarshal the configuration into the Config struct
	if err := viper.Unmarshal(&cfg); err != nil {
		return cfg, fmt.Errorf("unmarshaling config: %w", err)
	}

	handleExitFlags(cfg)

	return cfg, nil
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
		log.Println(cfg)
		os.Exit(0)
	}
}
