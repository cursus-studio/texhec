package internal

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5/plumbing/format/gitignore"
)

func getFilesRespectingGitignore(modulePath string) ([]string, error) {
	var files []string
	root, err := os.OpenRoot(modulePath)
	if err != nil {
		return nil, err
	}
	defer func() { _ = root.Close() }()

	var patterns []gitignore.Pattern
	// 2. Read .gitignore relative to the root jail safely
	if data, err := root.ReadFile(".gitignore"); err == nil {
		for line := range strings.SplitSeq(string(data), "\n") {
			line = strings.TrimSpace(line)
			if line == "" || strings.HasPrefix(line, "#") {
				continue
			}
			patterns = append(patterns, gitignore.ParsePattern(line, nil))
		}
	}

	matcher := gitignore.NewMatcher(patterns)
	err = fs.WalkDir(root.FS(), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if path == "." {
			return nil
		}
		pathSpecs := strings.Split(path, string(filepath.Separator))
		if matcher.Match(pathSpecs, d.IsDir()) {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		if d.IsDir() || strings.Contains(path, "readme/README.md") {
			return nil
		}
		fullPath := filepath.Join(modulePath, path)
		files = append(files, fullPath)
		return nil
	})

	return files, err
}
