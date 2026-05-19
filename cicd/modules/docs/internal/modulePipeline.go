package internal

import (
	"bytes"
	"cicd/modules/docs"
	"cicd/modules/docs/internal/deps"
	"cicd/modules/docs/internal/types"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/tools/go/packages"
)

// Readme pipeline:
//   - header
//   - `Architecture` section: comments above `package name`
//     Architecture should explain module what, why and how
//   - `Types` section: performs AST on module `*.go` files.
//     Sub-sections:
//     `type-$name`
//     `method-$typename-$name`
//     `var-$name`
//     `func-$name`
//   - `readme/BENCH.md` or fallback automatic benchmark as `Benchmarks` section
//   - `Challenges` section: `readme/CHALLENGES.md`
//     Challenges purpose is to show case module complexity and to give topics for discussion.
//     It shouldn't contain necessary information to understand how module works.
//   - `Lines of code` section contains cloc result excluding generated readmes
//   - `TODO` section: `readme/TODO.md`
//   - `Dependencies` section: performs AST on module `*/***.go` files and uses
//     import blocks and external method calls to deduce dependencies
func (s *service) GetModuleDocs(modulePath string) (string, error) {
	sections := []string{}
	if section, err := s.Title(modulePath); err == nil && section != "" {
		sections = append(sections, string(section))
	} else if err != nil {
		return "", err
	} else {
		return "", docs.ErrMissingPackage
	}
	if section, err := s.Architecture(modulePath); err == nil && section != "" {
		sections = append(sections, string(section))
	} else if err != nil {
		return "", err
	} else {
		return "", docs.ErrMissingPackageComments
	}
	if section, err := s.Types(modulePath); err == nil && section != "" {
		sections = append(sections, string(section))
	}
	if section, err := s.Bench(modulePath); err == nil && section != "" {
		sections = append(sections, string(section))
	}
	if section, err := s.Challenges(modulePath); err == nil && section != "" {
		sections = append(sections, string(section))
	}
	if section, err := s.LinesOfCode(modulePath); err == nil && section != "" {
		sections = append(sections, string(section))
	}
	if section, err := s.Todo(modulePath); err == nil && section != "" {
		sections = append(sections, string(section))
	}
	if section, err := s.Dependencies(modulePath); err == nil && section != "" {
		sections = append(sections, string(section))
	}
	doc := strings.Join(sections, "\n")
	doc = strings.Trim(doc, " \n")
	return doc, nil
}

// Readme pipeline:
//   - header
//   - `Architecture` section: comments above `package name`
//     Architecture should explain module what, why and how
//   - `Modules` section: reads all project modules
//   - `Challenges` section: `readme/CHALLENGES.md`
//     Challenges purpose is to show case module complexity and to give topics for discussion.
//     It shouldn't contain necessary information to understand how module works.
//   - `Lines of code` section contains cloc result excluding generated readmes
//   - `TODO` section: `readme/TODO.md`
//   - `Dependencies` section: performs AST on module `*/***.go` files and uses
//     import blocks and external method calls to deduce dependencies
func (s *service) GetProjectDocs(projectPath string) (string, error) {
	sections := []string{}
	if section, err := s.Title(projectPath); err == nil && section != "" {
		sections = append(sections, string(section))
	} else if err != nil {
		return "", err
	} else {
		return "", docs.ErrMissingPackage
	}
	if section, err := s.Architecture(projectPath); err == nil && section != "" {
		sections = append(sections, string(section))
	} else if err != nil {
		return "", err
	} else {
		return "", docs.ErrMissingPackageComments
	}
	if section, err := s.Modules(projectPath); err == nil && section != "" {
		sections = append(sections, string(section))
	}
	if section, err := s.Challenges(projectPath); err == nil && section != "" {
		sections = append(sections, string(section))
	}
	if section, err := s.LinesOfCode(projectPath); err == nil && section != "" {
		sections = append(sections, string(section))
	}
	if section, err := s.Todo(projectPath); err == nil && section != "" {
		sections = append(sections, string(section))
	}
	if section, err := s.ThirdPartyDependencies(projectPath); err == nil && section != "" {
		sections = append(sections, string(section))
	}
	doc := strings.Join(sections, "\n")
	doc = strings.Trim(doc, " \n")
	return doc, nil
}

//
//
//

func (s *service) Title(modulePath string) (string, error) {
	if data, err := os.ReadFile(fmt.Sprintf("%v/readme/TITLE.md", modulePath)); err == nil {
		doc := fmt.Sprintf("# %v", string(data))
		return doc, nil
	}
	cfg := &packages.Config{
		Mode: packages.NeedName,
		Dir:  modulePath,
	}
	pkgs, err := packages.Load(cfg, ".")
	if err != nil || len(pkgs) != 1 {
		return "", err
	}
	doc := fmt.Sprintf("# %v", pkgs[0].Name)
	return doc, nil
}

