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
	"engine/services/ecs"
	"errors"
	"fmt"
	"math"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
	"golang.org/x/exp/constraints"
)

type service struct {
	engine.World `inject:"1"`
	Tile         tile.Service            `inject:"1"`
	GameAssets   definitions.Definitions `inject:"1"`
	Ui           ui.Service              `inject:"1"`

	layer    float32
	dirtySet ecs.DirtySet

	units     ecs.ComponentsArray[unit.UnitComponent]
	coords    ecs.ComponentsArray[unit.CoordsComponent]
	rotations ecs.ComponentsArray[unit.RotationComponent]
}

func NewService(c ioc.Dic, layer float32) unit.Service {
	s := ioc.GetServices[*service](c)

	s.layer = layer
	s.dirtySet = ecs.NewDirtySet()

	s.units = ecs.GetComponentsArray[unit.UnitComponent](s)
	s.coords = ecs.GetComponentsArray[unit.CoordsComponent](s)
	s.rotations = ecs.GetComponentsArray[unit.RotationComponent](s)

	s.coords.AddDirtySet(s.dirtySet)
	s.rotations.AddDirtySet(s.dirtySet)

	s.Transform.Pos().AddDependency(s.coords)
	s.Transform.Size().AddDependency(s.coords)
	s.Transform.Rotation().AddDependency(s.rotations)

	s.Render.Mesh().BeforeGet(s.BeforeGet)
	s.Render.Texture().BeforeGet(s.BeforeGet)
	s.Transform.Pos().BeforeGet(s.BeforeGet)
	s.Transform.Size().BeforeGet(s.BeforeGet)
	s.Transform.Rotation().BeforeGet(s.BeforeGet)

	s.Collider.Component().BeforeGet(s.BeforeGet)
	s.Inputs.LeftClick().BeforeGet(s.BeforeGet)
	s.Inputs.Stack().BeforeGet(s.BeforeGet)

	events.Listen(s.EventsBuilder, s.OnClick)

	return s
}

func lerp[Number constraints.Float](a, b, t Number) Number {
	return a*(t) + b*(1.-t)
}

func (s *service) BeforeGet() {
	for _, entity := range s.dirtySet.Get() {
		construct, ok := s.units.Get(entity)
		if !ok {
			continue
		}
		coords, ok := s.coords.Get(entity)
		if !ok {
			continue
		}
		rotation, _ := s.rotations.Get(entity)

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
			lerp(posFloor.Pos[0], posCeil.Pos[0], float32(fractX)),
			lerp(posFloor.Pos[1], posCeil.Pos[1], float32(fractY)),
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

func (s *service) OnClick(e unit.ClickEvent) {
	coords, ok := s.coords.Get(e.Unit)
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

func (s *service) Unit() ecs.ComponentsArray[unit.UnitComponent] {
	return s.units
}
func (s *service) Coords() ecs.ComponentsArray[unit.CoordsComponent] {
	return s.coords
}
func (s *service) Rotation() ecs.ComponentsArray[unit.RotationComponent] {
	return s.rotations
}
