package assetspkg

import (
	"engine/modules/assets"
	"engine/modules/assets/internal"
	"engine/modules/ecs"
	"engine/modules/entityregistry"

	"github.com/ogiusek/ioc/v2"
)

type config struct {
	path string
}

func newConfig(defaultPath string) Config {
	c := &config{}
	c.SetPath(defaultPath)
	return c
}

func (c *config) SetPath(path string) {
	if len(path) != 0 && path[len(path)-1] != '/' {
		path += "/"
	}
	c.path = path
}
func (c *config) GetPath() string { return c.path }

type Config interface {
	SetPath(string)
	GetPath() string
}

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) Config {
		return newConfig("assets/")
	})
	ioc.Register(b, func(c ioc.Dic) assets.Service {
		return internal.NewService(c)
	})
	ioc.Wrap(b, func(c ioc.Dic, registry entityregistry.Service) {
		config := ioc.Get[Config](c)
		registry.Register("path", func(entity ecs.EntityID, structTagValue string) {
			assetsService := ioc.Get[assets.Service](c)
			path := assets.NewPath(config.GetPath() + structTagValue)
			assetsService.Path().Set(entity, path)
		})
	})
})
