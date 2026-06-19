package inputs

import (
	"engine/modules/ecs"
	"engine/services/datastructures"
	"reflect"
	"strings"

	"github.com/veandco/go-sdl2/sdl"
)

type KeyboardEvent struct {
	sdl.KeyboardEvent
}

func NewKeyboardEvent(e sdl.KeyboardEvent) KeyboardEvent {
	return KeyboardEvent{e}
}

//

type TextInputComponent struct {
	Lines []string // split in lines
	CursorRow,
	CursorCol int
}

func (c *TextInputComponent) Text() string {
	return strings.Join(c.Lines, "\n")
}

//

// handles inputs and saves change in component
type TextInputEvent struct {
	Entity ecs.EntityID
	KeyboardEvent
}

func NewTextInputEvent() TextInputEvent {
	return TextInputEvent{}
}

var textInputEventCaptures = datastructures.NewSet[reflect.Type]()

func (TextInputEvent) CapturesEvents() datastructures.SetReader[reflect.Type] {
	return textInputEventCaptures
}
func (e TextInputEvent) Capture(event any) any {
	e.KeyboardEvent = event.(KeyboardEvent)
	return e
}
func (e TextInputEvent) ApplyEntity(entity ecs.EntityID) any {
	e.Entity = entity
	return e
}
func (e TextInputEvent) Fallthrough() bool { return false }

func init() {
	textInputEventCaptures.Add(reflect.TypeFor[KeyboardEvent]())
}
