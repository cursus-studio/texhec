package systems

import (
	"engine"
	"engine/modules/inputs"
	"engine/modules/loop"
	"engine/services/ecs"
	"errors"
	"fmt"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
	"github.com/veandco/go-sdl2/sdl"
)

var (
	ErrNotHandledInput error = errors.New("not handled input")
)

type inputsSystem struct {
	engine.EngineWorld `inject:""`
}

func NewInputsSystem(c ioc.Dic) inputs.System {
	return ecs.NewSystemRegister(func() error {
		s := ioc.GetServices[*inputsSystem](c)
		events.Listen(s.EventsBuilder(), s.Listen)
		return nil
	})
}

func (s *inputsSystem) Listen(args loop.FrameEvent) {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		var e any
		switch event := event.(type) {
		case *sdl.AudioDeviceEvent:
			e = *event
		case *sdl.ClipboardEvent:
			e = *event
		case *sdl.CommonEvent:
			e = *event
		case *sdl.ControllerAxisEvent:
			e = *event
		case *sdl.ControllerButtonEvent:
			e = *event
		case *sdl.ControllerDeviceEvent:
			e = *event
		case *sdl.DisplayEvent:
			e = *event
		case *sdl.DollarGestureEvent:
			e = *event
		case *sdl.DropEvent:
			e = *event
		case *sdl.JoyAxisEvent:
			e = *event
		case *sdl.JoyBallEvent:
			e = *event
		case *sdl.JoyButtonEvent:
			e = *event
		case *sdl.JoyDeviceAddedEvent:
			e = *event
		case *sdl.JoyDeviceRemovedEvent:
			e = *event
		case *sdl.JoyHatEvent:
			e = *event
		case *sdl.KeyboardEvent:
			e = *event
		case *sdl.MouseButtonEvent:
			e = *event
		case *sdl.MouseMotionEvent:
			e = *event
		case *sdl.MouseWheelEvent:
			e = *event
		case *sdl.MultiGestureEvent:
			e = *event
		case *sdl.QuitEvent:
			e = *event
		case *sdl.RenderEvent:
			e = *event
		case *sdl.SensorEvent:
			e = *event
		case *sdl.TextInputEvent:
			e = *event
		case *sdl.TextEditingEvent:
			e = *event
		case *sdl.UserEvent:
			e = *event
		case *sdl.WindowEvent:
			e = *event
		case *sdl.TouchFingerEvent:
			e = *event
		default:
			s.Logger().Log(errors.Join(
				ErrNotHandledInput,
				fmt.Errorf("event not handled: type \"%d\": \"%v\"", event.GetType(), event),
			))
		}
		events.EmitAny(s.Events(), e)
	}
}
