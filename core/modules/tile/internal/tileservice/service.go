package tileservice

import (
	"core/modules/definitions"
	"core/modules/tile"
	"engine/modules/collider"
	"engine/modules/grid"
	"engine/modules/inputs"
	"engine/modules/render"
	"engine/modules/transform"
	"engine/services/ecs"

	"github.com/ogiusek/ioc/v2"
)

type service struct {
	C                     ioc.Dic
	definitions           *definitions.Definitions
	World                 ecs.World        `inject:"1"`
	Render                render.Service   `inject:"1"`
	Collider              collider.Service `inject:"1"`
	Inputs                inputs.Service   `inject:"1"`
	grid.Service[tile.ID] `inject:"1"`

	tile ecs.ComponentsArray[tile.Component]

	pos   ecs.ComponentsArray[tile.PosComponent]
	size  ecs.ComponentsArray[tile.SizeComponent]
	rot   ecs.ComponentsArray[tile.RotComponent]
	layer ecs.ComponentsArray[tile.LayerComponent]
}

func NewService(c ioc.Dic) tile.Service {
	s := ioc.GetServices[*service](c)
	s.C = c
	s.tile = ecs.GetComponentsArray[tile.Component](s.World)

	s.pos = ecs.GetComponentsArray[tile.PosComponent](s.World)
	s.size = ecs.GetComponentsArray[tile.SizeComponent](s.World)
	s.rot = ecs.GetComponentsArray[tile.RotComponent](s.World)
	s.layer = ecs.GetComponentsArray[tile.LayerComponent](s.World)

	s.size.SetEmpty(tile.NewSize(1, 1))
	s.layer.SetEmpty(tile.NewLayer(1))

	return s
}

func (t *service) Definitions() *definitions.Definitions {
	if t.definitions == nil {
		definitions := ioc.Get[definitions.Definitions](t.C)
		t.definitions = &definitions
	}
	return t.definitions
}

func (t *service) Tile() ecs.ComponentsArray[tile.Component] {
	return t.tile
}
func (t *service) Grid() ecs.ComponentsArray[grid.SquareGridComponent[tile.ID]] {
	return t.Component()
}

func (t *service) Pos() ecs.ComponentsArray[tile.PosComponent]     { return t.pos }
func (t *service) Size() ecs.ComponentsArray[tile.SizeComponent]   { return t.size }
func (t *service) Rot() ecs.ComponentsArray[tile.RotComponent]     { return t.rot }
func (t *service) Layer() ecs.ComponentsArray[tile.LayerComponent] { return t.layer }

func (t *service) GetPos(coords grid.Coords) transform.PosComponent {
	size := t.GetTileSize().Size
	return transform.NewPos(
		size.X()*(float32(coords.X)+.5),
		size.Y()*(float32(coords.Y)+.5),
		size.Z(),
	)
}
func (t *service) GetTileSize() transform.SizeComponent {
	return transform.NewSize(100, 100, 1)
}

func (s *service) Unit(entity, blueprint ecs.EntityID) {
	s.Layer().Set(entity, tile.NewLayer(3))

	s.Render.Mesh().Set(entity, render.NewMesh(s.Definitions().SquareMesh))
	s.Render.Texture().Set(entity, render.NewTexture(blueprint))

	s.Collider.Component().Set(entity, collider.NewCollider(s.Definitions().SquareCollider))
	s.Inputs.LeftClick().Set(entity, inputs.NewLeftClick(tile.NewClickObjectEvent()))
	s.Inputs.Stack().Set(entity, inputs.StackComponent{})
}

func (s *service) Construct(entity, blueprint ecs.EntityID) {
	s.Render.Mesh().Set(entity, render.NewMesh(s.Definitions().SquareMesh))
	s.Render.Texture().Set(entity, render.NewTexture(blueprint))

	s.Collider.Component().Set(entity, collider.NewCollider(s.Definitions().SquareCollider))
	s.Inputs.LeftClick().Set(entity, inputs.NewLeftClick(tile.NewClickObjectEvent()))
	s.Inputs.Stack().Set(entity, inputs.StackComponent{})
}
