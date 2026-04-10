package internal

import (
	"core/modules/definitions"
	"core/modules/settings"
	"core/modules/ui"
	gamescenes "core/scenes"
	"engine"
	"engine/modules/audio"
	"engine/modules/groups"
	"engine/modules/inputs"
	"engine/modules/scene"
	"engine/modules/text"
	"engine/modules/transform"
	"engine/services/ecs"
	"engine/services/frames"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

// 1. settings text
// 2. quit button

type system struct {
	Definitions definitions.Definitions `inject:"1"`

	engine.World `inject:"1"`
	Ui           ui.Service `inject:"1"`
}

type temporaryToggleColorComponent struct{}

func NewSystem(c ioc.Dic) settings.System {
	return ecs.NewSystemRegister(func() error {
		s := ioc.GetServices[system](c)

		events.ListenE(s.EventsBuilder, func(event settings.EnterSettingsForParentEvent) error {
			return s.ListenRender(event.Parent)
		})
		events.Listen(s.EventsBuilder, s.ListenOnTick)
		events.Listen(s.EventsBuilder, func(settings.EnterSettingsEvent) {
			for _, p := range s.Ui.Show() {
				event := settings.EnterSettingsForParentEvent{Parent: p}
				events.Emit(s.Events, event)
			}
		})

		return nil
	})
}

func (s *system) ListenOnTick(frames.TickEvent) {
	toggleArray := ecs.GetComponentsArray[temporaryToggleColorComponent](s)
	for _, entity := range toggleArray.GetEntities() {
		color, ok := s.Render.Color().Get(entity)
		if !ok {
			color.Color = mgl32.Vec4{1, 1, 1, 1}
		}

		color.Color[1] = 1 - color.Color[1]
		color.Color[2] = 1 - color.Color[2]

		s.Render.Color().Set(entity, color)
	}

}

func (s *system) ListenRender(parent ecs.EntityID) error {
	// render
	// collider
	// click

	// changes
	labelEntity := s.NewEntity()
	s.Hierarchy.SetParent(labelEntity, parent)
	s.Groups.Inherit().Set(labelEntity, groups.InheritGroupsComponent{})

	s.Transform.Parent().Set(labelEntity, transform.NewParent(transform.RelativePos|transform.RelativeSizeX))
	s.Transform.Size().Set(labelEntity, transform.NewSize(1, 50, 1))

	s.Text.Content().Set(labelEntity, text.TextComponent{Text: "SETTINGS"})
	s.Text.FontSize().Set(labelEntity, text.FontSizeComponent{FontSize: 25})
	s.Text.Align().Set(labelEntity, text.TextAlignComponent{Vertical: .5, Horizontal: .5})

	//

	type Button struct {
		text  string
		event any
	}
	btns := []Button{
		{"SHOT", audio.NewPlayEvent(gamescenes.EffectChannel, s.Definitions.ExampleAudio)},
		{"SHOT2", audio.NewPlayEvent(gamescenes.EffectChannel, s.Definitions.ExampleAudio)},
		{"SHOT3", audio.NewPlayEvent(gamescenes.EffectChannel, s.Definitions.ExampleAudio)},
		{"QUIT", scene.NewChangeSceneEvent(gamescenes.MenuID)},
	}

	for _, btn := range btns {
		btnEntity := s.Prototype.Clone(s.Definitions.Hud.Btn)
		s.Hierarchy.SetParent(btnEntity, parent)

		ecs.GetComponentsArray[temporaryToggleColorComponent](s).Set(btnEntity, temporaryToggleColorComponent{})

		s.Text.Content().Set(btnEntity, text.TextComponent{Text: btn.text})
		s.Inputs.LeftClick().Set(btnEntity, inputs.NewLeftClick(btn.event))
	}

	return nil
}
