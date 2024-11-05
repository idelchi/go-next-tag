# go-next-tag Overview

`go-next-tag` is a Go tool for calculating the next `major-minor` or `semantic` version given a tag.

## Installation

### From source

```sh
go install github.com/idelchi/go-next-tag@latest
```

### From installation script

```sh
curl -sSL https://raw.githubusercontent.com/idelchi/go-next-tag/refs/heads/dev/install.sh | sh -s -- -d ~/.local/bin
```

## Usage

```sh
go-next-tag [flags] [version]
```

Run `go-next-tag` with the desired flags. The available flags include:

```sh
--version: Show the version information of go-next-tag.
--bump: Bump the next tag. Possible values: patch, minor, major, none. Default is 'patch'.
--format: The format of the tag. Possible values: majorminor, semver, auto. Default is 'auto'.
```

Example:

```sh
go-next-tag --bump minor v1.2.3
```

```sh
echo "v1.2.3" | go-next-tag --bump minor
```

If the version is provided as both stdin and an argument, the argument will take precedence.

With `--format=auto`, it will be inferred as either `majorminor` or `semver` based on the input.
If no input is given, it will default to semver.

With `--bump`, you can specify the type of bump to apply to the version. The default is `patch`.
If the format is `majorminor`, selecting `patch` will be analogous to `minor`.

For more details on usage and configuration, run:

```sh
go-next-tag --help
```

This will display a comprehensive list of flags and their descriptions.

All flags can be set through environment variables. The prefix `GO_NEXT_TAG` is used to avoid conflicts.
For example, to set the bump strategy, use `GO_NEXT_TAG_BUMP`.
