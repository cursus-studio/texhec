package internal

import (
	"cicd/modules/projectfs"
	"cicd/world"
	"engine/modules/datastructures"
	"os"
	"path/filepath"
	"strings"

	"github.com/ogiusek/ioc/v2"
)

type service struct {
	world.CICDWorld `inject:""`
}

func NewService(c ioc.Dic) projectfs.Service {
	s := ioc.GetServices[*service](c)
	return s
}

func (s *service) AllProjects() ([]string, error) {
	entries, err := os.ReadDir(".")
	if err != nil {
		return nil, err
	}

	var projects []string
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		projectName := entry.Name()
		goModPath := filepath.Join(projectName, "go.mod")
		if _, err := os.Stat(goModPath); err == nil {
			projects = append(projects, projectName)
		}
	}
	return projects, nil
}
func (s *service) AllModules() ([]string, error) {
	projects, err := s.AllProjects()
	if err != nil {
		return nil, err
	}
	modules := []string{}
	for _, project := range projects {
		modulesDir := filepath.Join(project, "modules")

		entries, err := os.ReadDir(modulesDir)
		if err != nil || os.IsNotExist(err) {
			return nil, err
		}

		for _, entry := range entries {
			if !entry.IsDir() {
				continue
			}
			modulePath := filepath.Join(modulesDir, entry.Name())
			modules = append(modules, modulePath)
		}
	}
	return modules, nil
}

func extractPathModule(input string) (string, bool) {
	parts := strings.SplitN(input, "/modules/", 2)
	if len(parts) < 2 || !strings.Contains(parts[1], "/") {
		return "", false
	}
	return parts[0] + "/modules/" + strings.SplitN(parts[1], "/", 2)[0], true
}
func (s *service) FilesModules(files []string) []string {
	paths := datastructures.NewSet[string]()
	for _, file := range files {
		path, ok := extractPathModule(file)
		if !ok {
			continue
		}
		paths.Add(path)
	}
	return paths.Get()
}

func (s *service) FilesProjects(files []string) []string {
	paths := datastructures.NewSet[string]()
	for _, file := range files {
		parts := strings.SplitN(file, string(filepath.Separator), 2)
		if len(parts) == 0 || parts[0] == "" {
			continue
		}
		project := parts[0]
		goModPath := filepath.Join(project, "go.mod")
		if _, err := os.Stat(goModPath); err != nil {
			continue
		}
		paths.Add(project)
	}
	return paths.Get()
}

func (s *service) ProjectModules(project string) []string {
	entries, err := os.ReadDir(filepath.Join(project, "modules"))
	if err != nil {
		return nil
	}

	modules := []string{}
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		modules = append(modules, filepath.Join(project, "modules", entry.Name()))
	}
	return modules
}
func (s *service) ProjectServices(project string) []string {
	entries, err := os.ReadDir(filepath.Join(project, "services"))
	if err != nil {
		return nil
	}

	services := []string{}
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		services = append(services, filepath.Join(project, "services", entry.Name()))
	}
	return services
}

func (s *service) Save(file, content string) error {
	dir := filepath.Dir(file)
	// #nosec G301
	if err := os.MkdirAll(dir, 0777); err != nil {
		return err
	}
	// #nosec G306
	if err := os.WriteFile(file, []byte(content), 0666); err != nil {
		return err
	}
	// #nosec G302
	if err := os.Chmod(file, 0666); err != nil {
		return err
	}
	return nil
}
