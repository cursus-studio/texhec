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
	log.Printf("Comparing \"%v\" docs...\n", modulePath)
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
	return fmt.Errorf("module \"%v\" is outdated", readmePath)
}
