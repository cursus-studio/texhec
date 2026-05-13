package systems

import (
	"engine"
	"engine/modules/inputs"
	"engine/modules/text"
	"engine/services/ecs"
	"errors"
	"fmt"
	"strings"

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
		s.Inputs().TextInput().OnMod(func(ei ecs.EntityID) {
			if !s.World().EntityExists(ei) {
				return
			}
			input, _ := s.Inputs().TextInput().Get(ei)

			res := make([]string, len(input.Lines))
			copy(res, input.Lines)
			if len(res) == 0 {
				res = []string{"|"}
			}
			r := []rune(res[input.CursorRow])
			res[input.CursorRow] = string(r[:input.CursorCol]) + "|" + string(r[input.CursorCol:])

			s.Text().Content().Set(ei, text.NewText(strings.Join(res, "\n")))
		})

		events.Listen(s.EventsBuilder(), s.PollOnFrame)

		events.Listen(s.EventsBuilder(), s.OnDefaultFocus)
		events.Listen(s.EventsBuilder(), s.OnFocus)
		events.Listen(s.EventsBuilder(), s.OnUnfocus)

		events.Listen(s.EventsBuilder(), s.OnTextInputEvent)
		events.Listen(s.EventsBuilder(), s.OnKeyboardEvent)

		return nil
	})
}

func (s *inputsSystem) OnTextInputEvent(event inputs.TextInputEvent) {
	if event.Type != sdl.KEYDOWN {
		return
	}
	input, _ := s.Inputs().TextInput().Get(event.Entity)
	if len(input.Lines) == 0 {
		input.Lines = append(input.Lines, "")
	}
	input.CursorRow = min(len(input.Lines)-1, max(0, input.CursorRow))
	input.CursorCol = min(len(input.Lines[input.CursorRow]), max(0, input.CursorCol))
	switch event.Keysym.Sym {
	case sdl.K_LEFT:
		if input.CursorCol > 0 {
			input.CursorCol--
		} else if input.CursorRow > 0 {
			input.CursorRow--
			input.CursorCol = len(input.Lines[input.CursorRow])
		}
	case sdl.K_RIGHT:
		if input.CursorCol < len(input.Lines[input.CursorRow]) {
			input.CursorCol++
		} else if input.CursorRow < len(input.Lines)-1 {
			input.CursorRow++
			input.CursorCol = 0
		}
	case sdl.K_UP:
		if input.CursorRow > 0 {
			input.CursorRow--
			input.CursorCol = min(input.CursorCol, len(input.Lines[input.CursorRow]))
		}
	case sdl.K_DOWN:
		if input.CursorRow < len(input.Lines)-1 {
			input.CursorRow++
			input.CursorCol = min(input.CursorCol, len(input.Lines[input.CursorRow]))
		}
	case sdl.K_BACKSPACE:
		if input.CursorCol > 0 {
			line := input.Lines[input.CursorRow]
			input.Lines[input.CursorRow] = line[:input.CursorCol-1] + line[input.CursorCol:]
			input.CursorCol--
		} else if input.CursorRow > 0 {
			prevLineLen := len(input.Lines[input.CursorRow-1])
			input.Lines[input.CursorRow-1] += input.Lines[input.CursorRow]
			input.Lines = append(input.Lines[:input.CursorRow], input.Lines[input.CursorRow+1:]...)
			input.CursorRow--
			input.CursorCol = prevLineLen
		}
	case sdl.K_DELETE:
		if input.CursorCol < len(input.Lines[input.CursorRow]) {
			line := input.Lines[input.CursorRow]
			input.Lines[input.CursorRow] = line[:input.CursorCol] + line[input.CursorCol+1:]
		} else if input.CursorRow < len(input.Lines)-1 {
			input.Lines[input.CursorRow] += input.Lines[input.CursorRow+1]
			input.Lines = append(input.Lines[:input.CursorRow+1], input.Lines[input.CursorRow+2:]...)
		}
	case sdl.K_RETURN, sdl.K_KP_ENTER:
		line := input.Lines[input.CursorRow]
		remainingText := line[input.CursorCol:]
		input.Lines[input.CursorRow] = line[:input.CursorCol]

		input.Lines = append(input.Lines[:input.CursorRow+1], append([]string{remainingText}, input.Lines[input.CursorRow+1:]...)...)
		input.CursorRow++
		input.CursorCol = 0
	default:
		if char := sdl.GetKeyName(event.Keysym.Sym); len(char) == 1 {
			line := input.Lines[input.CursorRow]
			input.Lines[input.CursorRow] = line[:input.CursorCol] + char + line[input.CursorCol:]
			input.CursorCol++
		}
	}

	s.Inputs().TextInput().Set(event.Entity, input)
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
		event := comp.Capture(e)
		if apply, ok := event.(inputs.ApplyEntityEvent); ok {
			event = apply.ApplyEntity(entity)
		}
		events.EmitAny(s.Events(), event)
		if !comp.Fallthrough() {
			return
		}
	}

	events.Emit(s.Events(), e)
}
