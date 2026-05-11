package internal

import (
	"core/game"
	"core/modules/definitions"
	"core/modules/deploy"
	"core/modules/player"
	"core/modules/tile"
	"engine/modules/grid"
	"engine/modules/groups"
	"engine/modules/inputs"
	"engine/modules/render"
	"engine/services/ecs"
	"slices"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type PreviewedComponent struct {
	Event      deploy.PreviewEvent
	Collisions []grid.Coords
}

func (c1 *PreviewedComponent) Equal(c2 PreviewedComponent) bool {
	return c1.Event == c2.Event && slices.Equal(c1.Collisions, c2.Collisions)
}

type service struct {
	game.GameWorld `inject:""`

	component ecs.ComponentsArray[deploy.Component]
	previewed ecs.ComponentsArray[PreviewedComponent]
}

func NewService(c ioc.Dic) deploy.Service {
	s := ioc.GetServices[*service](c)

	s.component = ecs.GetComponentsArray[deploy.Component](s.World())
	s.previewed = ecs.GetComponentsArray[PreviewedComponent](s.World())

	events.Listen(s.EventsBuilder(), s.Execute)
	events.Listen(s.EventsBuilder(), s.Preview)
	events.Listen(s.EventsBuilder(), s.Select)

	return s
}

func (s *service) Component() ecs.ComponentsArray[deploy.Component] { return s.component }

func (s *service) Deploy(
	blueprint,
	owner ecs.EntityID,
	coords grid.Coords,
) (ecs.EntityID, error) {
	s.Ui().Objects().Unselect()
	// check can place:

	// - is position occuped
	pos := tile.NewPos(coords.Coords())
	size, _ := s.Tile().Size().Get(blueprint)
	obstruction, _ := s.Tile().Obstruction().Get(blueprint)
	aabb := tile.NewAABB(pos, size)
	if collisions := s.Tile().Collisions(aabb, obstruction.Obstruction); len(collisions) != 0 {
		return 0, tile.ErrPositionIsOccupied
	}

	// place
	deployed := s.Prototype().Clone(blueprint)
	s.Hierarchy().SetParent(deployed, s.Scene().Scene())

	s.Player().Owner().Set(deployed, player.NewOwner(owner))
	s.Tile().Deployed().Set(deployed, tile.NewDeployed())
	s.Inputs().LeftClick().Set(deployed, inputs.NewLeftClick(tile.NewClickEntityEvent()))
	s.Tile().Pos().Set(deployed, pos)
	return deployed, nil
}

func (s *service) Select(e deploy.SelectEvent) {
	events.Emit(s.Events(), tile.NewSelectEvent(deploy.NewPreviewEvent(e.By, e.Blueprint)))
}
func (s *service) Preview(e deploy.PreviewEvent) {
	pos := tile.NewPos(e.Coords.Coords())
	blueprintObstruction, _ := s.Tile().Obstruction().Get(e.Blueprint)
	size, _ := s.Tile().Size().Get(e.Blueprint)
	aabb := tile.NewAABB(pos, size)
	previewComp := PreviewedComponent{
		Event:      e,
		Collisions: s.Tile().Collisions(aabb, blueprintObstruction.Obstruction),
	}

	if len(s.previewed.GetEntities()) == 1 {
		previewed := s.previewed.GetEntities()[0]
		previousPreviewComp, _ := s.previewed.Get(previewed)
		if previousPreviewComp.Equal(previewComp) {
			return
		}
	}

	s.Ui().Actions().Unselect()

	placeholderEntity := s.Prototype().Clone(e.Blueprint)
	s.Hierarchy().SetParent(placeholderEntity, s.Scene().Scene())
	s.Tile().Layer().Set(placeholderEntity, tile.NewLayer(definitions.ObjectPlaceholderLayer))
	s.Tile().Pos().Set(placeholderEntity, pos)
	s.Ui().Actions().HideOnUnselect(placeholderEntity)
	s.Inputs().KeepSelected().Set(placeholderEntity, inputs.KeepSelectedComponent{})
	s.previewed.Set(placeholderEntity, previewComp)

	if len(previewComp.Collisions) == 0 {
		s.Render().Color().Set(placeholderEntity, render.NewColor(mgl32.Vec4{0, 1, 0, 1}))
		s.Inputs().LeftClick().Set(placeholderEntity, inputs.NewLeftClick(deploy.NewExecuteEvent(e.By, e.Blueprint).ApplyCoords(e.Coords)))
		return
	}

	// place indicator on occupied tiles
	for _, collision := range previewComp.Collisions {
		entity := s.Prototype().Clone(s.Definitions().Assets().Blank)
		s.Hierarchy().SetParent(entity, s.Scene().Scene())

		s.Tile().Layer().Set(entity, tile.NewLayer(definitions.TilePlaceholderLayer))
		s.Render().Mesh().Set(entity, render.NewMesh(s.Definitions().Assets().SquareMesh))
		s.Render().Texture().Set(entity, render.NewTexture(s.Definitions().Assets().Blank))
		s.Groups().Component().Set(entity, groups.EmptyGroups().Ptr().Enable(definitions.GameGroup).Val())

		s.Tile().Layer().Set(entity, tile.NewLayer(definitions.TilePlaceholderLayer))
		s.Tile().Pos().Set(entity, tile.NewPos(collision.Coords()))
		s.Ui().Actions().HideOnUnselect(entity)
		s.Render().Color().Set(entity, render.NewColor(mgl32.Vec4{1, 0, 0, 1}))
	}
	s.Render().Color().Set(placeholderEntity, render.NewColor(mgl32.Vec4{1, 0, 0, 1}))
}
func (s *service) Execute(e deploy.ExecuteEvent) {
	s.Ui().Objects().Unselect()

	// pay
	// ...

	// check can place
	pos := tile.NewPos(e.Coords.Coords())
	size, _ := s.Tile().Size().Get(e.Blueprint)
	blueprintObstruction, _ := s.Tile().Obstruction().Get(e.Blueprint)

	obstructionGridEntity := s.Tile().ObstructionGrid().GetEntities()[0]
	obstructed, ok := s.Tile().ObstructionGrid().Get(obstructionGridEntity)
	if !ok {
		s.Logger().Log(tile.ErrPositionIsOccupied)
		return
	}
	aabb := tile.NewAABB(pos, size)
	for _, coords := range aabb.Tiles {
		index, ok := obstructed.GetIndex(coords.Coords())
		if !ok {
			s.Logger().Log(tile.ErrPositionIsOccupied)
			return
		}
		coordsObstruction := obstructed.GetTile(index)
		if blueprintObstruction.Obstruction&coordsObstruction != 0 {
			s.Logger().Log(tile.ErrPositionIsOccupied)
			return
		}
	}

	// place
	deployed := s.Prototype().Clone(e.Blueprint)
	s.Hierarchy().SetParent(deployed, s.Scene().Scene())
	if owner, ok := s.Player().Owner().Get(e.By); ok {
		s.Player().Owner().Set(deployed, owner)
	}
	s.Tile().Deployed().Set(deployed, tile.NewDeployed())
	s.Inputs().LeftClick().Set(deployed, inputs.NewLeftClick(tile.NewClickEntityEvent()))
	s.Tile().Pos().Set(deployed, tile.NewPos(e.Coords.Coords()))
}
