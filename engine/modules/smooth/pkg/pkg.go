package smoothpkg

import (
	"engine/modules/smooth"
	"engine/modules/smooth/internal"
	"engine/modules/transition"
	"engine/services/ecs"
	"reflect"

	"github.com/ogiusek/ioc/v2"
)

type config struct {
	components   map[reflect.Type]struct{}
	services     []func(b ioc.Builder)
	firstSystems []func(c ioc.Dic) smooth.StartSystem
	lastSystems  []func(c ioc.Dic) smooth.StopSystem
}

type Config struct {
	*config
}

func NewConfig() Config {
	return Config{
		config: &config{
			components: make(map[reflect.Type]struct{}),
		},
	}
}

type startSystem[Component any] smooth.StartSystem
type stopSystem[Component any] smooth.StopSystem

func SmoothComponent[Component transition.LerpConstraint[Component]](config Config) {
	componentType := reflect.TypeFor[Component]()
	if _, ok := config.components[componentType]; ok {
		return
	}

	config.components[componentType] = struct{}{}
	config.services = append(config.services, func(b ioc.Builder) {
		ioc.Register(b, func(c ioc.Dic) *internal.Service[Component] {
			return internal.NewService[Component](c)
		})
		ioc.Register(b, func(c ioc.Dic) startSystem[Component] {
			return internal.NewFirstSystem[Component](c)
		})
		ioc.Register(b, func(c ioc.Dic) stopSystem[Component] {
			return internal.NewLastSystem[Component](c)
		})
	})
	config.firstSystems = append(config.firstSystems, func(c ioc.Dic) smooth.StartSystem {
		return ioc.Get[startSystem[Component]](c)
	})
	config.lastSystems = append(config.lastSystems, func(c ioc.Dic) smooth.StopSystem {
		return ioc.Get[stopSystem[Component]](c)
	})
}

var Pkg = ioc.NewPkgT(func(b ioc.Builder, config Config) {
	for _, register := range config.services {
		register(b)
	}
	ioc.Register(b, func(c ioc.Dic) smooth.StartSystem {
		return ecs.NewSystemRegister(func() error {
			for _, system := range config.firstSystems {
				if err := system(c).Register(); err != nil {
					return err
				}
			}
			return nil
		})
	})

	ioc.Register(b, func(c ioc.Dic) smooth.StopSystem {
		return ecs.NewSystemRegister(func() error {
			for _, system := range config.lastSystems {
				if err := system(c).Register(); err != nil {
					return err
				}
			}
			return nil
		})
	})
})
