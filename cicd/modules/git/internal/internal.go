package internal

import (
	"cicd/modules/git"
	"cicd/world"
	"context"
	"fmt"
	"os"
	"os/exec"

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
