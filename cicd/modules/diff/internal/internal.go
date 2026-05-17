package internal

import (
	"bytes"
	"cicd/modules/diff"
	"engine/services/datastructures"
	"fmt"
	"os"
	"os/exec"
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
func filesModules(files []string) (modules []string) {
	paths := datastructures.NewSet[string]()
	for _, file := range files {
		path, ok := extractPathModule(file)
		if !ok {
			continue
		}
		paths.Add(path)
	}
	return paths.Get()
}

func execCommand(command string) (string, error) {
	// #nosec G204
	cmd := exec.Command("bash", "-c", fmt.Sprintf("%v", command))

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("command failed: %w (stderr: %s)", err, stderr.String())
	}

	outputStr := out.String()
	outputStr = strings.TrimSpace(outputStr)
	return outputStr, nil
}

func (s *service) DiffUncommited() ([]string, error) {
	outputStr, err := execCommand("git status --porcelain -uall | awk '{print $2}'")
	if err != nil || outputStr == "" {
		return nil, err
	}
	files := []string{}
	for file := range strings.SplitSeq(outputStr, "\n") {
		if _, err := os.Stat(file); err == nil {
			files = append(files, file)
		}

	}
	paths := filesModules(files)
	return paths, nil
}

func (s *service) DiffCommited() ([]string, error) {
	outputStr, err := execCommand("git diff --name-only | tr '\n' '\n'")
	if err != nil || outputStr == "" {
		return nil, err
	}
	paths := filesModules(strings.Split(outputStr, "\n"))
	return paths, nil
}
