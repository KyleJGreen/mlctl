package main

import (
	"os"
	"os/exec"
)

func ForkRepository() (string, error) {
	parentDir := "/tmp"

	// Create a temporary directory for the forked repository.
	tmpDir, err := os.MkdirTemp(parentDir, "forked-repo-")
	repoURL := "git@github.com:KyleJGreen/mlctl-templates.git"
	if err != nil {
		return "", err
	}

	// Clone the repository into the temporary directory.
	cmd := exec.Command("git", "clone", repoURL, tmpDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return "", err
	}

	return tmpDir, nil
}
