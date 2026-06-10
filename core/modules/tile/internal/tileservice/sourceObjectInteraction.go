package tileservice

import (
	"core/modules/tile"
	"engine/modules/inputs"
	"engine/modules/interactions"
	"engine/modules/text"
	"engine/services/ecs"
	"errors"
	"fmt"
)

func (s *service) SourceObjectInteraction() interactions.InteractionService[tile.SourceObjectInteraction] {
	return s.SourceObjectInteractionService
}

// create a list of blueprints and how it in ui
func (s *service) OnMissingSourceObjectInteractionUpsert(ecs.EntityID) {
	type Button struct {
		text  string
		event any
	}
	btns := []Button{}
	featureEntity := s.Interactions().FeatureEntity()
	object, ok := s.ObjectInteraction().Interaction().Get(featureEntity)
	if !ok {
		s.Logger().Warn(fmt.Errorf("cannot select source object when object isn't selected"))
		return
	}

	link, ok := s.Metadata().Link().Get(object.State.Entity)
	if !ok {
		return
	}

	// TODO interactions deploy
	deployed, _ := s.Deploy().Component().Get(link.Entity)
	if len(deployed.Deployable) == 0 {
		return
	}
	for _, deployed := range deployed.Deployable {
		name, ok := s.Metadata().Name().Get(deployed)
		if !ok {
			s.Logger().Log(errors.New("expected entity to have name component"))
			continue
		}
		event := interactions.NewFinishMeasurementEvent(tile.NewSourceObjectInteraction(deployed))
		btn := Button{fmt.Sprintf("%v", name.Name), event}
		btns = append(btns, btn)
	}

	for _, p := range s.Ui().ShowMenu() {
		// i want here to display all actions which can be performed by entity
		// currently implement only building
		for _, btn := range btns {
			btnEntity := s.Prototype().Clone(s.Definitions().Hud().Btn)
			s.Inputs().LeftClick().Set(btnEntity, inputs.NewLeftClick(btn.event))
			s.Hierarchy().SetParent(btnEntity, p)
			s.Text().Content().Set(btnEntity, text.NewText(btn.text))
		}
	}
}
func (s *service) OnMissingSourceObjectInteractionRemove(ecs.EntityID) {
	s.Ui().HideMenu()
}
