package utils

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"
)

// GitNotFound indicates git binary is not installed in the machine
var ErrGitNotFound = errors.New("git command not found")

// cloneTemplate clones the service template repository
// returns an error if git not installed in the machine
func CloneTemplate(template string) (string, error) {
	tempDir := fmt.Sprint(os.TempDir(), "service-template-", time.Now().UnixNano())
	gitPath, err := exec.LookPath("git")

	if err != nil {
		return tempDir, ErrGitNotFound
	}

	clone := &exec.Cmd{
		Path:   gitPath,
		Args:   []string{gitPath, "clone", template, tempDir},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	err = clone.Run()
	return tempDir, err
}
