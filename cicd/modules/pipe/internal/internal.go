package internal

import (
	"cicd/modules/git"
	"cicd/modules/pipe"
	"cicd/world"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"os/exec"

	"github.com/ogiusek/ioc/v2"
)

type StageCtx struct {
	ChangedModules,
	Modules,
	Projects []string
}

func NewStageCtx(
	changedModules,
	modules,
	projects []string,
) StageCtx {
	return StageCtx{
		ChangedModules: changedModules,
		Modules:        modules,
		Projects:       projects,
	}
}

//

type Stage struct {
	Name string
	// fix can be nil
	Fix,
	Verify func(StageCtx) error
}

func NewStage(
	name string,
	verify func(ctx StageCtx) error,
) Stage {
	return Stage{
		Name:   name,
		Verify: verify,
	}
}
func (s Stage) SetFix(fix func(ctx StageCtx) error) Stage {
	s.Fix = fix
	return s
}

type service struct {
	world.CICDWorld `inject:""`
	hooksDirectory  string

	stages []Stage
}

func NewService(c ioc.Dic) pipe.Service {
	s := ioc.GetServices[*service](c)
	s.hooksDirectory = ".git-hooks"

	s.stages = []Stage{
		NewStage("Deps", func(ctx StageCtx) error {
			for _, dir := range ctx.Projects {
				cmd := exec.Command("sh", "-c", "go mod download && go mod tidy -diff && go mod verify")
				cmd.Stdout = io.Discard
				cmd.Stderr = os.Stderr
				cmd.Dir = dir
				if err := cmd.Run(); err != nil {
					return err
				}
			}
			return nil
		}).SetFix(func(ctx StageCtx) error {
			for _, dir := range ctx.Projects {
				cmd := exec.Command("go", "mod", "tidy")
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Dir = dir
				if err := cmd.Run(); err != nil {
					return err
				}
				if err := s.Git().Stage(
					fmt.Sprintf("%v/go.mod", dir),
					fmt.Sprintf("%v/go.sum", dir),
				); err != nil {
					return err
				}
			}
			if err := s.Git().Stage("go.work", "go.work.sum"); err != nil {
				return err
			}
			return nil
		}),

		NewStage("Build", func(ctx StageCtx) error {
			for _, dir := range ctx.Projects {
				cmd := exec.Command("go", "build", "-o", "/dev/null")
				cmd.Stdout = io.Discard
				cmd.Stderr = os.Stderr
				cmd.Dir = dir
				if err := cmd.Run(); err != nil {
					return err
				}
			}
			return nil
		}),

		NewStage("Security", func(ctx StageCtx) error {
			for _, proj := range ctx.Projects {
				cmd := exec.Command("gosec", "-quiet", "./...")
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Dir = proj
				if err := cmd.Run(); err != nil {
					return err
				}
			}
			return nil
		}),

		NewStage("Pipeline Security", func(ctx StageCtx) error {
			cmd := exec.Command("trivy", "config", "--exit-code", "1", "--quiet", "--severity", "HIGH,CRITICAL", ".")
			cmd.Stdout = io.Discard
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				return err
			}
			return nil
		}),

		NewStage("Lint", func(ctx StageCtx) error {
			for _, proj := range ctx.Projects {
				cmd := exec.Command("golangci-lint", "run")
				cmd.Stdout = io.Discard
				cmd.Stderr = os.Stderr
				cmd.Dir = proj
				if err := cmd.Run(); err != nil {
					return err
				}
			}
			return nil
		}),

		NewStage("Generated Docs", func(ctx StageCtx) error {
			for _, module := range ctx.ChangedModules {
				if err := s.Docs().DiffModuleDocs(module); err != nil {
					return err
				}
			}
			return nil
		}).SetFix(func(ctx StageCtx) error {
			for _, module := range ctx.ChangedModules {
				if err := s.Docs().GenerateModuleDocs(module); err != nil {
					return err
				}
				if err := s.Git().Stage(fmt.Sprintf("%v/readme/README.md", module)); err != nil {
					return err
				}
			}
			return nil
		}),

		NewStage("Markdown Lint", func(ctx StageCtx) error {
			cmd := exec.Command("lychee", "--root-dir", ".", "**/*.md")
			cmd.Stdout = io.Discard
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				return err
			}
			return nil
		}),

		NewStage("Test", func(ctx StageCtx) error {
			for _, proj := range ctx.Projects {
				cmd := exec.Command("go", "test", "./...", "-benchtime=1x")
				cmd.Stdout = io.Discard
				cmd.Stderr = os.Stderr
				cmd.Dir = proj
				if err := cmd.Run(); err != nil {
					return err
				}
			}
			return nil
		}),
	}

	return s
}

