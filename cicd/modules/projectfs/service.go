// responsible for accessing project file structure and project specific directories
package projectfs

type Service interface {
	AllProjects() ([]string, error)
	AllModules() ([]string, error)

	FilesModules(files []string) []string
	FilesProjects(files []string) []string

	Save(file, content string) error
}
