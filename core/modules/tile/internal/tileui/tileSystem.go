package tileui

import (
	"core/modules/tile"
	"core/modules/ui"
	"engine"
	"engine/modules/groups"
	"engine/modules/text"
	"engine/modules/transform"
	"engine/services/ecs"
	"fmt"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type system struct {
	engine.World `inject:"1"`
	Ui           ui.Service   `inject:"1"`
	Tile         tile.Service `inject:"1"`

	dirtySet ecs.DirtySet

	TileSize transform.SizeComponent
}

func NewSystem(c ioc.Dic) tile.System {
	return ecs.NewSystemRegister(func() error {
		s := ioc.GetServices[*system](c)

		s.TileSize = s.Tile.GetTileSize()

		s.dirtySet = ecs.NewDirtySet()

		s.Tile.Pos().AddDirtySet(s.dirtySet)
		s.Tile.Size().AddDirtySet(s.dirtySet)
		s.Tile.Rot().AddDirtySet(s.dirtySet)

		s.Transform.Pos().AddDependency(s.Tile.Pos())
		s.Transform.Size().AddDependency(s.Tile.Size())
		s.Transform.Rotation().AddDependency(s.Tile.Rot())

		s.Transform.ParentPivotPoint().BeforeGet(s.BeforeGet)
		s.Transform.Pos().BeforeGet(s.BeforeGet)
		s.Transform.Size().BeforeGet(s.BeforeGet)
		s.Transform.Rotation().BeforeGet(s.BeforeGet)

		events.Listen(s.EventsBuilder, s.OnClick)
		return nil
	})
}

func (s *system) BeforeGet() {
	for _, entity := range s.dirtySet.Get() {
		pos, posOk := s.Tile.Pos().Get(entity)
		size, sizeOk := s.Tile.Size().Get(entity)
		rot, rotOk := s.Tile.Rot().Get(entity)
		layer, layerOk := s.Tile.Layer().Get(entity)
		if !posOk && !sizeOk && !rotOk && !layerOk {
			continue
		}

		transformPos := transform.NewPos(
			s.TileSize.Size.X()*(float32(pos.X)+.5),
			s.TileSize.Size.Y()*(float32(pos.Y)+.5),
			float32(layer.Z),
		)
		transformSize := transform.NewSize(
			s.TileSize.Size[0]*float32(size.X),
			s.TileSize.Size[1]*float32(size.Y),
			s.TileSize.Size[2],
		)
		transformRot := transform.NewRotation(rot.Quat())

		s.Transform.ParentPivotPoint().Set(entity, transform.NewParentPivotPoint(0, 0, .5))

		s.Transform.Pos().Set(entity, transformPos)
		s.Transform.Size().Set(entity, transformSize)
		s.Transform.Rotation().Set(entity, transformRot)
	}
}

func (s *system) OnClick(e tile.ClickEvent) {
	grid, ok := s.Tile.Grid().Get(e.Grid)
	if !ok {
		s.Logger.Warn(fmt.Errorf("grid doesn't exist"))
		return
	}
	coords := grid.GetCoords(e.Tile)
	if !ok {
		return
	}
	for _, p := range s.Ui.Show() {
		entity := s.NewEntity()
		s.Hierarchy.SetParent(entity, p)
		s.Transform.Parent().Set(entity, transform.NewParent(transform.RelativePos|transform.RelativeSizeXYZ))
		s.Groups.Inherit().Set(entity, groups.InheritGroupsComponent{})

		s.Text.Content().Set(entity, text.TextComponent{Text: fmt.Sprintf("TILE: %v", coords)})
		s.Text.FontSize().Set(entity, text.FontSizeComponent{FontSize: 25})
		s.Text.Align().Set(entity, text.TextAlignComponent{Vertical: .5, Horizontal: .5})
	}
}
