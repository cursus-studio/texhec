package internal

import (
	"core/modules/actions"
	"core/modules/tile"
	"engine/modules/ecs"
	"engine/modules/feature"
	"engine/modules/render"
	"engine/modules/transform"
	"engine/modules/uuid"

	"github.com/go-gl/mathgl/mgl32"
)

func (s *service) OnEntityHover(event tile.HoverEntityEvent) {
	uuid, ok := s.UUID().Component().Get(event.Entity)
	if !ok {
		return
	}
	s.Feature().PreviewInteraction(actions.Entity(uuid.ID))
}
func (s *service) OnEntityClick(event tile.ClickEntityEvent) {
	uuid, ok := s.UUID().Component().Get(event.Entity)
	if !ok {
		return
	}
	s.Feature().SubmitInteraction(actions.Entity(uuid.ID))
}

func (s *service) OnOwnUpsert(entity ecs.EntityID) {
	comp, ok := s.Own().Render().Get(entity)
	if ok && (comp.Render == feature.RuleRenderPreviewRulesPass ||
		comp.Render == feature.RuleRenderCompletePreview) {
		goto render
	}
	if entityComp, ok := s.Entity().Render().Get(entity); ok &&
		entityComp.Render == feature.FieldRenderInteraction &&
		!s.Feature().Interactions().LockedFeature {
		comp.AppliedTo = entityComp.Value
		goto render
	}
	return
render:
	selectedEntity, ok := s.UUID().Entity(uuid.UUID(comp.AppliedTo))
	if !ok {
		return
	}

	s.Hierarchy().SetParent(entity, selectedEntity)
	s.Groups().InheritGroups(entity)

	s.Transform().Inherit().Set(entity, transform.NewInherit(transform.RelativePos|transform.RelativeSizeXY))
	s.Transform().Pos().Set(entity, transform.NewPos(0, 0, -1))

	s.Render().Mesh().Set(entity, render.NewMesh(s.Definitions().Assets().SquareMesh))
	s.Render().Texture().Set(entity, render.NewTexture(s.Definitions().Hud().Can))
}
func (s *service) OnEnemyUpsert(entity ecs.EntityID) {
	comp, ok := s.IsEnemy().Render().Get(entity)
	if ok && (comp.Render == feature.RuleRenderPreviewRulesPass ||
		comp.Render == feature.RuleRenderCompletePreview) {
		goto render
	}
	if canBeAttacked, ok := s.CanBeAttacked().Render().Get(entity); ok && (canBeAttacked.Render == feature.RuleRenderPreviewRulesPass ||
		canBeAttacked.Render == feature.RuleRenderCompletePreview) {
		comp.Render = canBeAttacked.Render
		comp.AppliedTo = canBeAttacked.AppliedTo
		goto render
	}
	return
render:
	selectedEntity, ok := s.UUID().Entity(uuid.UUID(comp.AppliedTo))
	if !ok {
		return
	}

	s.Hierarchy().SetParent(entity, selectedEntity)
	s.Groups().InheritGroups(entity)

	s.Transform().Inherit().Set(entity, transform.NewInherit(transform.RelativePos|transform.RelativeSizeXY))
	s.Transform().Pos().Set(entity, transform.NewPos(0, 0, -1))

	s.Render().Mesh().Set(entity, render.NewMesh(s.Definitions().Assets().SquareMesh))
	s.Render().Color().Set(entity, render.NewColor(mgl32.Vec4{1, 0, 0, 1}))
	s.Render().Texture().Set(entity, render.NewTexture(s.Definitions().Hud().Can))
}
