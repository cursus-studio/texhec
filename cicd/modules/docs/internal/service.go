package internal

import (
	"cicd/modules/docs"
	"cicd/world"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/ogiusek/ioc/v2"
)

type service struct {
	world.CICDWorld `inject:""`
}

func NewService(c ioc.Dic) docs.Service {
	return ioc.GetServices[*service](c)
}

// Matches and eliminates platform-specific headers entirely
var platformRegex = regexp.MustCompile(`(?m)^(goos|goarch|cpu|gpu):.*\n?`)

// Matches the CPU core count suffix (e.g., -8, -4) right before the metrics start
var coreSuffixRegex = regexp.MustCompile(`(-\d+)\s+\d+`)

// Matches the metric numbers precisely without swallowing preceding newlines
// Group 1 catches the iteration count, Group 2 catches the speed
var benchMetricRegex = regexp.MustCompile(`\s+(\d+)\s+([\d.]+)\s+ns/op`)

// Safely targets the final execution summary duration line
var benchDurationRegex = regexp.MustCompile(`(?m)^ok\s+\S+\s+([\d.]+)s$`)

func (s *service) prepareDocToCompare(dic string) string {
	prepared := strings.Trim(dic, " \n")
	prepared = platformRegex.ReplaceAllString(prepared, "")
	prepared = coreSuffixRegex.ReplaceAllString(prepared, "-X")
	prepared = benchMetricRegex.ReplaceAllString(prepared, " X ns/op")
	prepared = benchDurationRegex.ReplaceAllStringFunc(prepared, func(match string) string {
		lastSpace := strings.LastIndex(match, " ")
		if lastSpace != -1 {
			return match[:lastSpace+1] + "Xs"
		}
		return match
	})
	return strings.TrimSpace(prepared)
}

func (s *service) GenerateModuleDocs(modulePath string) error {
	if _, err := os.Stat(modulePath); os.IsNotExist(err) {
		return nil
	}
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
	if _, err := os.Stat(modulePath); os.IsNotExist(err) {
		return nil
	}
	log.Printf("Comparing \"%v\" module docs...\n", modulePath)
	readmePath := filepath.Join(filepath.Clean(modulePath), "readme", "README.md")
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
	return fmt.Errorf("module \"%v\" is outdated ```\n%v\n``` != ```\n%v\n```", readmePath, docPrepared, filePrepared)
}

func (s *service) GenerateProjectDocs(projectPath string) error {
	if _, err := os.Stat(projectPath); os.IsNotExist(err) {
		return nil
	}
	log.Printf("Generating \"%v\" docs...\n", projectPath)
	readmePath := fmt.Sprintf("%v/readme/README.md", projectPath)
	dir := filepath.Dir(readmePath)

	doc, err := s.GetProjectDocs(projectPath)
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
func (s *service) DiffProjectDocs(projectPath string) error {
	if _, err := os.Stat(projectPath); os.IsNotExist(err) {
		return nil
	}
	log.Printf("Comparing \"%v\" project docs...\n", projectPath)
	readmePath := filepath.Join(filepath.Clean(projectPath), "readme", "README.md")
	file, err := os.ReadFile(readmePath)
	if err != nil {
		return err
	}
	doc, err := s.GetProjectDocs(projectPath)
	if err != nil {
		return err
	}

	filePrepared := s.prepareDocToCompare(string(file))
	docPrepared := s.prepareDocToCompare(doc)
	if filePrepared == docPrepared {
		return nil
	}
	return fmt.Errorf("project \"%v\" is outdated ```\n%v\n``` != ```\n%v\n```", readmePath, docPrepared, filePrepared)
}
