package test

import (
	"engine/modules/codec"
	typeregistrypkg "engine/modules/typeregistry/pkg"
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
		typeregistrypkg.PkgT[Type],
	)
	return setup{ioc.Get[codec.Service](c)}
}
