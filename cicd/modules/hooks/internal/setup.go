package internal

import (
	"embed"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

//go:embed hooks
var embeddedHooks embed.FS
var embeddedHooksPath string = "hooks"
var hooksDir string = ".git/hooks"

func copyAndMod(srcFile io.Reader, dst string, perm os.FileMode) error {
	dstFile, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	defer func() { _ = dstFile.Close() }()
	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return err
	}
	return os.Chmod(dst, perm)
}

func (s *service) Setup() error {
	files, err := embeddedHooks.ReadDir(embeddedHooksPath)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(hooksDir, 0755); err != nil {
		return err
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		srcName := file.Name()
		dstName := strings.TrimSuffix(srcName, ".sh")
		dstPath := filepath.Join(hooksDir, dstName)

		srcFile, err := embeddedHooks.Open(fmt.Sprintf("%v/%v", embeddedHooksPath, srcName))
		if err != nil {
			return err
		}
		err = copyAndMod(srcFile, dstPath, 0755)
		_ = srcFile.Close()

		if err != nil {
			return err
		}
	}
	return nil
}
