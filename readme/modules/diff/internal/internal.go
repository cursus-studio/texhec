package internal

import (
	"bytes"
	"engine/services/datastructures"
	"fmt"
	"os/exec"
	"readme/modules/diff"
	"strings"
)

type service struct{}

func NewService() diff.Service {
	return &service{}
}

func extractPathModule(input string) (string, bool) {
	parts := strings.SplitN(input, "/modules/", 2)
	if len(parts) < 2 || !strings.Contains(parts[1], "/") {
		return "", false
	}
	return parts[0] + "/modules/" + strings.SplitN(parts[1], "/", 2)[0], true
}

func (s *service) GetModifiedModules() ([]string, error) {
	cmd := exec.Command("git", "diff-tree", "--no-commit-id", "--name-only", "-r", "HEAD")

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("git command failed: %w (stderr: %s)", err, stderr.String())
	}

	outputStr := out.String()
	outputStr = strings.TrimSpace(outputStr)
	if len(outputStr) == 0 {
		return []string{}, nil
	}

	files := strings.Split(outputStr, "\n")

	paths := datastructures.NewSet[string]()
	for _, file := range files {
		path, ok := extractPathModule(file)
		if !ok {
			continue
		}
		paths.Add(path)
	}

	return paths.Get(), nil
}
