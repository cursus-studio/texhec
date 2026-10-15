package internal

import (
	"core/modules/actions"
	"core/modules/definitions"
	"core/modules/obstruction"
	"core/modules/tile"
	"engine/modules/ecs"
	"engine/modules/feature"
	"engine/modules/grid"
	"engine/modules/render"
	"engine/modules/transform"
	"engine/modules/uuid"

	"github.com/go-gl/mathgl/mgl32"
)

func (s *service) OnCoordsHover(e grid.HoverEvent) {
	chunkCoords, ok := s.EngineWorld.Grid().Coords().Get(e.Chunk)
	if !ok {
		return
	}
	coords := s.EngineWorld.Grid().AbsoluteCoords(chunkCoords, e.Coords)
	s.Feature().PreviewInteraction(actions.Coords(coords))
}
func (s *service) OnCoordsClick(e grid.ClickEvent) {
	chunkCoords, ok := s.EngineWorld.Grid().Coords().Get(e.Chunk)
	if !ok {
		return
	}
	coords := s.EngineWorld.Grid().AbsoluteCoords(chunkCoords, e.Coords)
	s.Feature().SubmitInteraction(actions.Coords(coords))
}

func (s *service) OnCanMoveToUpsert(entity ecs.EntityID) {
	var movingEntity uuid.UUID
	var coords grid.Coords
	var ruleRender feature.RuleRender

	if comp, ok := s.CanMoveTo().Render().Get(entity); ok &&
		(comp.Render == feature.RuleRenderPreviewRulesPass ||
			comp.Render == feature.RuleRenderCompletePreview ||
			comp.Render == feature.RuleRenderPreviewRulesFail) {
		movingEntity = uuid.UUID(comp.Rule.Entity.From(comp.Fields))
		coords = grid.Coords(comp.Rule.Coords.From(comp.Fields))
		ruleRender = comp.Render
		goto render
	}
	return
render:
	selectedEntity, ok := s.UUID().Entity(movingEntity)
	if !ok {
		return
	}
	blueprint, ok := s.Tile().Blueprint().Get(selectedEntity)
	if !ok {
		return
	}
	uuidComp, ok := s.UUID().Component().Get(blueprint)
	if !ok {
		return
	}

	pos := tile.NewPos(coords.Coords())
	size, _ := s.Tile().Size().Get(blueprint)

	s.Tile().Blueprint().SetUUID(entity, uuidComp.ID)
	s.Render().Texture().Set(entity, render.NewTexture(s.Definitions().Hud().Can))
	s.Tile().Layer().Set(entity, tile.NewLayer(definitions.ObjectPlaceholderLayer))
	s.Tile().Pos().Set(entity, pos)

	if ruleRender == feature.RuleRenderPreviewRulesFail {
		s.Render().Color().Set(entity, render.NewColor(mgl32.Vec4{1, 0, 0, 1}))
	} else {
		s.Render().Color().Set(entity, render.NewColor(mgl32.Vec4{0, 1, 0, 1}))
	}

	blueprintObstruction, _ := s.Obstruction().Component().Get(blueprint)
	aabb := obstruction.NewAABB(pos, size)
	collisions := s.Obstruction().Collisions(aabb, blueprintObstruction.Obstruction)
	for _, collision := range collisions {
		ind := s.World().NewEntity()
		s.Hierarchy().SetParent(ind, entity)
		s.Transform().Inherit().Set(ind, transform.NewInherit(transform.Absolute))

		s.Tile().Layer().Set(ind, tile.NewLayer(definitions.TilePlaceholderLayer))
		s.Render().Mesh().Set(ind, render.NewMesh(s.Definitions().Assets().SquareMesh))
		s.Render().Texture().Set(ind, render.NewTexture(s.Definitions().Assets().Blank))
		s.Groups().InheritGroups(ind)

		s.Tile().Pos().Set(ind, tile.NewPos(collision.Coords()))
		s.Render().Color().Set(ind, render.NewColor(mgl32.Vec4{1, 0, 0, 1}))
	}
}
func (s *service) OnCanDeployAtUpsert(entity ecs.EntityID) {
	var blueprintEntity uuid.UUID
	var coords grid.Coords
	var ruleRender feature.RuleRender
	if comp, ok := s.CanDeployAt().Render().Get(entity); ok &&
		(comp.Render == feature.RuleRenderPreviewRulesPass ||
			comp.Render == feature.RuleRenderCompletePreview ||
			comp.Render == feature.RuleRenderPreviewRulesFail) {
		blueprintEntity = uuid.UUID(comp.Rule.Blueprint.From(comp.Fields))
		coords = grid.Coords(comp.Rule.Coords.From(comp.Fields))
		ruleRender = comp.Render
		goto render
	}
	return
render:
	blueprint, ok := s.UUID().Entity(blueprintEntity)
	if !ok {
		return
	}
	pos := tile.NewPos(coords.Coords())
	size, _ := s.Tile().Size().Get(blueprint)

	s.Tile().Blueprint().SetUUID(entity, blueprintEntity)
	s.Tile().Layer().Set(entity, tile.NewLayer(definitions.ObjectPlaceholderLayer))
	s.Tile().Pos().Set(entity, pos)
	if ruleRender == feature.RuleRenderPreviewRulesFail {
		s.Render().Color().Set(entity, render.NewColor(mgl32.Vec4{1, 0, 0, 1}))
	} else {
		s.Render().Color().Set(entity, render.NewColor(mgl32.Vec4{0, 1, 0, 1}))
	}

	blueprintObstruction, _ := s.Obstruction().Component().Get(blueprint)
	aabb := obstruction.NewAABB(pos, size)
	collisions := s.Obstruction().Collisions(aabb, blueprintObstruction.Obstruction)
	for _, collision := range collisions {
		ind := s.World().NewEntity()
		s.Hierarchy().SetParent(ind, entity)
		s.Transform().Inherit().Set(ind, transform.NewInherit(transform.Absolute))

		s.Tile().Layer().Set(ind, tile.NewLayer(definitions.TilePlaceholderLayer))
		s.Render().Mesh().Set(ind, render.NewMesh(s.Definitions().Assets().SquareMesh))
		s.Render().Texture().Set(ind, render.NewTexture(s.Definitions().Assets().Blank))
		s.Groups().InheritGroups(ind)

		s.Tile().Pos().Set(ind, tile.NewPos(collision.Coords()))
		s.Render().Color().Set(ind, render.NewColor(mgl32.Vec4{1, 0, 0, 1}))
	}
}
