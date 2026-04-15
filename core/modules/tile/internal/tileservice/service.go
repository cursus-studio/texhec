package tileservice

import (
	"core/modules/definitions"
	"core/modules/tile"
	"engine"
	"engine/modules/grid"
	"engine/modules/relation"
	"engine/modules/transform"
	"engine/services/ecs"

	"github.com/ogiusek/ioc/v2"
	"golang.org/x/exp/constraints"
)

type service struct {
	engine.World           `inject:"1"`
	TileGridService        grid.Service[tile.ID]          `inject:"1"`
	ObstructionGridService grid.Service[tile.Obstruction] `inject:"1"`
	TileTypeRelation       relation.Service[tile.ID]      `inject:"1"`

	tile ecs.ComponentsArray[tile.TypeComponent]

	pos   ecs.ComponentsArray[tile.PosComponent]
	size  ecs.ComponentsArray[tile.SizeComponent]
	rot   ecs.ComponentsArray[tile.RotComponent]
	layer ecs.ComponentsArray[tile.LayerComponent]

	placeholder ecs.ComponentsArray[tile.PlaceholderComponent]

	obstruction ecs.ComponentsArray[tile.ObstructionComponent]
	deployed    ecs.ComponentsArray[tile.DeployedComponent]

	speed ecs.ComponentsArray[tile.SpeedComponent]
	step  ecs.ComponentsArray[tile.StepComponent]
}

func NewService(c ioc.Dic) tile.Service {
	s := ioc.GetServices[*service](c)
	s.tile = ecs.GetComponentsArray[tile.TypeComponent](s.World)

	s.pos = ecs.GetComponentsArray[tile.PosComponent](s.World)
	s.size = ecs.GetComponentsArray[tile.SizeComponent](s.World)
	s.rot = ecs.GetComponentsArray[tile.RotComponent](s.World)
	s.layer = ecs.GetComponentsArray[tile.LayerComponent](s.World)

	s.placeholder = ecs.GetComponentsArray[tile.PlaceholderComponent](s.World)

	s.obstruction = ecs.GetComponentsArray[tile.ObstructionComponent](s.World)
	s.deployed = ecs.GetComponentsArray[tile.DeployedComponent](s.World)

	s.speed = ecs.GetComponentsArray[tile.SpeedComponent](s.World)
	s.step = ecs.GetComponentsArray[tile.StepComponent](s.World)

	s.size.SetEmpty(tile.NewSize(1, 1))
	s.layer.SetEmpty(tile.NewLayer(definitions.TileLayer))
	s.obstruction.SetEmpty(tile.NewObstruction(definitions.LowlandObstruction))

	return s
}

func (s *service) TileType() ecs.ComponentsArray[tile.TypeComponent] {
	return s.tile
}
func (s *service) TileGrid() ecs.ComponentsArray[grid.SquareGridComponent[tile.ID]] {
	return s.TileGridService.Component()
}
func (s *service) ObstructionGrid() ecs.ComponentsArray[grid.SquareGridComponent[tile.Obstruction]] {
	return s.ObstructionGridService.Component()
}
func (s *service) GetTileType(id tile.ID) (ecs.EntityID, bool) {
	return s.TileTypeRelation.Get(id)
}

func (s *service) Pos() ecs.ComponentsArray[tile.PosComponent]     { return s.pos }
func (s *service) Size() ecs.ComponentsArray[tile.SizeComponent]   { return s.size }
func (s *service) Rot() ecs.ComponentsArray[tile.RotComponent]     { return s.rot }
func (s *service) Layer() ecs.ComponentsArray[tile.LayerComponent] { return s.layer }

func (s *service) Placeholder() ecs.ComponentsArray[tile.PlaceholderComponent] { return s.placeholder }

func (s *service) Obstruction() ecs.ComponentsArray[tile.ObstructionComponent] { return s.obstruction }
func (s *service) Deployed() ecs.ComponentsArray[tile.DeployedComponent]       { return s.deployed }

func (s *service) Speed() ecs.ComponentsArray[tile.SpeedComponent] { return s.speed }
func (s *service) Step() ecs.ComponentsArray[tile.StepComponent]   { return s.step }

func (s *service) GetPos(coords grid.Coords) transform.PosComponent {
	size := s.GetTileSize().Size
	return transform.NewPos(
		size.X()*(float32(coords.X)+.5),
		size.Y()*(float32(coords.Y)+.5),
		size.Z(),
	)
}
func (s *service) GetTileSize() transform.SizeComponent {
	return transform.NewSize(100, 100, 1)
}

func (s *service) Collisions(aabb tile.AABB, obstruction tile.Obstruction) []grid.Coords {
	var collisions []grid.Coords
	obstructionGridEntity := s.ObstructionGrid().GetEntities()[0]
	obstructed, ok := s.ObstructionGrid().Get(obstructionGridEntity)
	if !ok {
		collisions = append(collisions, aabb.Tiles...)
		return collisions
	}
	for _, coords := range aabb.Tiles {
		index, ok := obstructed.GetIndex(coords.Coords())
		if !ok || obstruction&obstructed.GetTile(index) != 0 {
			collisions = append(collisions, coords)
			continue
		}
	}
	return collisions
}

func abs[Number constraints.Float | constraints.Integer](n Number) Number {
	return max(-n, n)
}

func (s *service) CanStep(
	pos tile.PosComponent,
	size tile.SizeComponent,
	obstruction tile.ObstructionComponent,
	step tile.StepComponent,
) bool {
	isValidStep := abs(step.X-grid.Coord(pos.X))+abs(step.Y-grid.Coord(pos.Y)) == 1
	if !isValidStep {
		return false
	}

	// is step destination occupied
	var aabbPos tile.PosComponent
	var aabbSize tile.SizeComponent

	// aabb size
	if grid.Coord(pos.X) != step.X {
		aabbSize = tile.NewSize(1, size.Y)
	} else if grid.Coord(pos.Y) != step.Y {
		aabbSize = tile.NewSize(size.X, 1)
	}
	// aabb pos
	if grid.Coord(pos.X) < step.X {
		aabbPos = tile.NewPos(step.X+size.X-1, step.Y)
	} else if grid.Coord(pos.Y) < step.Y {
		aabbPos = tile.NewPos(step.X, step.Y+size.Y-1)
	} else {
		aabbPos = tile.NewPos(step.Coords.Coords())
	}
	// perform is step destination occupied
	if collisions := s.Collisions(tile.NewAABB(aabbPos, aabbSize), obstruction.Obstruction); len(collisions) != 0 {
		return false
	}
	return true
}
