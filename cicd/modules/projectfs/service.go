// responsible for accessing project file structure and project specific directories
package projectfs

type Service interface {
	AllProjects() ([]string, error)
	AllModules() ([]string, error)

	FilesModules(files []string) []string
	FilesProjects(files []string) []string

	ProjectModules(project string) []string
	ProjectServices(project string) []string

	PxoFiles() ([]string, error)

	Save(file, content string) error
}
