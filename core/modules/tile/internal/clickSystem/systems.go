package clicksystem

import (
	"core/modules/definitions"
	"core/modules/deploy"
	"core/modules/pathfind"
	"core/modules/player"
	"core/modules/tile"
	"core/modules/ui"
	"engine"
	"engine/modules/inputs"
	"engine/modules/text"
	"engine/services/ecs"
	"errors"
	"fmt"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type system struct {
	engine.World `inject:"1"`
	Definitions  definitions.Definitions `inject:"1"`
	Deploy       deploy.Service          `inject:"1"`
	Player       player.Service          `inject:"1"`
	Tile         tile.Service            `inject:"1"`
	Ui           ui.Service              `inject:"1"`
}

func NewSystem(c ioc.Dic) ecs.SystemRegister {
	return ecs.NewSystemRegister(func() error {
		s := ioc.GetServices[*system](c)

		events.Listen(s.EventsBuilder, s.OnClickEntity)

		return nil
	})
}

func (s *system) OnClickEntity(e tile.ClickEntityEvent) {
	link, ok := s.Deploy.Link().Get(e.Entity)
	if !ok {
		s.Logger.Warn(errors.New("expected entity to have link component"))
		return
	}
	name, ok := s.Metadata.Name().Get(link.Deploy)
	if !ok {
		s.Logger.Warn(errors.New("expected link to have name component"))
		return
	}
	deployed, ok := s.Deploy.Component().Get(link.Deploy)
	if !ok {
		s.Logger.Warn(errors.New("expected link to have deploy component"))
		return
	}
	owner, ok := s.Player.Owner().Get(e.Entity)
	if !ok {
		s.Logger.Warn(errors.New("object without owner cannot build"))
		return
	}
	playerName, ok := s.Metadata.Name().Get(owner.Owner)
	if !ok {
		s.Logger.Warn(errors.New("expected player to have player component"))
		return
	}

	type Button struct {
		text  string
		event any
	}
	btns := []Button{
		{fmt.Sprintf("%v's %v", playerName.Name, name.Name), nil},
		{"Can deploy", nil},
	}
	for _, deployed := range deployed.Deployable {
		name, ok := s.Metadata.Name().Get(deployed)
		if !ok {
			s.Logger.Warn(errors.New("expected entity to have name component"))
			continue
		}
		btn := Button{fmt.Sprintf("%v", name.Name), deploy.NewSelectEvent(e.Entity, deployed)}
		btns = append(btns, btn)
	}
	if _, ok := s.Tile.Speed().Get(e.Entity); ok {
		btns = append(btns, Button{"Move", pathfind.NewSelectEvent(e.Entity)})
	}

	for _, p := range s.Ui.Show() {
		// i want here to display all actions which can be performed by entity
		// currently implement only building
		for _, btn := range btns {
			var btnEntity ecs.EntityID
			if btn.event != nil {
				btnEntity = s.Prototype.Clone(s.Definitions.Hud.Btn)
				s.Inputs.LeftClick().Set(btnEntity, inputs.NewLeftClick(btn.event))
			} else {
				btnEntity = s.Prototype.Clone(s.Definitions.Hud.Text)
			}
			s.Hierarchy.SetParent(btnEntity, p)
			s.Text.Content().Set(btnEntity, text.TextComponent{Text: btn.text})
		}
	}
}
