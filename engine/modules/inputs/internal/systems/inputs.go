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

func NewInputsSystem(c ioc.Dic) ecs.SystemRegister {
	return ecs.NewSystemRegister(func() error {
		s := ioc.GetServices[*inputsSystem](c)
		events.Listen(s.EventsBuilder(), s.OnDefaultFocus)
		events.Listen(s.EventsBuilder(), s.OnFocus)
		events.Listen(s.EventsBuilder(), s.OnUnfocus)
		events.Listen(s.EventsBuilder(), s.OnKeyboardEvent)
		events.Listen(s.EventsBuilder(), s.Listen)
		return nil
	})
}

func (s *inputsSystem) OnDefaultFocus(e inputs.DefaultFocusEvent) {
	for _, entity := range s.Inputs().DefaultFocused().GetEntities() {
		s.Inputs().DefaultFocused().Remove(entity)
	}
	s.Inputs().DefaultFocused().Set(e.Entity, inputs.NewDefaultFocused())
}
func (s *inputsSystem) OnFocus(e inputs.FocusEvent) {
	for _, focusedEntity := range s.Inputs().Focused().GetEntities() {
		s.Inputs().Focused().Remove(focusedEntity)
	}
	s.Inputs().Focused().Set(e.Entity, inputs.NewFocused())
}
func (s *inputsSystem) OnUnfocus(inputs.UnfocusEvent) {
	for _, focusedEntity := range s.Inputs().Focused().GetEntities() {
		s.Inputs().Focused().Remove(focusedEntity)
	}
}

func (s *inputsSystem) OnKeyboardEvent(event sdl.KeyboardEvent) {
	e := inputs.NewKeyboardEvent(event)
	focusedEntities := s.Inputs().Focused().GetEntities()
	if len(focusedEntities) > 1 {
		s.Logger().Log(fmt.Errorf("expected most one focused element"))
		for _, focusedEntity := range focusedEntities {
			s.Inputs().Focused().Remove(focusedEntity)
		}
	}

	if len(focusedEntities) == 0 {
		focusedEntities = s.Inputs().DefaultFocused().GetEntities()
	}

	if len(focusedEntities) == 0 {
		events.Emit(s.Events(), e)
		return
	}

	focusedEntity := focusedEntities[0]
	parents := append([]ecs.EntityID{focusedEntity}, s.Hierarchy().GetOrderedParents(focusedEntity)...)

	for _, entity := range parents {
		comp, ok := s.Inputs().CaptureKeyboard().Get(entity)
		if !ok {
			continue
		}
		events.EmitAny(s.Events(), comp.Capture(e))
		if !comp.Fallthrough() {
			return
		}
	}

	events.Emit(s.Events(), e)
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
