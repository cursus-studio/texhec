package internal

import (
	"engine"
	"engine/modules/assets"
	"fmt"
	"strings"

	"github.com/ogiusek/ioc/v2"
)

type extensions struct {
	engine.EngineWorld `inject:""`
	extensions         map[string]func(assets.PathComponent) (assets.Asset, error)
}

func NewExtensions(c ioc.Dic) *extensions {
	e := ioc.GetServices[*extensions](c)
	e.extensions = make(map[string]func(assets.PathComponent) (assets.Asset, error))
	return e
}

func (s *extensions) Register(
	/* shouldn't have dots and be after dots in asset */ extension string,
	dispatcher func(path assets.PathComponent) (assets.Asset, error),
) {
	extension = strings.Trim(extension, ".")
	if _, ok := s.extensions[extension]; ok {
		s.Logger().Log(fmt.Errorf("extension \"%v\" is already taken", extension))
		return
	}
	s.extensions[extension] = dispatcher
}

func (s *extensions) ExtensionDispatcher(extension string) (func(assets.PathComponent) (assets.Asset, error), bool) {
	extension = strings.Trim(extension, ".")
	dispatcher, ok := s.extensions[extension]
	return dispatcher, ok
}
