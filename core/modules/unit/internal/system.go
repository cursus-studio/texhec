package internal

import (
	"core/modules/definitions"
	"core/modules/tile"
	"core/modules/ui"
	"core/modules/unit"
	"engine"
	"engine/modules/collider"
	"engine/modules/grid"
	"engine/modules/groups"
	"engine/modules/inputs"
	"engine/modules/render"
	"engine/modules/text"
	"engine/modules/transform"
	"engine/modules/transition"
	"engine/services/ecs"
	"engine/services/frames"
	"errors"
	"fmt"
	"math"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type system struct {
	engine.World `inject:"1"`
	Tile         tile.Service            `inject:"1"`
	GameAssets   definitions.Definitions `inject:"1"`
	Ui           ui.Service              `inject:"1"`
	Unit         unit.Service            `inject:"1"`

	layer    float32
	dirtySet ecs.DirtySet
}

func NewSystem(c ioc.Dic, layer float32) error {
	s := ioc.GetServices[*system](c)

	s.layer = layer
	s.dirtySet = ecs.NewDirtySet()

	s.Unit.Coords().AddDirtySet(s.dirtySet)
	s.Unit.Rotation().AddDirtySet(s.dirtySet)

	s.Transform.Pos().AddDependency(s.Unit.Coords())
	s.Transform.Size().AddDependency(s.Unit.Coords())
	s.Transform.Rotation().AddDependency(s.Unit.Rotation())

	s.Render.Mesh().BeforeGet(s.BeforeGet)
	s.Render.Texture().BeforeGet(s.BeforeGet)
	s.Transform.Pos().BeforeGet(s.BeforeGet)
	s.Transform.Size().BeforeGet(s.BeforeGet)
	s.Transform.Rotation().BeforeGet(s.BeforeGet)

	s.Collider.Component().BeforeGet(s.BeforeGet)
	s.Inputs.LeftClick().BeforeGet(s.BeforeGet)
	s.Inputs.Stack().BeforeGet(s.BeforeGet)

	events.Listen(s.EventsBuilder, s.OnTick)
	events.Listen(s.EventsBuilder, s.OnClick)

	return nil
}

func (s *system) BeforeGet() {
	for _, entity := range s.dirtySet.Get() {
		construct, ok := s.Unit.Unit().Get(entity)
		if !ok {
			continue
		}
		coords, ok := s.Unit.Coords().Get(entity)
		if !ok {
			continue
		}
		rotation, _ := s.Unit.Rotation().Get(entity)

		posFloor := s.Tile.GetPos(grid.NewCoords(
			grid.Coord(math.Floor(float64(coords.X))),
			grid.Coord(math.Floor(float64(coords.Y))),
		))
		posCeil := s.Tile.GetPos(grid.NewCoords(
			grid.Coord(math.Ceil(float64(coords.X))),
			grid.Coord(math.Ceil(float64(coords.Y))),
		))
		_, fractX := math.Modf(float64(coords.X))
		_, fractY := math.Modf(float64(coords.Y))
		pos := transform.NewPos(
			transition.Lerp(posFloor.Pos[0], posCeil.Pos[0], float32(fractX)),
			transition.Lerp(posFloor.Pos[1], posCeil.Pos[1], float32(fractY)),
			s.layer,
		)

		s.Render.Mesh().Set(entity, render.NewMesh(s.GameAssets.SquareMesh))
		s.Render.Texture().Set(entity, render.NewTexture(construct.Unit))

		s.Transform.ParentPivotPoint().Set(entity, transform.NewParentPivotPoint(0, 0, .5))
		s.Transform.Pos().Set(entity, pos)
		s.Transform.Size().Set(entity, s.Tile.GetTileSize())
		s.Transform.Rotation().Set(entity, transform.NewRotation(rotation.Quat()))

		s.Collider.Component().Set(entity, collider.NewCollider(s.GameAssets.SquareCollider))
		s.Inputs.LeftClick().Set(entity, inputs.NewLeftClick(unit.NewClickEvent(entity)))
		s.Inputs.Stack().Set(entity, inputs.StackComponent{})
	}
}

func (s *system) OnTick(e frames.TickEvent) {
	for _, unit := range s.Unit.Coords().GetEntities() {
		coords, _ := s.Unit.Coords().Get(unit)
		rot, _ := s.Unit.Rotation().Get(unit)

		coords.X += 1
		coords.Y += 1
		rot.Radians += mgl32.DegToRad(90)

		s.Unit.Coords().Set(unit, coords)
		s.Unit.Rotation().Set(unit, rot)
	}
}

func (s *system) OnClick(e unit.ClickEvent) {
	coords, ok := s.Unit.Coords().Get(e.Unit)
	if !ok {
		s.Logger.Warn(errors.New("expected unit to have unit coords component"))
		return
	}
	for _, p := range s.Ui.Show() {
		entity := s.NewEntity()
		s.Hierarchy.SetParent(entity, p)
		s.Transform.Parent().Set(entity, transform.NewParent(transform.RelativePos|transform.RelativeSizeXYZ))
		s.Groups.Inherit().Set(entity, groups.InheritGroupsComponent{})

		s.Text.Content().Set(entity, text.TextComponent{Text: fmt.Sprintf("UNIT: %v", coords)})
		s.Text.FontSize().Set(entity, text.FontSizeComponent{FontSize: 25})
		s.Text.Align().Set(entity, text.TextAlignComponent{Vertical: .5, Horizontal: .5})
	}
}
