package utils

import "os/exec"

// isGitInstalled checks if Git is installed on the system.
func IsGitInstalled() bool {
	_, err := exec.LookPath("git")
	return err == nil
}

// isGitRepo checks if the current directory is a Git repository.
func IsGitRepo(dir string) bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	cmd.Dir = dir
	err := cmd.Run()
	return err == nil
}

// findGitRoot finds the root directory of the Git repository.
func FindGitRoot(startDir string) (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	cmd.Dir = startDir
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
