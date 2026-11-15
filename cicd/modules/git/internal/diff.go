package internal

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

func (s *service) handleListing(comparedCommit string) ([]string, error) {
	args := []string{"--no-pager", "diff", "--name-only", "--staged"}
	if comparedCommit != "" {
		args = append(args, comparedCommit)
	} else {
		args = append(args, "HEAD")
	}
	cmd := exec.Command("git", args...)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return nil, errors.New(stderr.String())
	}
	outputStr := out.String()
	outputStr = strings.TrimSpace(outputStr)
	files := strings.Split(outputStr, "\n")

	return files, nil
}
func (s *service) DiffNotCommited() ([]string, error) {
	return s.handleListing("")
}
func (s *service) DiffPrevCommit() ([]string, error) {
	return s.handleListing("HEAD~1")
}
func (s *service) DiffCompare(commitHash string) ([]string, error) {
	return s.handleListing(commitHash)
}
