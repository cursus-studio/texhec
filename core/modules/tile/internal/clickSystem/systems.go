package clicksystem

import (
	"core/game"
	"core/modules/deploy"
	"core/modules/pathfind"
	"core/modules/tile"
	"core/modules/ui"
	"engine/modules/collider"
	"engine/modules/inputs"
	"engine/modules/render"
	"engine/modules/text"
	"engine/modules/transform"
	"engine/services/ecs"
	"errors"
	"fmt"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type system struct {
	game.GameWorld `inject:""`
}

func NewSystem(c ioc.Dic) ecs.SystemRegister {
	return ecs.NewSystemRegister(func() error {
		s := ioc.GetServices[*system](c)

		events.Listen(s.EventsBuilder(), s.OnClickEntity)
		events.Listen(s.EventsBuilder(), s.SelectEntity)
		return nil
	})
}

func (s *system) OnClickEntity(e tile.ClickEntityEvent) {
	events.Emit(s.Events(), ui.NewSelect[ui.ObjectComponent](e.Entity))

	link, ok := s.Metadata().Link().Get(e.Entity)
	if !ok {
		s.Logger().Log(errors.New("expected entity to have link component"))
		return
	}
	name, ok := s.Metadata().Name().Get(link.Entity)
	if !ok {
		s.Logger().Log(errors.New("expected link to have name component"))
		return
	}
	owner, ok := s.Player().Owner().Get(e.Entity)
	if !ok {
		s.Logger().Log(errors.New("object without owner cannot build"))
		return
	}
	playerName, ok := s.Metadata().Name().Get(owner.Owner)
	if !ok {
		s.Logger().Log(errors.New("expected player to have player component"))
		return
	}

	type Button struct {
		text  string
		event any
	}
	btns := []Button{
		{fmt.Sprintf("%v's %v", playerName.Name, name.Name), nil},
	}

	if deployed, _ := s.Deploy().Component().Get(link.Entity); len(deployed.Deployable) != 0 {
		btns = append(btns, Button{"Deploy", deploy.NewFeatureDeployEvent()})
	}
	if _, ok := s.Pathfind().Speed().Get(e.Entity); ok {
		btns = append(btns, Button{"Move", pathfind.NewFeatureFindPathEvent()})
	}

	for _, p := range s.Ui().ShowMenu() {
		// i want here to display all actions which can be performed by entity
		// currently implement only building
		for _, btn := range btns {
			var btnEntity ecs.EntityID
			if btn.event != nil {
				btnEntity = s.Prototype().Clone(s.Definitions().Hud().Btn)
				s.Inputs().LeftClick().Set(btnEntity, inputs.NewLeftClick(btn.event))
			} else {
				btnEntity = s.Prototype().Clone(s.Definitions().Hud().Text)
			}
			s.Hierarchy().SetParent(btnEntity, p)
			s.Text().Content().Set(btnEntity, text.NewText(btn.text))
		}
	}
}

func (s *system) SelectEntity(e ui.SelectEvent[ui.ObjectComponent]) {
	for _, entity := range e.Entities {
		marker := s.World().NewEntity()
		s.Hierarchy().SetParent(marker, entity)

		s.Render().Mesh().Set(marker, render.NewMesh(s.Definitions().Assets().SquareMesh))
		s.Render().Texture().Set(marker, render.NewTexture(s.Definitions().Hud().Selected))
		s.Groups().InheritGroups(marker)

		s.Collider().Component().Set(marker, collider.NewCollider(s.Definitions().Assets().SquareCollider))

		s.Transform().Pos().Set(marker, transform.NewPos(0, 0, -.1))
		s.Transform().Parent().Set(marker, transform.NewParent(transform.RelativePos|transform.RelativeSizeXYZ))
		s.Ui().Objects().Set(marker, ui.ObjectComponent{})
	}
}
