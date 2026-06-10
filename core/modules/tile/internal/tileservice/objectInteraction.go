package tileservice

import (
	"core/modules/tile"
	"engine/modules/interactions"
	"engine/services/ecs"
)

func (s *service) ObjectInteraction() interactions.InteractionService[tile.ObjectInteraction] {
	return s.ObjectInteractionService
}

// create a list of object and show it (TODO later. Currently just on object click select it)
func (s *service) OnMissingObjectInteractionUpsert(entity ecs.EntityID) {
}
func (s *service) OnMissingObjectInteractionRemove(entity ecs.EntityID) {
}

func (s *service) OnClickEntity(event tile.ClickEntityEvent) {
	featureEntity := s.Interactions().FeatureEntity()
	comp := interactions.NewInteraction(tile.NewObjectInteraction(event.Entity))
	s.ObjectInteractionService.Interaction().Set(featureEntity, comp)
	s.ObjectInteractionService.MissingInteraction().Remove(featureEntity)
}
