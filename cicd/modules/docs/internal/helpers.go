package internal

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5/plumbing/format/gitignore"
)

func findGitignore(startPath string) (string, string) {
	absStart, err := filepath.Abs(startPath)
	if err != nil {
		return "", ""
	}

	current := absStart
	for {
		target := filepath.Join(current, ".gitignore")
		if _, err := os.Stat(target); err == nil {
			return current, ".gitignore"
		}

		parent := filepath.Dir(current)
		if parent == current {
			break
		}
		current = parent
	}
	return "", ""
}

func getFilesRespectingGitignore(modulePath string) ([]string, error) {
var files []string

	absModulePath, err := filepath.Abs(modulePath)
	if err != nil {
		return nil, err
	}

	var patterns []gitignore.Pattern
	repoRoot, gitignoreName := findGitignore(absModulePath)

	if repoRoot != "" {
		root, err := os.OpenRoot(repoRoot)
		if err == nil {
			if data, err := root.ReadFile(gitignoreName); err == nil {
				for line := range strings.SplitSeq(string(data), "\n") {
					line = strings.TrimSpace(line)
					if line == "" || strings.HasPrefix(line, "#") {
						continue
					}
					patterns = append(patterns, gitignore.ParsePattern(line, nil))
				}
			}
			_ = root.Close()
		}
	}

	matcher := gitignore.NewMatcher(patterns)

	err = filepath.WalkDir(absModulePath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		var relToRepo string
		if repoRoot != "" {
			relToRepo, err = filepath.Rel(repoRoot, path)
			if err != nil {
				return err
			}
		} else {
			relToRepo, _ = filepath.Rel(absModulePath, path)
		}

		if relToRepo == "." {
			return nil
		}

		pathSpecs := strings.Split(relToRepo, string(filepath.Separator))
		if matcher.Match(pathSpecs, d.IsDir()) {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		if d.IsDir() || strings.Contains(path, "readme/README.md") {
			return nil
		}

		files = append(files, path)
		return nil
	})

	return files, err}
