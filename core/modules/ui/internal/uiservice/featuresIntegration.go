package uiservice

import (
	"engine/modules/ecs"
	"engine/modules/feature"
	"engine/modules/inputs"
	"engine/modules/text"
)

func (s *service) onAvailableFeaturesMod(entity ecs.EntityID) {
	availableFeatures := s.Feature().Interactions()
	if len(availableFeatures.Features) == 0 {
		s.HideMenu()
		return
	}
	type Button struct {
		text  string
		event any
	}
	btns := []Button{}

	for _, feat := range availableFeatures.Features {
		btns = append(btns, Button{feat.Name(), feature.NewLockFeatureEvent(feat)})
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
