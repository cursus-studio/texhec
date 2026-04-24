package audiopkg

import (
	"engine/modules/assets"
	"engine/modules/audio"
	"engine/modules/audio/internal"
	codecpkg "engine/modules/codec/pkg"
	"os"

	"github.com/ogiusek/ioc/v2"
	"github.com/veandco/go-sdl2/mix"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		codecpkg.PkgT[audio.StopEvent],
		codecpkg.PkgT[audio.PlayEvent],
		codecpkg.PkgT[audio.QueueEvent],
		codecpkg.PkgT[audio.QueueEndlessEvent],
		codecpkg.PkgT[audio.SetMasterVolumeEvent],
		codecpkg.PkgT[audio.SetChannelVolumeEvent],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
	ioc.Register(b, func(c ioc.Dic) internal.Service {
		return internal.NewService(c)
	})
	ioc.Register(b, func(c ioc.Dic) audio.PlayerService { return ioc.Get[internal.Service](c) })
	ioc.Register(b, func(c ioc.Dic) audio.VolumeService { return ioc.Get[internal.Service](c) })
	ioc.Register(b, func(c ioc.Dic) audio.Service { return ioc.Get[internal.Service](c) })

	ioc.Register(b, func(c ioc.Dic) audio.System {
		return internal.NewSystem(c)
	})

	ioc.Wrap(b, func(c ioc.Dic, b assets.Service) {
		b.Register("wav", func(id assets.PathComponent) (assets.Asset, error) {
			source, err := os.ReadFile(id.Path)
			if err != nil {
				return nil, err
			}
			chunk, err := mix.QuickLoadWAV(source)
			if err != nil {
				return nil, err
			}
			audio := audio.NewAudioAsset(chunk, source)
			return audio, nil
		})
	})
})
