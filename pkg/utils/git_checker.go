package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetGitRoot(startDir string) (string, error) {
	gitInstalled := isGitInstalled()

	if !gitInstalled {
		return "", fmt.Errorf("git is not installed")
	}

	isGitRepo := isGitRepo(startDir)

	if !isGitRepo {
		return "", fmt.Errorf("not running inside git repo")
	}

	gitRoot, err := findGitRoot(startDir)

	if err != nil {
		return "", fmt.Errorf("can't find git root")
	}

	return gitRoot, nil
}

func isGitInstalled() bool {
	_, err := exec.LookPath("git")
	return err == nil
}

func isGitRepo(dir string) bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	cmd.Dir = dir
	err := cmd.Run()
	return err == nil
}

func findGitRoot(startDir string) (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	cmd.Dir = startDir
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}
