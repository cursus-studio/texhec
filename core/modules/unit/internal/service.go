package internal

import (
	"core/modules/definitions"
	"core/modules/tile"
	"core/modules/unit"
	"engine"
	"engine/modules/grid"
	"engine/modules/render"
	"engine/modules/transform"
	"engine/services/ecs"
	"math"

	"github.com/ogiusek/ioc/v2"
	"golang.org/x/exp/constraints"
)

type service struct {
	engine.World `inject:"1"`
	Tile         tile.Service            `inject:"1"`
	GameAssets   definitions.Definitions `inject:"1"`

	layer    float32
	dirtySet ecs.DirtySet

	units      ecs.ComponentsArray[unit.UnitComponent]
	unitCoords ecs.ComponentsArray[unit.CoordsComponent]
}

func NewService(c ioc.Dic, layer float32) unit.Service {
	s := ioc.GetServices[*service](c)

	s.layer = layer
	s.dirtySet = ecs.NewDirtySet()

	s.units = ecs.GetComponentsArray[unit.UnitComponent](s)
	s.unitCoords = ecs.GetComponentsArray[unit.CoordsComponent](s)

	s.unitCoords.AddDirtySet(s.dirtySet)

	s.Transform.Pos().AddDependency(s.unitCoords)
	s.Transform.Size().AddDependency(s.unitCoords)

	s.Render.Mesh().BeforeGet(s.BeforeGet)
	s.Render.Texture().BeforeGet(s.BeforeGet)
	s.Transform.Pos().BeforeGet(s.BeforeGet)
	s.Transform.Size().BeforeGet(s.BeforeGet)

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
		coords, ok := s.unitCoords.Get(entity)
		if !ok {
			continue
		}

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

	}
}

func (s *service) Unit() ecs.ComponentsArray[unit.UnitComponent] {
	return s.units
}
func (s *service) Coords() ecs.ComponentsArray[unit.CoordsComponent] {
	return s.unitCoords
}
