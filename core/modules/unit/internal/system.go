package internal

import (
	"core/modules/definitions"
	"core/modules/tile"
	"core/modules/ui"
	"core/modules/unit"
	"engine"
	"engine/modules/collider"
	"engine/modules/groups"
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
	engine.World `inject:"1"`
	Tile         tile.Service            `inject:"1"`
	GameAssets   definitions.Definitions `inject:"1"`
	Ui           ui.Service              `inject:"1"`
	Unit         unit.Service            `inject:"1"`

	dirtySet ecs.DirtySet
}

func NewSystem(c ioc.Dic) error {
	s := ioc.GetServices[*system](c)

	s.dirtySet = ecs.NewDirtySet()

	s.Unit.Unit().AddDirtySet(s.dirtySet)

	s.Tile.Layer().BeforeGet(s.BeforeGet)

	s.Render.Mesh().BeforeGet(s.BeforeGet)
	s.Render.Texture().BeforeGet(s.BeforeGet)

	s.Collider.Component().BeforeGet(s.BeforeGet)
	s.Inputs.LeftClick().BeforeGet(s.BeforeGet)
	s.Inputs.Stack().BeforeGet(s.BeforeGet)

	events.Listen(s.EventsBuilder, s.OnClick)

	return nil
}

func (s *system) BeforeGet() {
	for _, entity := range s.dirtySet.Get() {
		unitComp, ok := s.Unit.Unit().Get(entity)
		if !ok {
			continue
		}

		s.Tile.Layer().Set(entity, tile.NewLayer(3))

		s.Render.Mesh().Set(entity, render.NewMesh(s.GameAssets.SquareMesh))
		s.Render.Texture().Set(entity, render.NewTexture(unitComp.Unit))

		s.Collider.Component().Set(entity, collider.NewCollider(s.GameAssets.SquareCollider))
		s.Inputs.LeftClick().Set(entity, inputs.NewLeftClick(unit.NewClickEvent(entity)))
		s.Inputs.Stack().Set(entity, inputs.StackComponent{})
	}
}

func (s *system) OnClick(e unit.ClickEvent) {
	unit, ok := s.Unit.Unit().Get(e.Unit)
	if !ok {
		s.Logger.Warn(errors.New("expected unit to have unit coords component"))
		return
	}
	for _, p := range s.Ui.Show() {
		entity := s.NewEntity()
		s.Hierarchy.SetParent(entity, p)
		s.Transform.Parent().Set(entity, transform.NewParent(transform.RelativePos|transform.RelativeSizeXYZ))
		s.Groups.Inherit().Set(entity, groups.InheritGroupsComponent{})

		s.Text.Content().Set(entity, text.TextComponent{Text: fmt.Sprintf("UNIT: %v", unit)})
		s.Text.FontSize().Set(entity, text.FontSizeComponent{FontSize: 25})
		s.Text.Align().Set(entity, text.TextAlignComponent{Vertical: .5, Horizontal: .5})
	}
}
