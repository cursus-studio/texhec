package internal

import (
	"core/modules/actions"
	"core/modules/tile"
	"engine/modules/ecs"
	"engine/modules/feature"
	"engine/modules/render"
	"engine/modules/transform"
	"engine/modules/uuid"
)

func (s *service) OnEntityHover(event tile.HoverEntityEvent) {
	uuid, ok := s.UUID().UUID().Get(event.Entity)
	if !ok {
		return
	}
	s.Feature().PreviewInteraction(actions.Entity(uuid.ID))
}
func (s *service) OnEntityClick(event tile.ClickEntityEvent) {
	uuid, ok := s.UUID().UUID().Get(event.Entity)
	if !ok {
		return
	}
	s.Feature().SubmitInteraction(actions.Entity(uuid.ID))
}

func (s *service) OnEntityUpsert(entity ecs.EntityID) {
	var entityUUID uuid.UUID
	if comp, ok := s.Own().Render().Get(entity); ok && (comp.Render == feature.RuleRenderPreviewRulesPass ||
		comp.Render == feature.RuleRenderCompletePreview) {
		entityUUID = uuid.UUID(comp.Rule.Entity.From(comp.Fields))
		goto render
	}
	if comp, ok := s.IsEnemy().Render().Get(entity); ok &&
		(comp.Render == feature.RuleRenderPreviewRulesPass ||
			comp.Render == feature.RuleRenderCompletePreview) {
		entityUUID = uuid.UUID(comp.Rule.Entity.From(comp.Fields))
		goto render
	}
	if comp, ok := s.Entity().Render().Get(entity); ok &&
		comp.Render == feature.FieldRenderInteraction &&
		!s.Feature().Interactions().LockedFeature {
		entityUUID = uuid.UUID(comp.Value)
		goto render
	}
	if comp, ok := s.CanBeAttacked().Render().Get(entity); ok && (comp.Render == feature.RuleRenderPreviewRulesPass ||
		comp.Render == feature.RuleRenderCompletePreview) {
		entityUUID = uuid.UUID(comp.Rule.Attacked.From(comp.Fields))
		goto render
	}
	return
render:
	selectedEntity, ok := s.UUID().Entity(entityUUID)
	if !ok {
		return
	}
	owner, ok := s.Player().Owner().Get(selectedEntity)
	if !ok {
		return
	}
	player, ok := s.Player().Color().Get(owner)
	if !ok {
		return
	}

	s.Hierarchy().SetParent(entity, selectedEntity)
	s.Groups().InheritGroups(entity)

	s.Transform().Inherit().Set(entity, transform.NewInherit(transform.RelativePos|transform.RelativeSizeXY))
	s.Transform().Pos().Set(entity, transform.NewPos(0, 0, -1))

	s.Render().Mesh().Set(entity, render.NewMesh(s.Definitions().Assets().SquareMesh))
	s.Render().Color().Set(entity, render.NewColor(player.Color.Vec4(1)))
	s.Render().Texture().Set(entity, render.NewTexture(s.Definitions().Hud().Can))
}
