package clicksystems

import (
	"core/modules/definitions"
	"core/modules/tile"
	"core/modules/ui"
	"engine"
	"engine/modules/groups"
	"engine/modules/text"
	"engine/modules/transform"
	"engine/services/ecs"
	"errors"
	"fmt"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type system struct {
	engine.World `inject:"1"`
	Tile         tile.Service            `inject:"1"`
	Definitions  definitions.Definitions `inject:"1"`
	Ui           ui.Service              `inject:"1"`
}

func NewSystems(c ioc.Dic) ecs.SystemRegister {
	return ecs.NewSystemRegister(func() error {
		s := ioc.GetServices[*system](c)

		events.Listen(s.EventsBuilder, s.OnClickObject)

		return nil
	})
}

func (s *system) OnClickObject(e tile.ClickObjectEvent) {
	unit, ok := s.Tile.Pos().Get(e.Unit)
	if !ok {
		s.Logger.Warn(errors.New("expected object to have coords component"))
		return
	}
	for _, p := range s.Ui.Show() {
		// i want here to display all actions which can be performed by entity
		// currently implement only building
		textEntity := s.NewEntity()
		s.Hierarchy.SetParent(textEntity, p)
		s.Transform.Parent().Set(textEntity, transform.NewParent(transform.RelativePos|transform.RelativeSizeXYZ))
		s.Groups.Inherit().Set(textEntity, groups.InheritGroupsComponent{})

		s.Text.Content().Set(textEntity, text.TextComponent{Text: fmt.Sprintf("OBJECT: %v", unit)})
		s.Text.FontSize().Set(textEntity, text.FontSizeComponent{FontSize: 25})
		s.Text.Align().Set(textEntity, text.TextAlignComponent{Vertical: .5, Horizontal: .5})
	}
}
