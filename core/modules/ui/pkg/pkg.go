package uipkg

import (
	"core/modules/ui"
	"core/modules/ui/internal/systems"
	"core/modules/ui/internal/uiservice"
	typeregistrypkg "engine/modules/typeregistry/pkg"
	"engine/services/ecs"
	"time"

	"github.com/ogiusek/ioc/v2"
)

type Config interface {
	AnimationDuration(time.Duration)
	BgFrameDuration(time.Duration)
}

type config struct {
	animationDuration time.Duration
	bgTimePerFrame    time.Duration
}

func newConfig() Config {
	return &config{
		animationDuration: time.Millisecond * 300, // animation duration
		bgTimePerFrame:    time.Second / 12,       // bgTimePerFrame
	}
}
func (c *config) AnimationDuration(d time.Duration) { c.animationDuration = d }
func (c *config) BgFrameDuration(d time.Duration)   { c.bgTimePerFrame = d }

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		typeregistrypkg.PkgT[ui.AnimatedBackgroundComponent],
		typeregistrypkg.PkgT[ui.CursorCameraComponent],
		typeregistrypkg.PkgT[ui.UiCameraComponent],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
	ioc.Register(b, func(c ioc.Dic) Config { return newConfig() })

	ioc.Register(b, func(c ioc.Dic) ui.Service {
		config := ioc.Get[Config](c).(*config)
		return uiservice.NewService(c, []ecs.SystemRegister{
			systems.NewSystem(c, config.bgTimePerFrame),
			systems.NewCursorSystem(c),
		}, config.animationDuration)
	})
})
