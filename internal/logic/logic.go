// Package logic implements the core execution flow for generating the next version tag.
package logic

import (
	"fmt"

	"github.com/idelchi/go-next-tag/internal/parse"
	"github.com/idelchi/go-next-tag/internal/versioning"
)

// Run executes the core logic for generating the next version tag.
func Run(version string) error {
	cfg, err := parse.Parse(version)
	if err != nil {
		return fmt.Errorf("parsing flags: %w", err)
	}

	if err := cfg.Validate(); err != nil {
		return fmt.Errorf("application configuration: %w", err)
	}

	nextTag, err := versioning.Next(cfg.Tag, cfg.Bump)
	if err != nil {
		return fmt.Errorf("calculating next tag: %w", err)
	}

	fmt.Println( //nolint:forbidigo // Print the next tag to stdout
		cfg.Prefix + versioning.ToFormat(nextTag, cfg.Format),
	)

	return nil
}
