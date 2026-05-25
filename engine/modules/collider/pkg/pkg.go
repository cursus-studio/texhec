package colliderpkg

import (
	"engine/modules/collider"
	"engine/modules/collider/internal/collisions"
	typeregistrypkg "engine/modules/typeregistry/pkg"

	"github.com/ogiusek/ioc/v2"
)

type config struct {
	chunkSize float32
}

func NewConfig() Config {
	return &config{
		1000,
	}
}
func (c *config) GetChunkSize() float32          { return c.chunkSize }
func (c *config) SetChunkSize(chunkSize float32) { c.chunkSize = chunkSize }

type Config interface {
	GetChunkSize() float32
	SetChunkSize(float32)
}

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		typeregistrypkg.PkgT[collider.Component],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
	ioc.Register(b, func(c ioc.Dic) Config {
		return NewConfig()
	})
	ioc.Register(b, func(c ioc.Dic) collider.Service {
		return collisions.NewService(c, ioc.Get[Config](c).GetChunkSize())
	})
})
