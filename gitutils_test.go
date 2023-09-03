package main

import (
	"errors"
	"testing"
)

type MockGitClient struct {
	Err error
}

func (m *MockGitClient) Clone(repoURL string, tmpDir string) error {
	return m.Err
}

func TestForkRepository(t *testing.T) {
	mockClient := &MockGitClient{Err: errors.New("Unable to fetch respository")}

	// Test error handling with call to mock git client
	_, err := ForkRepository(mockClient)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	} else if err.Error() != "Unable to fetch respository" {
		t.Errorf("Expected Unable to fetch respository' error, but got: %v", err)
	}

}