func (s *service) Setup() error {
	root, err := os.OpenRoot(s.hooksDirectory)
	if err != nil {
		return err
	}
	defer func() { _ = root.Close() }()

	if err := fs.WalkDir(root.FS(), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		//#nosec G302
		return root.Chmod(path, 0755)
	}); err != nil {
		return err
	}
	//#nosec G204
	if err := exec.Command("git", "config", "core.hooksPath", s.hooksDirectory).Run(); err != nil {
		return err
	}
	return nil
}

func (s *service) Sync() error {
	projects, err := s.ProjectFS().AllProjects()
	if err != nil {
		return err
	}
	modules, err := s.ProjectFS().AllModules()
	if err != nil {
		return err
	}

	ctx := NewStageCtx(
		// sync runs on all modules (doesn't cache certain steps)
		modules, modules,
		projects,
	)

	for _, stage := range s.stages {
		log.Printf("Stage \"%v\"", stage.Name)
		if stage.Fix != nil {
			if err := stage.Fix(ctx); err != nil {
				return err
			}
		}
		if err := stage.Verify(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (s *service) Fix() error {
	projects, err := s.ProjectFS().AllProjects()
	if err != nil {
		return err
	}
	modules, err := s.ProjectFS().AllModules()
	if err != nil {
		return err
	}
	changedFiles, err := s.Git().DiffNotCommited()
	if err != nil {
		return err
	}
	changedModules := s.ProjectFS().FilesModules(changedFiles)

	ctx := NewStageCtx(
		changedModules,
		modules,
		projects,
	)

	for _, stage := range s.stages {
		log.Printf("Stage \"%v\"", stage.Name)
		if stage.Fix != nil {
			if err := stage.Fix(ctx); err != nil {
				return err
			}
		}
		if err := stage.Verify(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (s *service) Cloud(commitHash string) error {
	projects, err := s.ProjectFS().AllProjects()
	if err != nil {
		return err
	}
	modules, err := s.ProjectFS().AllModules()
	if err != nil {
		return err
	}
	changedFiles, err := s.Git().DiffCompare(commitHash)
	if err != nil {
		return err
	}
	changedModules := s.ProjectFS().FilesModules(changedFiles)

	ctx := NewStageCtx(
		changedModules,
		modules,
		projects,
	)
	for _, stage := range s.stages {
		msg := fmt.Sprintf("Stage \"%v\"", stage.Name)
		log.Print(msg)
		if err := s.Git().SetStatus(git.Pending, msg); err != nil {
			return err
		}
		if err := stage.Verify(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (s *service) Verify(commitHash string) error {
	projects, err := s.ProjectFS().AllProjects()
	if err != nil {
		return err
	}
	modules, err := s.ProjectFS().AllModules()
	if err != nil {
		return err
	}
	changedFiles, err := s.Git().DiffCompare(commitHash)
	if err != nil {
		return err
	}
	changedModules := s.ProjectFS().FilesModules(changedFiles)

	ctx := NewStageCtx(
		changedModules,
		modules,
		projects,
	)
	for _, stage := range s.stages {
		log.Printf("Stage \"%v\"", stage.Name)
		if err := stage.Verify(ctx); err != nil {
			return err
		}
	}
	return nil
}
