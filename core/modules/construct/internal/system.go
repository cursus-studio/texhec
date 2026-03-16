package internal

import (
	"core/modules/construct"
	"core/modules/definitions"
	"core/modules/tile"
	"core/modules/ui"
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
	Construct    construct.Service       `inject:"1"`

	layer    float32
	dirtySet ecs.DirtySet
}

func NewSystem(c ioc.Dic, layer float32) error {
	s := ioc.GetServices[*system](c)

	s.layer = layer
	s.dirtySet = ecs.NewDirtySet()

	s.Construct.Coords().AddDirtySet(s.dirtySet)

	s.Transform.Pos().AddDependency(s.Construct.Coords())
	s.Transform.Size().AddDependency(s.Construct.Coords())

	s.Render.Mesh().BeforeGet(s.BeforeGet)
	s.Render.Texture().BeforeGet(s.BeforeGet)
	s.Transform.Pos().BeforeGet(s.BeforeGet)
	s.Transform.Size().BeforeGet(s.BeforeGet)

	s.Collider.Component().BeforeGet(s.BeforeGet)
	s.Inputs.LeftClick().BeforeGet(s.BeforeGet)
	s.Inputs.Stack().BeforeGet(s.BeforeGet)

	events.Listen(s.EventsBuilder, s.OnClick)

	return nil
}

func (s *system) BeforeGet() {
	for _, entity := range s.dirtySet.Get() {
		constructComp, ok := s.Construct.Construct().Get(entity)
		if !ok {
			continue
		}
		coords, ok := s.Construct.Coords().Get(entity)
		if !ok {
			continue
		}

		pos := s.Tile.GetPos(coords.Coords)
		pos.Pos[2] = s.layer
		s.Render.Mesh().Set(entity, render.NewMesh(s.GameAssets.SquareMesh))
		s.Render.Texture().Set(entity, render.NewTexture(constructComp.Construct))

		s.Transform.ParentPivotPoint().Set(entity, transform.NewParentPivotPoint(0, 0, .5))
		s.Transform.Pos().Set(entity, pos)
		s.Transform.Size().Set(entity, s.Tile.GetTileSize())

		s.Collider.Component().Set(entity, collider.NewCollider(s.GameAssets.SquareCollider))
		s.Inputs.LeftClick().Set(entity, inputs.NewLeftClick(construct.NewClickEvent(entity)))
		s.Inputs.Stack().Set(entity, inputs.StackComponent{})
	}
}

func (s *system) OnClick(e construct.ClickEvent) {
	coords, ok := s.Construct.Coords().Get(e.Construct)
	if !ok {
		s.Logger.Warn(errors.New("expected construct to have construct coords component"))
		return
	}
	for _, p := range s.Ui.Show() {
		entity := s.NewEntity()
		s.Hierarchy.SetParent(entity, p)
		s.Transform.Parent().Set(entity, transform.NewParent(transform.RelativePos|transform.RelativeSizeXYZ))
		s.Groups.Inherit().Set(entity, groups.InheritGroupsComponent{})

		s.Text.Content().Set(entity, text.TextComponent{Text: fmt.Sprintf("CONSTRUCT: %v", coords)})
		s.Text.FontSize().Set(entity, text.FontSizeComponent{FontSize: 25})
		s.Text.Align().Set(entity, text.TextAlignComponent{Vertical: .5, Horizontal: .5})
	}
}
