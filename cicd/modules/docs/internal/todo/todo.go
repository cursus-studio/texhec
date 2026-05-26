package todo

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Meta struct {
	Name string
	TODO string
}
type Project struct {
	Meta
	Modules []Meta
}

type TODO struct {
	Projects []Project
}

// accept only:
// ./$projectName/readme/TODO.md
// ./$projectName/modules/$moduleName/readme/TODO.md
func ReadProjects(files []string) TODO {
	sort.Strings(files)
	var todo TODO

	for _, file := range files {
		file = strings.TrimPrefix(file, "./")
		file = filepath.Clean(file)
		cleanPath := filepath.ToSlash(filepath.Clean(file))
		parts := strings.Split(cleanPath, "/")

		if (len(parts) != 3 && len(parts) != 5) || !strings.HasSuffix(file, "/readme/TODO.md") {
			continue
		}

		projectName, moduleName := parts[0], ""
		isModule := strings.Contains(file, fmt.Sprintf("%v/modules", projectName))
		if isModule {
			moduleName = parts[2]
		}

		var currentProject *Project
		if len(todo.Projects) == 0 || todo.Projects[len(todo.Projects)-1].Name != projectName {
			todo.Projects = append(todo.Projects, Project{Meta: Meta{Name: projectName}})
		}
		currentProject = &todo.Projects[len(todo.Projects)-1]

		contentBytes, err := os.ReadFile(file)
		if err != nil {
			continue
		}
		todoContent := strings.TrimSpace(string(contentBytes))

		if !isModule {
			currentProject.TODO = todoContent
			continue
		}

		moduleMeta := Meta{Name: moduleName, TODO: todoContent}
		currentProject.Modules = append(currentProject.Modules, moduleMeta)
	}

	return todo
}

func (t *TODO) String() string {
	if len(t.Projects) == 0 {
		return ""
	}

	res := &strings.Builder{}
	res.WriteString("# TODO\n")
	res.WriteString("This list contains a list of tasks to keep in mind. Often architectural changes to revise and either implement it or omit it entirely\n")

	for _, proj := range t.Projects {
		fmt.Fprintf(res, "## [%s](/%s/readme/README.md)\n", proj.Name, proj.Name)
		if proj.TODO != "" {
			fmt.Fprintf(res, "%s\n", proj.TODO)
		}

		for _, mod := range proj.Modules {
			fmt.Fprintf(res, "- ### [%s](/%s/modules/%s/readme/README.md)\n", mod.Name, proj.Name, mod.Name)
			if mod.TODO != "" {
				fmt.Fprintf(res, "%s\n", mod.TODO)
			}
		}
	}

	return res.String()
}
