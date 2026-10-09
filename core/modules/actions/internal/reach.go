package internal

import (
	"core/modules/definitions"
	"core/modules/tile"
	"engine/modules/ecs"
	"engine/modules/feature"
	"engine/modules/groups"
	"engine/modules/render"
	"engine/modules/transform"
	"engine/modules/uuid"

	"github.com/go-gl/mathgl/mgl32"
)

func (s *service) OnCanAttackUpsert(entity ecs.EntityID) {
	var attacker uuid.UUID
	if comp, ok := s.CanAttack().Render().Get(entity); ok &&
		(comp.Render == feature.RuleRenderPreviewRulesPass ||
			comp.Render == feature.RuleRenderCompletePreview) {
		attacker = uuid.UUID(comp.Rule.Entity.From(comp.Fields))
		goto render
	}
	return
render:
	selectedEntity, ok := s.UUID().Entity(attacker)
	if !ok {
		return
	}

	for _, reachCoords := range s.GameWorld.Attack().Reach().TilesWithinReach(selectedEntity) {
		ind := s.World().NewEntity()
		s.Hierarchy().SetParent(ind, entity)
		s.Transform().Inherit().Set(ind, transform.NewInherit(transform.Absolute))

		s.Render().Mesh().Set(ind, render.NewMesh(s.Definitions().Assets().SquareMesh))
		s.Render().Texture().Set(ind, render.NewTexture(s.Definitions().Assets().Border))
		s.Groups().Component().Set(ind, groups.EmptyGroups().Enable(definitions.GameGroup))

		s.Tile().Layer().Set(ind, tile.NewLayer(definitions.RangePlaceholderLayer))
		s.Tile().Pos().Set(ind, tile.NewPos(reachCoords.Coords()))
		s.Render().Color().Set(ind, render.NewColor(mgl32.Vec4{.5, 0, 0, 1}))
	}
}
func (s *service) OnCanDeployUpsert(entity ecs.EntityID) {
	var builder uuid.UUID
	if comp, ok := s.CanDeploy().Render().Get(entity); ok && (comp.Render == feature.RuleRenderPreviewRulesPass ||
		comp.Render == feature.RuleRenderCompletePreview) {
		builder = uuid.UUID(comp.Rule.Entity.From(comp.Fields))
		goto render
	}
	return
render:
	selectedEntity, ok := s.UUID().Entity(builder)
	if !ok {
		return
	}

	for _, reachCoords := range s.GameWorld.Deploy().Reach().TilesWithinReach(selectedEntity) {
		ind := s.World().NewEntity()
		s.Hierarchy().SetParent(ind, entity)
		s.Transform().Inherit().Set(ind, transform.NewInherit(transform.Absolute))

		s.Render().Mesh().Set(ind, render.NewMesh(s.Definitions().Assets().SquareMesh))
		s.Render().Texture().Set(ind, render.NewTexture(s.Definitions().Assets().Border))
		s.Groups().Component().Set(ind, groups.EmptyGroups().Enable(definitions.GameGroup))

		s.Tile().Layer().Set(ind, tile.NewLayer(definitions.RangePlaceholderLayer))
		s.Tile().Pos().Set(ind, tile.NewPos(reachCoords.Coords()))
		s.Render().Color().Set(ind, render.NewColor(mgl32.Vec4{0, 0, .5, 1}))
	}
}
