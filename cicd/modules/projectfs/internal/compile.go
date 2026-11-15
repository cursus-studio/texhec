package internal

import (
	"fmt"
	"os"
	"os/exec"
)

func (s *service) Compile() (files []string, cleanUp func(), err error) {
	addedFiles := []string{}
	cleanUp = func() {
		for _, file := range addedFiles {
			_ = os.Remove(file)
		}
	}
	{ // windows
		windowsCmd := exec.Command("go", "-C", "core", "build", "-tags", "release", `-ldflags=-s -w -extldflags '-static'`, "-o", "texhec.exe", ".")
		windowsCmd.Env = append(os.Environ(),
			"CGO_ENABLED=1",
			"GOOS=windows",
			"GOARCH=amd64",
			"CC=x86_64-w64-mingw32-gcc",
			`CGO_LDFLAGS=-static -lmingw32 -lSDL2main -lSDL2 -lSDL2_image -lSDL2_mixer -lm -ldinput8 -ldxguid -lsetupapi -lgdi32 -lwinmm -limm32 -lole32 -loleaut32 -lshell32 -lversion -luuid`,
		)
		if out, err := windowsCmd.CombinedOutput(); err != nil {
			return nil, cleanUp, fmt.Errorf("failed compiling for windows due to: '%s'", out)
		}
		addedFiles = append(addedFiles, "core/texhec.exe")
		zip := "windows.zip"
		if err := zipFiles(zip, "core/assets", "core/texhec.exe"); err != nil {
			return nil, cleanUp, err
		}
		addedFiles = append(addedFiles, "windows.zip")
		files = append(files, zip)
	}
	{ // linux
		linuxCmd := exec.Command("go", "-C", "core", "build", "-tags", "release", "-o", "texhec", ".")
		linuxCmd.Env = append(os.Environ(),
			"CGO_ENABLED=1",
			"GOOS=linux",
			"GOARCH=amd64",
		)
		if out, err := linuxCmd.CombinedOutput(); err != nil {
			return nil, cleanUp, fmt.Errorf("failed compiling for windows due to: '%s'", out)
		}
		// #nosec G302
		if err := os.Chmod("core/texhec", 0755); err != nil {
			return nil, cleanUp, fmt.Errorf("failed setting executable permissions: %w", err)
		}
		zip := "linux.zip"
		addedFiles = append(addedFiles, "core/texhec")
		if err := zipFiles(zip, "core/assets", "core/texhec"); err != nil {
			return nil, cleanUp, err
		}
		files = append(files, zip)
		addedFiles = append(addedFiles, "linux.zip")
	}
	return files, cleanUp, err
}
