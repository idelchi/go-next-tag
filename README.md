# go-next-tag

[![Go Reference](https://pkg.go.dev/badge/github.com/idelchi/go-next-tag.svg)](https://pkg.go.dev/github.com/idelchi/go-next-tag)
[![Go Report Card](https://goreportcard.com/badge/github.com/idelchi/go-next-tag)](https://goreportcard.com/report/github.com/idelchi/go-next-tag)
[![Build Status](https://github.com/idelchi/go-next-tag/actions/workflows/github-actions.yml/badge.svg)](https://github.com/idelchi/go-next-tag/actions/workflows/github-actions.yml/badge.svg)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

`go-next-tag` is a Go tool for calculating the next `major-minor` or `semantic` version given a tag.

## Installation

```sh
curl -sSL https://raw.githubusercontent.com/idelchi/go-next-tag/refs/heads/main/install.sh | sh -s -- -d ~/.local/bin
```

## Usage

```sh
go-next-tag [flags] [version|STDIN]
```

### Configuration

| Flag         | Environment Variable | Description                    | Default | Valid Values                      |
| ------------ | -------------------- | ------------------------------ | ------- | --------------------------------- |
| `--version`  | -                    | Show version                   | -       | -                                 |
| `--show`     | -                    | Show configuration             | -       | -                                 |
| `--bump`     | `GO_NEXT_TAG_BUMP`   | Version component to increment | `patch` | `patch`, `minor`, `major`, `none` |
| `--format`   | `GO_NEXT_TAG_FORMAT` | Version format to use          | `auto`  | `majorminor`, `semver`, `auto`    |
| `-h, --help` | -                    | Help for go-next-tag           | -       | -                                 |

### Format Types

| Format       | Description                              | Example  |
| ------------ | ---------------------------------------- | -------- |
| `majorminor` | Two-component version number             | `v1.2`   |
| `semver`     | Three-component semantic version         | `v1.2.3` |
| `auto`       | Inferred from input (defaults to semver) | -        |

### Examples

```sh
# Bump minor version of a semver tag
go-next-tag --bump minor v1.2.3
# Output: v1.3.0

# Read version from stdin
echo "v1.2.3" | go-next-tag --bump minor
# Output: v1.3.0

# Use majorminor format implicitly
go-next-tag v1.2
# Output: v1.3

# Use majorminor format explicitly
go-next-tag --format majorminor v1.2.3
# Output: v1.3
```

### Notes

- When version is provided as both stdin and argument, the argument takes precedence
- With `--format=auto`, format is inferred from input, defaulting to semver if no input
- When using `majorminor` format, `--bump patch` is equivalent to `--bump minor`

For detailed help:

```sh
go-next-tag --help
```
