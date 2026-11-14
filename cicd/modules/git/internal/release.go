package internal

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func getLatestTag() (major, minor, patch int) {
	cmd := exec.Command("git", "describe", "--tags", "--abbrev=0")
	out, err := cmd.Output()
	if err != nil {
		return 0, 0, 0
	}
	tag := strings.TrimSpace(string(out))
	cleanVersion := strings.TrimPrefix(tag, "v")
	parts := strings.Split(cleanVersion, ".")
	major, _ = strconv.Atoi(parts[0])
	minor, _ = strconv.Atoi(parts[1])
	patch, _ = strconv.Atoi(parts[2])
	return major, minor, patch
}

func (s *service) Version(arg string) (version string, err error) {
	major, minor, patch := getLatestTag()

	switch arg {
	case "major":
		major++
		minor = 0
		patch = 0
	case "minor":
		minor++
		patch = 0
	case "patch":
		patch++
	default:
		return arg, nil
	}

	return fmt.Sprintf("v%d.%d.%d", major, minor, patch), nil
}

func (s *service) Release(version, title string, files []string) error {
	args := []string{"release", "create", version, "--title", title, "--generate-notes"}
	args = append(args, files...)
	// #nosec G204
	cmd := exec.Command("gh", args...)
	out, err := cmd.Output()
	if err != nil {
		return errors.New(string(out))
	}
	return nil
}
