package main

import (
	"log/slog"
	"os"
)

// ConfigureLogger sets up the verbosity and output mode of the logger.
// Supported log formats are "json" and "text".
// The function panics if an unknown log format is provided.
func ConfigureLogger(verbose bool, logFormat string) *slog.Logger {
	options := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	if verbose {
		options.Level = slog.LevelDebug
	}

	var handler slog.Handler

	switch logFormat {
	case "json":
		handler = slog.NewJSONHandler(os.Stdout, options)
	case "text":
		if !verbose {
			replace := func(_ []string, attr slog.Attr) slog.Attr {
				if attr.Key != slog.MessageKey {
					return slog.Attr{}
				}

				return attr
			}
			options.ReplaceAttr = replace
		}

		handler = slog.NewTextHandler(os.Stdout, options)
	default:
		panic("unknown log format")
	}

	return slog.New(handler)
}
