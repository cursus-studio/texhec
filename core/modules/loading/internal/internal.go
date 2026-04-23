package internal

import (
	"core/modules/loading"
	"core/modules/ui"
	gamescenes "core/scenes"
	"engine/modules/camera"
	"engine/modules/loop"
	"engine/modules/text"
	"engine/modules/transform"
	"engine/services/ecs"
	"fmt"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type CamComp struct{}
type TextComp struct{}

type system struct {
	gamescenes.GameWorld `inject:""`

	CamArr  ecs.ComponentsArray[CamComp]
	TextArr ecs.ComponentsArray[TextComp]
}

func NewSystem(c ioc.Dic) loading.System {
	s := ioc.GetServices[*system](c)
	s.CamArr = ecs.GetComponentsArray[CamComp](s.World())
	s.TextArr = ecs.GetComponentsArray[TextComp](s.World())
	return ecs.NewSystemRegister(func() error {
		events.Listen(s.EventsBuilder(), s.Listen)
		return nil
	})
}

func (s *system) Hide() {
	for _, e := range s.CamArr.GetEntities() {
		s.World().RemoveEntity(e)
	}
	for _, e := range s.TextArr.GetEntities() {
		s.World().RemoveEntity(e)
	}
}

func (s *system) Render(message string) {
	if len(s.TextArr.GetEntities()) == 1 {
		textEntity := s.TextArr.GetEntities()[0]
		s.Text().Content().Set(textEntity, text.NewText(message))
		return
	}

	cameraEntity := s.World().NewEntity()
	s.Camera().Ortho().Set(cameraEntity, camera.NewOrtho(-5, 5))
	s.CamArr.Set(cameraEntity, CamComp{})

	background := s.World().NewEntity()
	s.Hierarchy().SetParent(background, cameraEntity)
	s.Transform().Pos().Set(background, transform.NewPos(0, 0, 1))
	s.Transform().PivotPoint().Set(background, transform.NewPivotPoint(.5, .5, 0))
	s.Transform().ParentPivotPoint().Set(background, transform.NewParentPivotPoint(.5, .5, 0))
	s.Ui().AnimatedBackground().Set(background, ui.AnimatedBackgroundComponent{})

	textEntity := s.World().NewEntity()
	s.TextArr.Set(textEntity, TextComp{})
	s.Hierarchy().SetParent(textEntity, cameraEntity)
	s.Transform().Pos().Set(textEntity, transform.NewPos(0, 0, 2))
	s.Transform().Parent().Set(textEntity, transform.NewParent(transform.RelativePos))

	s.Text().Content().Set(textEntity, text.NewText(message))
	s.Text().FontSize().Set(textEntity, text.NewFontSize(32))
	s.Text().Break().Set(textEntity, text.NewBreak(text.BreakNone))
	s.Text().Align().Set(textEntity, text.NewAlign(.5, .5))
}

func (s *system) Listen(loop.FrameEvent) {
	progress := s.Batcher().Progress()
	if progress == -1 {
		s.Hide()
		return
	}

	message := fmt.Sprintf("Loading... %6.2f%%", progress*100)
	s.Render(message)
}
