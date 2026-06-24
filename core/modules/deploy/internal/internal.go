package internal

import (
	"core/game"
	"core/modules/deploy"
	"core/modules/obstruction"
	"core/modules/player"
	"core/modules/reach"
	"core/modules/tile"
	"engine/modules/ecs"
	"engine/modules/grid"
	"engine/modules/inputs"
	"engine/modules/seed"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type service struct {
	game.GameWorld `inject:""`
	ReachT         reach.ServiceT[deploy.Component] `inject:""`

	component ecs.ComponentArray[deploy.Component]
}

func NewService(c ioc.Dic) deploy.Service {
	s := ioc.GetServices[*service](c)
	s.ReachT.Component().SetEmpty(reach.NewReach[deploy.Component](1))

	s.component = ecs.GetComponentArray[deploy.Component](s.World())

	events.Listen(s.EventsBuilder(), s.DeployEvent)

	return s
}

func (s *service) Component() ecs.ComponentArray[deploy.Component] { return s.component }
func (s *service) Reach() reach.ServiceT[deploy.Component]         { return s.ReachT }

func (s *service) Deploy(
	blueprint,
	owner ecs.EntityID,
	coords grid.Coords,
) (ecs.EntityID, error) {
	worldEntity, ok := s.Seed().WorldSeed()
	if !ok {
		return 0, seed.ErrWorldCanHaveOneSeed
	}
	// check can place:

	// - is position occuped
	pos := tile.NewPos(coords.Coords())
	size, _ := s.Tile().Size().Get(blueprint)
	obstructionComp, _ := s.Obstruction().Component().Get(blueprint)
	aabb := obstruction.NewAABB(pos, size)
	if collisions := s.Obstruction().Collisions(aabb, obstructionComp.Obstruction); len(collisions) != 0 {
		return 0, obstruction.ErrPositionIsOccupied
	}

	// place
	deployed := s.Prototype().Clone(blueprint)
	s.Hierarchy().SetParent(deployed, worldEntity)

	s.Player().Owner().Set(deployed, player.NewOwner(owner))
	s.Obstruction().Deployed().Set(deployed, obstruction.NewDeployed())
	s.Inputs().LeftClick().Set(deployed, inputs.NewLeftClick(tile.NewClickEntityEvent()))
	s.Tile().Pos().Set(deployed, pos)
	return deployed, nil
}

func (s *service) DeployEvent(e deploy.DeployEvent) {
	worldEntity, ok := s.Seed().WorldSeed()
	if !ok {
		return
	}

	// by
	byPos, ok := s.Tile().Pos().Get(e.By)
	if !ok {
		s.Logger().Log(obstruction.ErrPositionIsOccupied)
		return
	}
	bySize, _ := s.Tile().Size().Get(e.By)
	reachComp, _ := s.GameWorld.Deploy().Reach().Component().Get(e.By)

	// target
	pos := tile.NewPos(e.Coords.Coords())
	size, _ := s.Tile().Size().Get(e.Blueprint)

	// check can place
	{ // reach
		dist := s.GameWorld.Reach().Distance(byPos, bySize, pos, size)
		isInReach := dist <= tile.Coord(reachComp.Reach)
		if !isInReach {
			s.Logger().Log(reach.ErrOutsideOfReach)
			return
		}
	}
	{ // obstruction
		blueprintObstruction, _ := s.Obstruction().Component().Get(e.Blueprint)

		aabb := obstruction.NewAABB(pos, size)
		collisions := s.Obstruction().Collisions(aabb, blueprintObstruction.Obstruction)

		if len(collisions) != 0 {
			s.Logger().Log(obstruction.ErrPositionIsOccupied)
			return
		}
	}

	// pay
	// ...

	// place
	deployed := s.Prototype().Clone(e.Blueprint)
	s.Hierarchy().SetParent(deployed, worldEntity)
	if owner, ok := s.Player().Owner().Get(e.By); ok {
		s.Player().Owner().Set(deployed, owner)
	}
	s.Obstruction().Deployed().Set(deployed, obstruction.NewDeployed())
	s.Inputs().LeftClick().Set(deployed, inputs.NewLeftClick(tile.NewClickEntityEvent()))
	s.Tile().Pos().Set(deployed, tile.NewPos(e.Coords.Coords()))
}
