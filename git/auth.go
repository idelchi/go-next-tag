package gitmanager

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

func FindDefaultSSHKeyPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("getting user home directory: %w", err)
	}

	possiblePaths := []string{
		filepath.Join(homeDir, ".ssh", "id_rsa"),
		filepath.Join(homeDir, ".ssh", "id_ed25519"),
	}

	for _, path := range possiblePaths {
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}

	return "", fmt.Errorf("%w: no SSH key found in default locations %q", os.ErrNotExist, possiblePaths)
}

// AuthMethod is an interface for different authentication methods
type AuthMethod interface {
	Authenticate() (transport.AuthMethod, error)
}

// BasicAuthMethod implements AuthMethod for username/password authentication
type BasicAuthMethod struct {
	Username string
	Password string
}

func (b *BasicAuthMethod) Authenticate() (transport.AuthMethod, error) {
	return &http.BasicAuth{
		Username: b.Username,
		Password: b.Password,
	}, nil
}

// SSHAuthMethod implements AuthMethod for SSH key authentication
type SSHAuthMethod struct {
	PrivateKeyPath string
	Password       string
}

// Authenticate creates a new public key authentication method
func (s *SSHAuthMethod) Authenticate() (transport.AuthMethod, error) {
	publicKeys, err := ssh.NewPublicKeysFromFile("git", s.PrivateKeyPath, s.Password)
	if err != nil {
		return nil, fmt.Errorf("creating public keys: %w", err)
	}
	return publicKeys, nil
}
