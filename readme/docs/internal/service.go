package internal

import (
	"errors"
	"fmt"
	"os"
	"readme/docs"
)

type Service interface {
	// Generates module documentation in `$modulePath/readme/README.md`
	GenerateModuleDocs(modulePath string) error

	// Generates modules documentation in `$modulesPath/$moduleName/readme/README.md`
	GenerateModulesDocs(modulesPath string) []error
}

//

type service struct {
	ModulePipeline
}

func NewService() docs.Service {
	return &service{}
}

func (s *service) GenerateModulesDocs(modulesPath string) []error {
	errs := []error{}
	entries, err := os.ReadDir(modulesPath)
	if err != nil {
		errs = append(errs, err)
		return errs
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		modulePath := fmt.Sprintf("%v/%v", modulesPath, entry.Name())
		if err := s.GenerateModuleDocs(modulePath); err != nil {
			errs = append(errs, errors.Join(fmt.Errorf("missing path in \"%v\"", modulePath), err))
		}
	}
	return errs
}
