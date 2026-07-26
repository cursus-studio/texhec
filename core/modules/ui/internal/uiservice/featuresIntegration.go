package uiservice

import (
	"engine/modules/ecs"
	"engine/modules/inputs"
	"engine/modules/interactions"
	"engine/modules/text"
)

func (s *service) onAvailableFeaturesMod(entity ecs.EntityID) {
	availableFeatures, _ := s.Interactions().AvailableFeatures().Get(entity)
	if len(availableFeatures.Features) < 2 {
		s.HideMenu()
		return
	}
	type Button struct {
		text  string
		event any
	}
	btns := []Button{}

	for _, feature := range availableFeatures.Features {
		btns = append(btns, Button{feature.Name(), interactions.NewSelectFeatureEvent(feature)})
	}

	for _, p := range s.Ui().ShowMenu() {
		// i want here to display all actions which can be performed by entity
		// currently implement only building
		for _, btn := range btns {
			var btnEntity ecs.EntityID
			if btn.event != nil {
				btnEntity = s.Prototype().Clone(s.Definitions().Hud().Btn)
				s.Inputs().LeftClick().Set(btnEntity, inputs.NewLeftClick(btn.event))
			} else {
				btnEntity = s.Prototype().Clone(s.Definitions().Hud().Text)
			}
			s.Hierarchy().SetParent(btnEntity, p)
			s.Text().Content().Set(btnEntity, text.NewText(btn.text))
		}
	}

}
