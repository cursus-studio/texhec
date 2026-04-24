package test

import (
	"engine/modules/codec"
	codecpkg "engine/modules/codec/pkg"
	enginepkg "engine/pkg"

	"github.com/ogiusek/ioc/v2"
)

type Type struct {
	Value int
}

type setup struct {
	codec codec.Service
}

func NewSetup() setup {
	c := ioc.NewContainer(
		enginepkg.Pkg,
		codecpkg.PkgT[Type],
	)
	return setup{ioc.Get[codec.Service](c)}
}
