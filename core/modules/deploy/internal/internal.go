package internal

import (
	"core/game"
	"core/modules/deploy"
	"core/modules/economy"
	"core/modules/obstruction"
	"core/modules/player"
	"core/modules/reach"
	"core/modules/tile"
	"engine/modules/ecs"
	"engine/modules/grid"
	"engine/modules/loop"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type BoughtComponent deploy.DeployEvent

func NewBought(event deploy.DeployEvent) BoughtComponent { return BoughtComponent(event) }

//

type service struct {
	game.GameWorld `inject:""`
	ReachT         reach.ServiceT[deploy.Component] `inject:""`

	boughtComponent ecs.ComponentArray[BoughtComponent]
	component       ecs.ComponentArray[deploy.Component]
}

func NewService(c ioc.Dic) deploy.Service {
	s := ioc.GetServices[*service](c)
	s.ReachT.Component().SetEmpty(reach.NewReach[deploy.Component](1))

	s.boughtComponent = ecs.GetComponentArray[BoughtComponent](s.World())
	s.component = ecs.GetComponentArray[deploy.Component](s.World())

	events.Listen(s.EventsBuilder(), s.DeployEvent)
	events.Listen(s.EventsBuilder(), s.DestroyEvent)

	return s
}

func (s *service) Register() error {
	events.Listen(s.EventsBuilder(), s.OnTick)
	return nil
}

func (s *service) Component() ecs.ComponentArray[deploy.Component] { return s.component }
func (s *service) Reach() reach.ServiceT[deploy.Component]         { return s.ReachT }

func (s *service) Deploy(
	blueprint,
	owner ecs.EntityID,
	coords grid.Coords,
) (ecs.EntityID, error) {
	// check can place:

	// - is position occuped
	pos := tile.NewPos(coords.Coords())
	size, _ := s.Tile().Size().Get(blueprint)
	obstructionComp, _ := s.Obstruction().Component().Get(blueprint)
	aabb := obstruction.NewAABB(pos, size)
	if collisions := s.Obstruction().Collisions(aabb, obstructionComp.Obstruction); len(collisions) != 0 {
		return 0, obstruction.ErrPositionIsOccupied
	}

	blueprintUUID, ok := s.UUID().Component().Get(blueprint)
	if !ok {
		s.Logger().Fatal(tile.ErrBlueprintIsMissingUUID)
	}
	ownerUUID, ok := s.UUID().Component().Get(owner)
	if !ok {
		s.Logger().Fatal(player.ErrRequiresOwner)
	}

	// place
	deployed := s.World().NewEntity()
	s.Player().Owner().SetUUID(deployed, ownerUUID.ID)
	s.Obstruction().Deployed().Set(deployed, obstruction.NewDeployed())
	s.Tile().Link().SetUUID(deployed, blueprintUUID.ID)
	s.Tile().Pos().Set(deployed, pos)
	return deployed, nil
}

func (s *service) DeployEvent(e deploy.DeployEvent) {
	entity := s.World().NewEntity()
	s.boughtComponent.Set(entity, NewBought(e))
}
func (s *service) OnTick(loop.TickEvent) {
	entities := s.boughtComponent.GetEntities()
	for _, entity := range entities {
		event, ok := s.boughtComponent.Get(entity)
		if !ok {
			continue
		}
		s.World().RemoveEntity(entity)

		// by
		byPos, ok := s.Tile().Pos().Get(event.By)
		if !ok {
			s.Logger().Log(obstruction.ErrPositionIsOccupied)
			continue
		}
		bySize, _ := s.Tile().Size().Get(event.By)
		reachComp, _ := s.GameWorld.Deploy().Reach().Component().Get(event.By)

		// target
		pos := tile.NewPos(event.Coords.Coords())
		size, _ := s.Tile().Size().Get(event.Blueprint)

		// check can place
		{ // reach
			dist := s.GameWorld.Reach().Distance(byPos, bySize, pos, size)
			isInReach := dist <= tile.Coord(reachComp.Reach)
			if !isInReach {
				s.Logger().Log(reach.ErrOutsideOfReach)
				continue
			}
		}
		{ // obstruction
			blueprintObstruction, _ := s.Obstruction().Component().Get(event.Blueprint)

			aabb := obstruction.NewAABB(pos, size)
			collisions := s.Obstruction().Collisions(aabb, blueprintObstruction.Obstruction)

			if len(collisions) != 0 {
				s.Logger().Log(obstruction.ErrPositionIsOccupied)
				continue
			}
		}

		owner, ok := s.Player().Owner().Get(event.By)
		if !ok {
			s.Logger().Log(player.ErrRequiresOwner)
			continue
		}

		// pay
		if cost, ok := s.Economy().Cost().Get(event.Blueprint); ok {
			wallet, ok := s.Economy().Wallet().Get(owner)
			if !ok || cost.Cost > wallet.Money {
				s.Logger().Log(economy.ErrToExpensive)
				continue
			}
			s.Economy().Wallet().Set(owner, wallet.Pay(cost))
		}

		blueprintUUID, ok := s.UUID().Component().Get(event.Blueprint)
		if !ok {
			s.Logger().Fatal(tile.ErrBlueprintIsMissingUUID)
		}
		ownerUUID, ok := s.UUID().Component().Get(event.By)
		if !ok {
			s.Logger().Fatal(player.ErrRequiresOwner)
		}

		// place
		deployed := s.World().NewEntity()
		s.Player().Owner().SetUUID(deployed, ownerUUID.ID)
		s.Obstruction().Deployed().Set(deployed, obstruction.NewDeployed())
		s.Tile().Link().SetUUID(deployed, blueprintUUID.ID)
		s.Tile().Pos().Set(deployed, tile.NewPos(event.Coords.Coords()))
	}
}

func (s *service) DestroyEvent(e deploy.DestroyEvent) {
	s.World().RemoveEntity(e.Entity)
}
