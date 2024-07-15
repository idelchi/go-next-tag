package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

// GitManager encapsulates the operations related to managing Git repositories,
// including tagging and configuration.
type GitManager struct {
	repo *git.Repository // repo represents the Git repository managed by this GitManager.

	auth *http.BasicAuth
}

// NewGitManager initializes a new GitManager by opening the current Git repository.
// It returns a pointer to the GitManager and any error encountered.
func NewGitManager(token string) (gm *GitManager, err error) {
	gm = &GitManager{}

	// Attempt to open the current directory as a Git repository.
	gm.repo, err = git.PlainOpen(".")
	if err != nil {
		return nil, fmt.Errorf("opening repository at '.': %w", err)
	}

	gm.auth = &http.BasicAuth{
		Username: "user",
		Password: token,
	}

	return gm, nil
}

// ConfigureGit sets the username and email for Git operations within the repository.
// It applies the configuration globally for the repository.
func (gm *GitManager) ConfigureGit(user, email string) error {
	cfg, err := gm.repo.Config()
	if err != nil {
		return fmt.Errorf("getting configuration: %w", err)
	}

	// Set user name and email in the Git configuration.
	cfg.Raw.Section("user").SetOption("name", user)
	cfg.Raw.Section("user").SetOption("email", email)

	if err := gm.repo.Storer.SetConfig(cfg); err != nil {
		return fmt.Errorf("setting configuration: %w", err)
	}

	return nil
}

// Checkout switches the current working directory to the specified branch or commit.
func (gm *GitManager) Checkout(branch, commit string) error {
	// Fetch the latest changes from the remote repository
	err := gm.repo.Fetch(&git.FetchOptions{
		Auth: gm.auth,
	})
	if err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
		return fmt.Errorf("fetching changes: %w", err)
	}

	// Branch does not exist, attempt to checkout as a new branch
	workTree, err := gm.repo.Worktree()
	if err != nil {
		return fmt.Errorf("getting worktree: %w", err)
	}

	// Checkout options with Create set to true to allow creating a new branch
	opts := &git.CheckoutOptions{
		Force:  true,                                    // Force checkout, discarding local changes (if necessary)
		Branch: plumbing.NewBranchReferenceName(branch), // The branch to checkout or create
		Hash:   plumbing.NewHash(commit),                // The commit hash to checkout
	}

	if err = workTree.Checkout(opts); err != nil {
		return fmt.Errorf("checking out branch and commit: %w", err)
	}

	return nil
}

// GetLatestTag retrieves the latest tag from the repository and returns it as a semver.Version.
// It fetches tags from the remote repository and identifies the latest one.
func (gm *GitManager) GetLatestTag() (*semver.Version, error) {
	// Fetch tags from the remote repository.
	if err := gm.repo.Fetch(&git.FetchOptions{}); err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
		return nil, fmt.Errorf("fetching tags: %w", err)
	}

	tags, err := gm.repo.Tags()
	if err != nil {
		return nil, fmt.Errorf("getting tags: %w", err)
	}

	var latestVersion *semver.Version
	// Iterate over tags to find the latest.
	err = tags.ForEach(func(t *plumbing.Reference) error {
		version, err := semver.NewVersion(t.Name().Short())
		if err != nil {
			return fmt.Errorf("parsing tag %q: %w", t.Name().Short(), err)
		}

		if latestVersion == nil || version.GreaterThan(latestVersion) {
			latestVersion = version
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("iterating over tags: %w", err)
	}

	if latestVersion == nil {
		return &semver.Version{}, nil
	}

	return latestVersion, nil
}

// CalculateNextTag generates the next tag based on the current latest tag and the specified bump rule.
// It supports semantic versioning and a simplified <major.minor> format.
func (gm *GitManager) CalculateNextTag(bump string) (semver.Version, error) {
	latest, err := gm.GetLatestTag()
	if err != nil {
		return semver.Version{}, err
	}

	if latest == nil {
		defaultVersion, err := semver.NewVersion("0.0.0")
		if err != nil {
			return semver.Version{}, fmt.Errorf("constructing default version: %w", err)
		}

		return *defaultVersion, nil
	}

	// Generate the next tag based on the bump rule and format.
	switch bump {
	case "none":
		return *latest, nil
	case "patch":
		return latest.IncPatch(), nil
	case "minor":
		return latest.IncMinor(), nil
	case "major":
		return latest.IncMajor(), nil
	default:
		return semver.Version{}, fmt.Errorf("%w: unexpected bump type %q", semver.ErrInvalidSemVer, bump)
	}
}

// CreateTag creates a new tag in the local repository.
func (gm *GitManager) CreateTag(tag string) error {
	head, err := gm.repo.Head()
	if err != nil {
		return fmt.Errorf("getting HEAD: %w", err)
	}

	cfg, err := gm.repo.Config()
	if err != nil {
		return fmt.Errorf("getting configuration: %w", err)
	}

	_, err = gm.repo.CreateTag(tag, head.Hash(), &git.CreateTagOptions{
		Tagger: &object.Signature{
			Name:  cfg.Raw.Section("user").Option("name"),
			Email: cfg.Raw.Section("user").Option("email"),
			When:  time.Now(),
		},
		Message: "Automatic tagging",
	})
	if err != nil {
		return fmt.Errorf("creating tag: %w", err)
	}

	return nil
}

// PushTag pushes the specified tag to the remote repository.
func (gm *GitManager) PushTag(tag string) error {
	remote, err := gm.repo.Remote("origin")
	if err != nil {
		return fmt.Errorf("getting remote: %w", err)
	}

	err = remote.Push(&git.PushOptions{
		FollowTags: true,
		RemoteName: "origin",
		RefSpecs: []config.RefSpec{
			config.RefSpec("refs/tags/" + tag + ":refs/tags/" + tag),
		},
		Auth: gm.auth,
	})
	if err != nil {
		return fmt.Errorf("pushing tag: %w", err)
	}

	return nil
}
