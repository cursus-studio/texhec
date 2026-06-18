package tilesystem

import (
	"core/game"
	"engine/modules/transform"
	"engine/services/ecs"

	"github.com/ogiusek/ioc/v2"
)

type system struct {
	game.GameWorld `inject:""`

	tileSize transform.SizeComponent
}

func NewSystem(c ioc.Dic) ecs.SystemRegister {
	return ecs.NewSystemRegister(func() error {
		s := ioc.GetServices[*system](c)

		s.tileSize = s.Tile().GetTileSize()

		//

		s.Tile().Pos().OnRemove(s.OnTilePosRemove)

		s.Tile().Pos().OnUpsert(s.OnTilePosSizeRotUpsert)
		s.Tile().Size().OnUpsert(s.OnTilePosSizeRotUpsert)
		s.Tile().Rot().OnUpsert(s.OnTilePosSizeRotUpsert)

		//

		return nil
	})
}

func (s *system) OnTilePosRemove(entity ecs.EntityID) {
	s.Transform().Size().Remove(entity)
	s.Inputs().Stack().Remove(entity)
}

func (s *system) OnTilePosSizeRotUpsert(entity ecs.EntityID) {
	pos, ok := s.Tile().Pos().Get(entity)
	if !ok {
		return
	}
	size, _ := s.Tile().Size().Get(entity)
	rot, _ := s.Tile().Rot().Get(entity)
	layer, _ := s.Tile().Layer().Get(entity)

	transformPos := transform.NewPos(
		s.tileSize.Size.X()*float32(pos.X),
		s.tileSize.Size.Y()*float32(pos.Y),
		float32(layer.Z),
	)
	transformSize := transform.NewSize(
		s.tileSize.Size[0]*float32(size.X),
		s.tileSize.Size[1]*float32(size.Y),
		s.tileSize.Size[2],
	)
	transformRot := transform.NewRotation(rot.Quat())

	s.Transform().PivotPoint().Set(entity, transform.NewPivotPoint(0, 0, .5))
	s.Transform().Pos().Set(entity, transformPos)
	s.Transform().Size().Set(entity, transformSize)
	s.Transform().Rotation().Set(entity, transformRot)
}

// V1:
// relations by feature:
// move  (affects coords cursor look): object(affects coords cursor size), coords
// build: object, coords, building(affects coords cursor size&look)

// interactions:
// object (can be first)    : interaction on object click
// coords (can't be first)  : interaction on hover if it's missing
//                            cursor look and size can be modified
// building (can't be first): interaction on button click

// V2:
// legend:
// ccc: configures coords cursor

// features:
// | name  | interactions             | ccc look | ccc size
// | move  | object, coords           | default  | object
// | build | object, coords, building | building | size

// interactions:
// | name     | can be first | config           | how works
// | object   | yes          | none             | on object click
// | coords   | no           | cursor look&size | on hover if it's missing
// | building | no           | none             | on button click

// game:
//   scene
//     map
//     feature
//   assets