func (s *service) Architecture(modulePath string) (string, error) {
	if data, err := os.ReadFile(fmt.Sprintf("%v/readme/ARCHITECTURE.md", modulePath)); err == nil {
		doc := fmt.Sprintf("## Architecture\n%v", string(data))
		return doc, nil
	}
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedSyntax | packages.NeedTypesInfo,
		Dir:  modulePath,
	}
	pkgs, err := packages.Load(cfg, ".")
	if err != nil || len(pkgs) != 1 {
		return "", err
	}

	docString := ""
	for _, file := range pkgs[0].Syntax {
		if file.Doc == nil {
			continue
		}
		doc := file.Doc.Text()
		if doc == "" {
			continue
		}
		if docString != "" && doc != "" && docString != doc {
			return "", docs.ErrInconsistentPackageComments
		}
		docString = doc
	}

	if docString == "" {
		return "", nil
	}

	doc := fmt.Sprintf("## Architecture\n%v", docString)
	return doc, nil
}

func (s *service) Types(modulePath string) (string, error) {
	meta, err := types.NewAST(modulePath)
	if err != nil {
		return "", err
	}
	return meta.String(), nil
}

func (s *service) Bench(modulePath string) (string, error) {
	if data, err := os.ReadFile(fmt.Sprintf("%v/readme/BENCH.md", modulePath)); err == nil {
		doc := fmt.Sprintf("## Benchmarks\n%v", string(data))
		return doc, nil
	}
	dir := fmt.Sprintf("%v/test", modulePath)
	info, err := os.Stat(dir)
	if os.IsNotExist(err) || !info.IsDir() {
		return "", nil
	}

	cmd := exec.Command("go", "test", "./...", "-bench=.")
	cmd.Dir = dir

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	if err := cmd.Run(); err != nil {
		return "", err
	}
	output := out.String()
	output = strings.Trim(output, " \n")
	if !strings.Contains(output, "PASS") {
		return "", nil
	}
	doc := fmt.Sprintf("## Benchmarks\n```\n$ go test ./... -bench=.\n%v\n```", output)
	return doc, nil
}

func (s *service) LinesOfCode(modulePath string) (string, error) {
	var files []string
	err := filepath.WalkDir(modulePath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || strings.Contains(path, "readme/README.md") {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil || len(files) == 0 {
		return "", err
	}
	cmd := exec.Command("cloc", "--list-file=-")
	var stdinBuffer bytes.Buffer
	for _, file := range files {
		stdinBuffer.WriteString(file + "\n")
	}
	cmd.Stdin = &stdinBuffer

	var stdoutBuffer bytes.Buffer
	cmd.Stdout = &stdoutBuffer
	cmd.Stderr = io.Discard

	if err := cmd.Run(); err != nil {
		return "", err
	}

	prepared := stdoutBuffer.String()
	// Breakdown:
	// (github\.com/\S+) -> Group 1: Captures the base URL/path (e.g., github.com/AlDanial/cloc)
	// \s+v\s+[\d.]+    -> Matches the version string (e.g., v 2.08)
	// .*$              -> Matches everything else to the end of the line (T=0.01 s...)
	var clocCleanupRegex = regexp.MustCompile(`(?m)^(github\.com/\S+)\s+v\s+[\d.]+.*$`)
	prepared = clocCleanupRegex.ReplaceAllString(prepared, "$1")
	return fmt.Sprintf("## Lines of code\n```\n%v\n```", prepared), nil
}

func (s *service) Modules(projectPath string) (string, error) {
	b := &strings.Builder{}

	modules := s.ProjectFS().ProjectModules(projectPath)
	if len(modules) != 0 {
		b.WriteString("## Modules\n")
	}
	for _, module := range modules {
		name := strings.Split(module, "/")
		fmt.Fprintf(b, "- [%v](/%v)\n", name[len(name)-1], module)
	}

	services := s.ProjectFS().ProjectServices(projectPath)
	if len(services) != 0 {
		b.WriteString("## Services\n")
	}
	for _, service := range services {
		name := strings.Split(service, "/")
		fmt.Fprintf(b, "- [%v](/%v)\n", name[len(name)-1], service)
	}

	doc := b.String()
	return doc, nil
}

func (s *service) Challenges(modulePath string) (string, error) {
	data, err := os.ReadFile(fmt.Sprintf("%v/readme/CHALLENGES.md", modulePath))
	if err != nil {
		return "", nil
	}
	doc := fmt.Sprintf("## Challenges\n%v", string(data))
	return doc, nil
}

func (s *service) Todo(modulePath string) (string, error) {
	data, err := os.ReadFile(fmt.Sprintf("%v/readme/TODO.md", modulePath))
	if err != nil {
		return "", nil
	}
	doc := fmt.Sprintf("## TODO\n%v", string(data))
	return doc, nil
}

func (s *service) Dependencies(modulePath string) (string, error) {
	deps, err := deps.NewAST(modulePath)
	if err != nil {
		return "", err
	}
	return deps.String(), nil
}
func (s *service) ThirdPartyDependencies(modulePath string) (string, error) {
	deps, err := deps.NewAST(modulePath)
	if err != nil {
		return "", err
	}
	return deps.ThirdPartyString(), nil
}
