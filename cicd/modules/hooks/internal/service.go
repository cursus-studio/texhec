package internal

import (
	"cicd/modules/hooks"
	"cicd/world"
	"fmt"

	"github.com/ogiusek/ioc/v2"
)

type service struct {
	world.CICDWorld `inject:""`
}

func NewHooks(c ioc.Dic) hooks.Service {
	return ioc.GetServices[*service](c)
}

// setup is in setup.go

func (s *service) Handle(hook string) error {
	switch hook {
	case "pre-commit":
		modules, err := s.Git().DiffUncommited()
		if err != nil {
			return err
		}
		for _, module := range modules {
			if err := s.Docs().GenerateModuleDocs(module); err != nil {
				return err
			}
			if err := s.Git().Stage(fmt.Sprintf("%v/readme/README.md", module)); err != nil {
				return err
			}
		}
		return nil
	default:
		return fmt.Errorf("http status code 501")
	}
}
