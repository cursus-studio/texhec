package clicksystem

import (
	"core/game"
	"core/modules/deploy"
	"core/modules/pathfind"
	"core/modules/tile"
	"engine/modules/inputs"
	"engine/modules/text"
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

		return nil
	})
}

func (s *system) OnClickEntity(e tile.ClickEntityEvent) {
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
	deployed, _ := s.Deploy().Component().Get(link.Entity)
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

	// deploy
	if len(deployed.Deployable) == 0 {
		goto skipDeploy
	}
	btns = append(btns, Button{"Can deploy", nil})
	for _, deployed := range deployed.Deployable {
		name, ok := s.Metadata().Name().Get(deployed)
		if !ok {
			s.Logger().Log(errors.New("expected entity to have name component"))
			continue
		}
		btn := Button{fmt.Sprintf("%v", name.Name), deploy.NewSelectEvent(e.Entity, deployed)}
		btns = append(btns, btn)
	}
skipDeploy:
	if _, ok := s.Tile().Speed().Get(e.Entity); ok {
		btns = append(btns, Button{"Move", pathfind.NewSelectEvent(e.Entity)})
	}

	for _, p := range s.Ui().Show() {
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
