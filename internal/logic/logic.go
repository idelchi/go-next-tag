package logic

import (
	"fmt"

	"github.com/idelchi/go-next-tag/internal/parse"
	"github.com/idelchi/go-next-tag/internal/versioning"
)

func Run(version string) error {
	cfg, err := parse.Parse()
	if err != nil {
		return fmt.Errorf("parsing flags: %w", err)
	}

	if err := cfg.Validate(); err != nil {
		return fmt.Errorf("validating configuration: %w", err)
	}

	nextTag, err := versioning.Next(cfg.Tag, cfg.Bump)
	if err != nil {
		return fmt.Errorf("calculating next tag: %w", err)
	}

	fmt.Println(cfg.Prefix + versioning.ToFormat(nextTag, cfg.Format))

	return nil
}
