package internal

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"readme/modules/docs"
	"readme/modules/docs/internal/deps"
	"readme/modules/docs/internal/types"
	"regexp"
	"strings"

	"golang.org/x/tools/go/packages"
)

type ModulePipeline struct{}

// Pipeline.
//
// Legend:
// + stages listed using '+' are required
// * stages listed using '*' are required but have automatic fallback for a section
// - stages listed using '-' are optional
//
// Readme pipeline:
//   - header: `package name`
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

// Breakdown:
// (\d+)\s+([\d.]+)\s+ns/op
// \s+: matches any number of spaces separating the values
// Group 1: (\d+) matches the number of operations
// \s+: matches any number of spaces separating the values
// Group 2: ([\d.]+) matches the floating-point timing performance
// ns/op: ensures we only target actual benchmark measurement lines
var benchMetricRegex = regexp.MustCompile(`\s+(\d+)\s+([\d.]+)\s+ns/op`)

// Breakdown:
// ^ok\s+        -> Starts with 'ok' followed by one or more spaces
// \S+           -> Match package path characters (alphanumeric, slashes, dashes, etc.)
// \s+           -> Spaces before the time duration
// ([\d.]+s)     -> Capture group targeting the float number directly attached to 's' (e.g., 5.919s)
// $             -> Assures it evaluates the exact format at the line's end
var benchDurationRegex = regexp.MustCompile(`(?m)^ok\s+\S+\s+([\d.]+)s$`)

func (s *service) prepareDocToCompare(dic string) string {
	prepared := strings.Trim(dic, " \n")
	prepared = benchMetricRegex.ReplaceAllString(prepared, "X ns/op")
	prepared = benchDurationRegex.ReplaceAllStringFunc(prepared, func(match string) string {
		lastSpace := strings.LastIndex(match, " ")
		if lastSpace != -1 {
			return match[:lastSpace+1] + "Xs"
		}
		return match
	})
	return prepared
}

func (s *service) GenerateModuleDocs(modulePath string) error {
	log.Printf("Generating \"%v\" docs...\n", modulePath)
	readmePath := fmt.Sprintf("%v/readme/README.md", modulePath)
	dir := filepath.Dir(readmePath)

	doc, err := s.GetModuleDocs(modulePath)
	if err != nil {
		return err
	}

	// #nosec G301
	if err := os.MkdirAll(dir, 0777); err != nil {
		return err
	}
	// #nosec G306
	if err := os.WriteFile(readmePath, []byte(doc), 0666); err != nil {
		return err
	}
	// #nosec G302
	if err := os.Chmod(readmePath, 0666); err != nil {
		return err
	}
	return nil
}

func (s *service) DiffModuleDocs(modulePath string) error {
	log.Printf("Comparing \"%v\" docs...\n", modulePath)
	readmePath := fmt.Sprintf("%v/readme/README.md", modulePath)

	file, err := os.ReadFile(readmePath)
	if err != nil {
		return err
	}
	doc, err := s.GetModuleDocs(modulePath)
	if err != nil {
		return err
	}

	filePrepared := s.prepareDocToCompare(string(file))
	docPrepared := s.prepareDocToCompare(doc)
	if filePrepared == docPrepared {
		return nil
	}
	return fmt.Errorf("module \"%v\" is outdated", readmePath)
}

//
//
//

func (s *ModulePipeline) Title(modulePath string) (string, error) {
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

func (s *ModulePipeline) Architecture(modulePath string) (string, error) {
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

func (s *ModulePipeline) Types(modulePath string) (string, error) {
	meta, err := types.NewAST(modulePath)
	if err != nil {
		return "", err
	}
	return meta.String(), nil
}

func (s *ModulePipeline) Bench(modulePath string) (string, error) {
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

func (s *ModulePipeline) Challenges(modulePath string) (string, error) {
	data, err := os.ReadFile(fmt.Sprintf("%v/readme/CHALLENGES.md", modulePath))
	if err != nil {
		return "", nil
	}
	doc := fmt.Sprintf("## Challenges\n%v", string(data))
	return doc, nil
}

func (s *ModulePipeline) Todo(modulePath string) (string, error) {
	data, err := os.ReadFile(fmt.Sprintf("%v/readme/TODO.md", modulePath))
	if err != nil {
		return "", nil
	}
	doc := fmt.Sprintf("## TODO\n%v", string(data))
	return doc, nil
}

func (s *ModulePipeline) Dependencies(modulePath string) (string, error) {
	deps, err := deps.NewAST(modulePath)
	if err != nil {
		return "", err
	}
	return deps.String(), nil
}
