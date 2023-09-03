package main

import (
	"os"
	"os/exec"
)

type GitClient interface {
	Clone(repoURL string, tmpDir string) error
}

type DefaultGitClient struct{}

func (c *DefaultGitClient) Clone(repoURL string, tmpDir string) error {
	// Clone the repository into the temporary directory.
	cmd := exec.Command("git", "clone", repoURL, tmpDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func ForkRepository(client GitClient) (string, error) {
	parentDir := "/tmp"
	repoURL := "git@github.com:KyleJGreen/mlctl-templates.git"

	// Create a temporary directory for the forked repository.
	tmpDir, err := os.MkdirTemp(parentDir, "forked-repo-")
	if err != nil {
		return "", err
	}

	if err := client.Clone(repoURL, tmpDir); err != nil {
		return "", err
	}

	return tmpDir, nil
}
