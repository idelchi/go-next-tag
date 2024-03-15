# go-next-tag Overview

`go-next-tag` is a Go tool for automatically calculating and applying the next `major-minor` or `semantic`
version tag to your git repository.

## Features

- Automatically calculates the next version tag based on the current state of the repository.
- Supports semantic versioning and allows for major, minor, and patch version bumps.
- Customizable tag format and prefix settings.
- Option to push the newly created tag to the remote repository.
- Integrates with existing git configuration for user authentication.

## Getting Started

### Prerequisites

- Go 1.22 or higher
- Git

### Installation

Clone the repository and build the binary with:

    git clone ssh://git@code.swisscom.com:2222/swisscom/scsa-shared-tools/go-next-tag.git
    cd go-next-tag
    go build -o go-next-tag .

Alternatively, you can install it directly using:

    go install code.swisscom.com/swisscom/scsa-shared-tools/go-next-tag@latest

### Usage

Run `go-next-tag` with the desired flags. The available flags include:

    --version: Show the version information of go-next-tag.
    --token: Access token to authenticate to the git server. Defaults to GO_NEXT_TAG_TOKEN environment variable.
    --user-name: Username for git operations. Defaults to GO_NEXT_TAG_USER_NAME environment variable.
    --user-email: Email for git operations. Defaults to GO_NEXT_TAG_USER_EMAIL environment variable.
    --bump: Bump the next tag. Possible values: patch, minor, major, none. Default is 'patch'.
    --push: Push the tag to the remote repository. Default is false.
    --format: The format of the tag. Possible values: majorminor, semver. Default is 'majorminor'.
    --prefix: The prefix to use for the tag. Default is 'v'.

Example:

    go-next-tag \
        --token <your_access_token> \
        --user-name <your_username> \
        --user-email <your_email> \
        --bump minor \
        --format semver \
        --prefix v \
        --push

For more details on usage and configuration, run:

    go-next-tag --help

This will display a comprehensive list of flags and their descriptions.
