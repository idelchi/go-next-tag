// go-next-tag is a simple command-line application for managing version tagging in a git repository.
// Allows for the automatic calculation of the next tag based on the current repository state and the bump strategy.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// Global variable for CI stamping.
var version = "unknown - unofficial & generated by unknown"

func main() { //nolint: funlen,cyclop
	user := User{}
	credentials := Credentials{}
	action := Action{}

	allowedBumps := []string{"patch", "minor", "major", "none"}
	allowedFormats := []string{"majorminor", "semver"}

	// Define command-line flags for the application.
	versionFlag := flag.Bool("version", false, "Show the version information")
	flag.StringVar(
		&credentials.Token,
		"token",
		os.Getenv("GO_NEXT_TAG_TOKEN"),
		"Access token to authenticate to the git server.",
	)
	flag.StringVar(&user.Name, "user-name", os.Getenv("GO_NEXT_TAG_USER_NAME"), "Username to use for git operations.")
	flag.StringVar(&user.Email, "user-email", os.Getenv("GO_NEXT_TAG_USER_EMAIL"), "Email to use for git operations.")
	flag.StringVar(&action.Bump, "bump", "patch", fmt.Sprintf("Bump the next tag. Possible values: %v", allowedBumps))
	flag.BoolVar(&action.Push, "push", false, "Push the tag to the remote repository.")
	flag.StringVar(
		&action.Format,
		"format",
		"majorminor",
		fmt.Sprintf("The format of the tag. Possible values: %v", allowedFormats),
	)
	flag.StringVar(&action.Prefix, "prefix", "v", "The prefix to use for the tag")

	// Custom usage function to provide more detailed help text.
	flag.Usage = func() {
		fmt.Println("Usage: go-next-tag [flags]")
		fmt.Println("Flags:")
		flag.PrintDefaults()
	}

	// Parse the command-line flags.
	flag.Parse()

	// If the version flag is provided, print the version and exit.
	if *versionFlag {
		fmt.Println(version)
		os.Exit(0)
	}

	if err := action.validate(allowedBumps, allowedFormats); err != nil {
		log.Fatalf("error validating action: %s", err)
	}

	// Initialize a new GitManager with the provided flags.
	gitManager, err := NewGitManager(credentials.Token)
	if err != nil {
		log.Fatalf("error initializing GitManager: %s", err)
	}

	// Configure git with the provided settings.
	if err := gitManager.ConfigureGit(user); err != nil {
		log.Fatalf("error configuring Git: %s", err)
	}

	// Calculate the next tag based on the current repository state and the bump strategy.
	tag, err := gitManager.CalculateNextTag(action.Bump)
	if err != nil {
		log.Fatalf("error calculating next tag: %s", err)
	}

	var tagString string

	switch action.Format {
	case "semver":
		tagString = fmt.Sprintf("%s%s", action.Prefix, tag.String())
	case "majorminor":
		tagString = fmt.Sprintf("%s%d.%d", action.Prefix, tag.Major(), tag.Minor())
	}

	log.Printf("Next tag: %s\n", tagString)

	// Create the calculated tag.
	if err := gitManager.CreateTag(tagString); err != nil {
		log.Fatalf("error creating tag: %s", err)
	}

	// If the push flag is not set, exit.
	if !action.Push {
		os.Exit(0)
	}

	// Push the tag to the remote repository.
	if err := gitManager.PushTag(tagString); err != nil {
		log.Fatalf("error pushing tag: %s", err)
	}

	log.Printf("Tag %q pushed successfully.\n", tagString)
}
