# go-next-tag Overview

`go-next-tag` is a Go tool for calculating the next `major-minor` or `semantic` version given a tag.

## Getting Started

### Prerequisites

- Go 1.22 or higher
- Git

### Installation

    go install github.com/idelchi/cmd/go-next-tag.git@latest

### Usage

Run `go-next-tag` with the desired flags. The available flags include:

    --version: Show the version information of go-next-tag.
    --bump: Bump the next tag. Possible values: patch, minor, major, none. Default is 'patch'.
    --format: The format of the tag. Possible values: majorminor, semver. Default is 'majorminor'.
    --prefix: The prefix to use for the tag. Default is 'v'.

Example:

    go-next-tag --bump minor --format semver --prefix v v1.2.3

    echo "v1.2.3" | go-next-tag --bump minor --format semver --prefix v

For more details on usage and configuration, run:

    go-next-tag --help

This will display a comprehensive list of flags and their descriptions.

All flags can be set through environment variables. The prefix "NEXT*TAG*" is used to avoid conflicts. For example, to set the bump strategy, use `NEXT_TAG_BUMP`.
