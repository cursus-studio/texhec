package internal

import (
	"core/modules/definitions"
	"core/modules/tile"
	"engine/modules/ecs"
	"engine/modules/feature"
	"engine/modules/grid"
	"engine/modules/groups"
	"engine/modules/render"
	"engine/modules/transform"
	"engine/modules/uuid"

	"github.com/go-gl/mathgl/mgl32"
)

func (s *service) OnReachUpsert(entity ecs.EntityID) {
	var gridCoords []grid.Coords
	var color mgl32.Vec4
	if comp, ok := s.CanAttack().Render().Get(entity); ok &&
		(comp.Render == feature.RuleRenderPreviewRulesPass ||
			comp.Render == feature.RuleRenderCompletePreview) {
		selectedUUID := uuid.UUID(comp.Rule.Entity.From(comp.Fields))
		selectedEntity, ok := s.UUID().Entity(selectedUUID)
		if !ok {
			return
		}
		gridCoords = s.GameWorld.Attack().Reach().TilesWithinReach(selectedEntity)
		color = mgl32.Vec4{.5, 0, 0, 1}
		goto render
	}
	if comp, ok := s.CanDeploy().Render().Get(entity); ok &&
		(comp.Render == feature.RuleRenderPreviewRulesPass ||
			comp.Render == feature.RuleRenderCompletePreview) {
		selectedUUID := uuid.UUID(comp.Rule.Entity.From(comp.Fields))
		selectedEntity, ok := s.UUID().Entity(selectedUUID)
		if !ok {
			return
		}
		gridCoords = s.GameWorld.Deploy().Reach().TilesWithinReach(selectedEntity)
		color = mgl32.Vec4{0, 0, .5, 1}
		goto render
	}
	return
render:
	for _, reachCoords := range gridCoords {
		ind := s.World().NewEntity()
		s.Hierarchy().SetParent(ind, entity)
		s.Transform().Inherit().Set(ind, transform.NewInherit(transform.Absolute))

		s.Render().Mesh().Set(ind, render.NewMesh(s.Definitions().Assets().SquareMesh))
		s.Render().Texture().Set(ind, render.NewTexture(s.Definitions().Assets().Border))
		s.Groups().Component().Set(ind, groups.EmptyGroups().Enable(definitions.GameGroup))

		s.Tile().Layer().Set(ind, tile.NewLayer(definitions.RangePlaceholderLayer))
		s.Tile().Pos().Set(ind, tile.NewPos(reachCoords.Coords()))
		s.Render().Color().Set(ind, render.NewColor(color))
	}
}
