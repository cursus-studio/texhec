package internal

import (
	"cicd/modules/git"
	"cicd/modules/pipe"
	"cicd/world"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/ogiusek/ioc/v2"
)

type StageCtx struct {
	ChangedModules,
	ChangedProjects,
	Modules,
	Projects,
	PxoFiles []string
}

func NewStageCtx(
	changedModules,
	changedProjects,
	modules,
	projects,
	pxoFiles []string,
) StageCtx {
	return StageCtx{
		ChangedModules:  changedModules,
		ChangedProjects: changedProjects,
		Modules:         modules,
		Projects:        projects,
		PxoFiles:        pxoFiles,
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
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	}
	pwd += "/"

	// pipeline
	// shaders
	shaders := []string{}
	// golang
	// docs
	s.stages = []Stage{
		NewStage("Loading Data", func(ctx StageCtx) error {
			cmd := exec.Command(
				"find", ".", "-type", "f", "-regextype", "posix-extended",
				"-iregex", `.*\.(vert|frag|geom|comp|tesc|tese|glsl)`)
			out, err := cmd.CombinedOutput()
			if err != nil {
				return err
			}
			s := string(out)
			s = strings.Trim(s, " \n")
			shaders = strings.Split(s, "\n")
			return nil
		}),
		// pipeline
		NewStage("Pipeline Security", func(ctx StageCtx) error {
			cmd := exec.Command("trivy", "config", "--skip-dirs", ".cache", "--exit-code", "1", "--quiet", "--severity", "HIGH,CRITICAL", ".")
			if output, err := cmd.CombinedOutput(); err != nil {
				return fmt.Errorf("%s", string(output))
			}
			return nil
		}),

		// shaders
		NewStage("Shaders Validation", func(ctx StageCtx) error {
			// #nosec G204
			cmd := exec.Command("glslangValidator", shaders...)
			if output, err := cmd.CombinedOutput(); err != nil {
				return fmt.Errorf("%s", string(output))
			}
			return nil
		}),
		NewStage("Shaders Lint", func(ctx StageCtx) error {
			cmd := exec.Command("clang-format", "--dry-run", "--Werror")
			cmd.Args = append(cmd.Args, shaders...)
			if output, err := cmd.CombinedOutput(); err != nil {
				return fmt.Errorf("%s", string(output))
			}
			return nil
		}),

		// golang
		NewStage("Deps", func(ctx StageCtx) error {
			for _, proj := range ctx.Projects {
				fmt.Printf("  \"%v\"\n", proj)
				cmd := exec.Command("sh", "-c", "go mod download && go mod tidy -diff && go mod verify")
				cmd.Dir = proj
				if output, err := cmd.CombinedOutput(); err != nil {
					return fmt.Errorf("%s", output)
				}
			}
			return nil
		}).SetFix(func(ctx StageCtx) error {
			for _, proj := range ctx.Projects {
				fmt.Printf("  \"%v\"\n", proj)
				cmd := exec.Command("go", "mod", "tidy")
				cmd.Dir = proj
				if output, err := cmd.CombinedOutput(); err != nil {
					return fmt.Errorf("%s", string(output))
				}
				if err := s.Git().Stage(
					fmt.Sprintf("%v/go.mod", proj),
					fmt.Sprintf("%v/go.sum", proj),
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
			for _, proj := range ctx.Projects {
				fmt.Printf("  \"%v\"\n", proj)
				cmd := exec.Command("go", "build", "-o", "/dev/null")
				cmd.Dir = proj
				if output, err := cmd.CombinedOutput(); err != nil {
					return fmt.Errorf("%s", string(output))
				}
			}
			return nil
		}),
		NewStage("Lint", func(ctx StageCtx) error {
			for _, proj := range ctx.ChangedProjects {
				fmt.Printf("  \"%v\"\n", proj)
				cmd := exec.Command("golangci-lint", "run")
				cmd.Dir = proj
				if output, err := cmd.CombinedOutput(); err != nil {
					return fmt.Errorf("%s", string(output))
				}
			}
			return nil
		}),
		NewStage("Security", func(ctx StageCtx) error {
			for _, proj := range ctx.Projects {
				fmt.Printf("  \"%v\"\n", proj)
				cmd := exec.Command("gosec", "-quiet", "./...")
				cmd.Dir = proj
				if output, err := cmd.CombinedOutput(); err != nil {
					return fmt.Errorf("%s", string(output))
				}
			}
			return nil
		}),
		NewStage("Test", func(ctx StageCtx) error {
			for _, proj := range ctx.Projects {
				fmt.Printf("  \"%v\"\n", proj)
				cmd := exec.Command("go", "test", "./...", "-benchtime=1x")
				cmd.Dir = proj
				if output, err := cmd.CombinedOutput(); err != nil {
					return fmt.Errorf("%s", string(output))
				}
			}
			return nil
		}),

		// assets exports
		NewStage("Assets Export", func(ctx StageCtx) error {
			for _, file := range ctx.PxoFiles {
				fmt.Printf("  checking \"%v\"\n", file)
				ext := filepath.Ext(file)
				base := file[:len(file)-len(ext)]
				base = strings.Replace(base, "src/", "dist/", 1)
				pxoFile := pwd + file

				if _, err := os.Stat(file); os.IsNotExist(err) {
					continue
				}

				if _, err := os.Stat(base + ".png"); os.IsNotExist(err) {
					return fmt.Errorf("missing PNG asset: %s.png", base)
				}

				if _, err := os.Stat(base + ".gif"); err == nil {
					continue
				}

				//pixelorama --headless --quit -- --framecount $(pwd)/assets/src/backgrounds/1.pxo
				//#nosec G204
				countCmd := exec.Command("pixelorama", "--headless", "--quit", "--", "--framecount", pxoFile)
				out, err := countCmd.CombinedOutput()
				if err != nil {
					//#nosec G204
					boxCmd := exec.Command("box64", "pixelorama", "--headless", "--quit", "--", "--framecount", pxoFile)
					boxOut, boxErr := boxCmd.CombinedOutput()
					if boxErr != nil {
						return fmt.Errorf("failed to get framecount for %s:\n - pixelorama err: %v (out: %q)\n - box64 err: %v (out: %q)",
							pxoFile, err, string(out), boxErr, string(boxOut))
					}
					out = boxOut
				}

				frames, err := parseFrameCount(string(out))
				if err != nil {
					return fmt.Errorf("failed to parse framecount for %s: %w", file, err)
				}

				if frames > 1 {
					return fmt.Errorf("missing GIF asset: %s.gif", base)
				}
			}
			return nil
		}).SetFix(func(ctx StageCtx) error {
			if len(ctx.PxoFiles) == 0 {
				return nil
			}
			handleFile := func(file string) error {
				fmt.Printf("  exporting \"%v\"\n", file)
				ext := filepath.Ext(file)
				base := file[:len(file)-len(ext)]
				base = strings.Replace(base, "src/", "dist/", 1)

				if _, err := os.Stat(file); os.IsNotExist(err) {
					return nil
				}
				if err := os.MkdirAll(filepath.Dir(base), 0750); err != nil {
					return fmt.Errorf("failed to create target directory for %s: %w", file, err)
				}

				pngFile := pwd + base + ".png"
				spiteSheetFile := pwd + base + ".spitesheet.png"
				gifFile := pwd + base + ".gif"
				defer func() { _ = os.Remove(spiteSheetFile) }()
				//pixelorama --headless --quit -- --framecount -e -f 1-1 -o $(pwd)/assets/dist/backgrounds/1.png -s -o $(pwd)/assets/dist/backgrounds/1.spritesheet.png $(pwd)/assets/src/backgrounds/1.pxo
				//#nosec G204
				cmd := exec.Command("pixelorama", "--headless", "--quit", "--", "--framecount",
					"-e", "-f 1-1", "-o", pngFile,
					"-s", "-o", spiteSheetFile,
					pwd+file)

				out, err := cmd.CombinedOutput()
				if err != nil {
					return fmt.Errorf("failed to run command for %s: %s %s", file, string(out), err.Error())
				}
				frames, err := parseFrameCount(string(out))
				if err != nil {
					return fmt.Errorf("failed to parse command output for %s: %w", file, err)
				}
				if frames <= 1 {
					return nil
				}

				if err := SpritesheetToGIF(spiteSheetFile, gifFile, frames); err != nil {
					return err
				}
				return nil
			}
			for _, file := range ctx.PxoFiles {
				if err := handleFile(file); err != nil {
					return err
				}
			}
			time.Sleep(time.Second)
			if err := s.Git().Stage("assets/dist"); err != nil {
				return fmt.Errorf("failed to stage file: %v", err.Error())
			}
			return nil
		}),

		// docs
		NewStage("Docs Genaration", func(ctx StageCtx) error {
			for _, module := range ctx.ChangedModules {
				fmt.Printf("  checking \"%v\" module\n", module)
				if err := s.Docs().DiffModule(module); err != nil {
					return err
				}
			}
			for _, proj := range ctx.ChangedProjects {
				fmt.Printf("  checking \"%v\" project\n", proj)
				if err := s.Docs().DiffProject(proj); err != nil {
					return err
				}
			}
			if err := s.Docs().DiffTODO(); err != nil {
				return err
			}
			_ = s.Git().Stage("/readme/TODO.md")
			return nil
		}).SetFix(func(ctx StageCtx) error {
			for _, module := range ctx.ChangedModules {
				fmt.Printf("  generating \"%v\" module\n", module)
				if err := s.Docs().GenerateModule(module); err != nil {
					return err
				}
				// if we cannot stage readme we do not stage it.
				// if error is returned it means that we just generated additional readme.
				_ = s.Git().Stage(fmt.Sprintf("%v/readme/README.md", module))
			}
			for _, proj := range ctx.ChangedProjects {
				fmt.Printf("  generating \"%v\" project\n", proj)
				if err := s.Docs().GenerateProject(proj); err != nil {
					return err
				}
				// if we cannot stage readme we do not stage it.
				// if error is returned it means that we just generated additional readme.
				_ = s.Git().Stage(fmt.Sprintf("%v/readme/README.md", proj))
			}
			if err := s.Docs().GenerateTODO(); err != nil {
				return err
			}
			_ = s.Git().Stage("readme/TODO.md")
			return nil
		}),
		NewStage("Docs Lint", func(ctx StageCtx) error {
			cmd := exec.Command("lychee", "--root-dir", ".", "**/*.md")
			if output, err := cmd.CombinedOutput(); err != nil {
				return fmt.Errorf("%s", string(output))
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
	if output, err := exec.Command("git", "config", "core.hooksPath", s.hooksDirectory).CombinedOutput(); err != nil {
		return fmt.Errorf("%s", string(output))
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
	pxoFiles, err := s.ProjectFS().PxoFiles()
	if err != nil {
		return err
	}

	ctx := NewStageCtx(
		// sync runs on all modules (doesn't cache certain steps)
		modules, projects,
		modules, projects,
		pxoFiles,
	)

	for _, stage := range s.stages {
		if stage.Fix != nil {
			log.Printf("Fix Stage \"%v\"", stage.Name)
			if err := stage.Fix(ctx); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *service) Fix() error {
	// read all files
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
	pxoFiles := slices.DeleteFunc(slices.Clone(changedFiles), func(file string) bool {
		return !strings.HasSuffix(file, ".pxo")
	})

	ctx := NewStageCtx(
		s.ProjectFS().FilesModules(changedFiles),
		s.ProjectFS().FilesProjects(changedFiles),
		modules,
		projects,
		pxoFiles,
	)

	for _, stage := range s.stages {
		if stage.Fix != nil {
			log.Printf("Fix Stage \"%v\"", stage.Name)
			if err := stage.Fix(ctx); err != nil {
				return err
			}
		}
		log.Printf("Verify Stage \"%v\"", stage.Name)
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
	pxoFiles := slices.DeleteFunc(changedFiles, func(file string) bool {
		return !strings.HasSuffix(file, ".pxo")
	})

	ctx := NewStageCtx(
		s.ProjectFS().FilesModules(changedFiles),
		s.ProjectFS().FilesProjects(changedFiles),
		modules,
		projects,
		pxoFiles,
	)
	for _, stage := range s.stages {
		msg := fmt.Sprintf("Verify Stage \"%v\"", stage.Name)
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
	pxoFiles := slices.DeleteFunc(changedFiles, func(file string) bool {
		return !strings.HasSuffix(file, ".pxo")
	})

	ctx := NewStageCtx(
		s.ProjectFS().FilesModules(changedFiles),
		s.ProjectFS().FilesProjects(changedFiles),
		modules,
		projects,
		pxoFiles,
	)
	for _, stage := range s.stages {
		log.Printf("Verify Stage \"%v\"", stage.Name)
		if err := stage.Verify(ctx); err != nil {
			return err
		}
	}
	return nil
}
