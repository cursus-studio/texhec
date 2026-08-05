package internal

import (
	"bytes"
	"cicd/modules/git"
	"cicd/world"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/google/go-github/v60/github"
	"github.com/ogiusek/ioc/v2"
	"golang.org/x/oauth2"
)

type service struct {
	world.CICDWorld `inject:""`
}

func NewService(c ioc.Dic) git.Service {
	return ioc.GetServices[*service](c)
}

func (s *service) FilterRemoved(allFiles []string) []string {
	files := []string{}
	for _, file := range allFiles {
		if _, err := os.Stat(file); err == nil {
			files = append(files, file)
		}
	}
	return files
}

func (s *service) handleListing(comparedCommit string) ([]string, error) {
	args := []string{"--no-pager", "diff", "--name-only", "--staged"}
	if comparedCommit != "" {
		args = append(args, comparedCommit)
	}
	args = append(args, "HEAD")
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

func (s *service) Stage(directories ...string) error {
	if len(directories) == 0 {
		return nil
	}
	args := []string{"add"}
	args = append(args, directories...)
	cmd := exec.Command("git", args...)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("%s", string(output))
	}
	return nil
}

func (s *service) SetStatus(state git.State, desc string) error {
	ctx := context.Background()
	token := os.Getenv("TOKEN")
	owner := os.Getenv("OWNER")
	repo := os.Getenv("REPO")
	sha := os.Getenv("GIT_COMMIT")
	buildURL := os.Getenv("BUILD_URL")
	contextName := os.Getenv("CONTEXT")

	if token == "" || owner == "" || repo == "" || sha == "" {
		return fmt.Errorf("missing required environment variables (TOKEN, OWNER, REPO, or GIT_COMMIT)")
	}

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	status := &github.RepoStatus{
		State:       github.String(string(state)),
		TargetURL:   github.String(buildURL),
		Description: github.String(desc),
		Context:     github.String(contextName),
	}

	_, _, err := client.Repositories.CreateStatus(ctx, owner, repo, sha, status)
	if err != nil {
		return err
	}

	return nil
}
