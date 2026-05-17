package internal

import (
	"cicd/modules/docs"
	"cicd/world"
	"errors"
	"fmt"
	"os"

	"github.com/ogiusek/ioc/v2"
)

type service struct {
	world.CICDWorld `inject:""`
}

func NewService(c ioc.Dic) docs.Service {
	return ioc.GetServices[*service](c)
}

func (s *service) modules(modulesPath string) ([]string, error) {
	entries, err := os.ReadDir(modulesPath)
	if err != nil {
		return nil, err
	}

	modulePaths := make([]string, 0, len(modulesPath))
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		modulePaths = append(modulePaths, fmt.Sprintf("%v/%v", modulesPath, entry.Name()))
	}
	return modulePaths, nil
}

func (s *service) GenerateModulesDocs(modulesPath string) []error {
	modulePaths, err := s.modules(modulesPath)
	if err != nil {
		return []error{err}
	}
	errs := []error{}
	for _, modulePath := range modulePaths {
		if err := s.GenerateModuleDocs(modulePath); err != nil {
			errs = append(errs, errors.Join(fmt.Errorf("module \"%v\"", modulePath), err))
		}
	}
	return errs
}

func (s *service) DiffModulesDocs(modulesPath string) []error {
	modulePaths, err := s.modules(modulesPath)
	if err != nil {
		return []error{err}
	}
	errs := []error{}
	for _, modulePath := range modulePaths {
		if err := s.DiffModuleDocs(modulePath); err != nil {
			errs = append(errs, errors.Join(fmt.Errorf("module \"%v\"", modulePath), err))
		}
	}
	return errs
}
