package tileservice

import (
	"core/modules/tile"
	"engine/modules/ecs"
	"engine/modules/inputs"
	"engine/modules/interactions"
	"engine/modules/text"
	"errors"
	"fmt"
)

func (s *service) BlueprintInteraction() interactions.InteractionService[tile.BlueprintInteraction] {
	return s.BlueprintInteractionService
}

func (s *service) OnClickBlueprint(event tile.ClickBlueprintEvent) {
	propertiesEntity := s.World().NewEntity()
	s.CoordsCursor().Set(propertiesEntity, tile.NewCoordsCursor(event.Entity, true))
	s.CoordsAnchor().Set(propertiesEntity, tile.NewCoordsAnchor(event.Entity))

	s.BlueprintInteraction().Save(propertiesEntity, tile.NewBlueprintInteraction(event.Entity))
}

func (s *service) OnBlueprintMissingUpsert(entity ecs.EntityID) {
	type Button struct {
		text  string
		event any
	}
	btns := []Button{}
	canDeploy, ok := s.CanDeploy().Get(entity)
	if !ok {
		s.Logger().Warn(fmt.Errorf("cannot select blueprint without CanDeploy component"))
		return
	}

	link, ok := s.Metadata().Link().Get(canDeploy.Entity)
	if !ok {
		return
	}

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
		btn := Button{fmt.Sprintf("%v", name.Name), tile.NewClickBlueprintEvent(deployed)}
		btns = append(btns, btn)
	}

	for _, p := range s.Ui().ShowMenu() {
		for _, btn := range btns {
			btnEntity := s.Prototype().Clone(s.Definitions().Hud().Btn)
			s.Inputs().LeftClick().Set(btnEntity, inputs.NewLeftClick(btn.event))
			s.Hierarchy().SetParent(btnEntity, p)
			s.Text().Content().Set(btnEntity, text.NewText(btn.text))
		}
	}
}
func (s *service) OnBlueprintMissingRemove(ecs.EntityID) {
	s.Ui().HideMenu()
}

func (s *service) OnBlueprintStateUpsert(entity ecs.EntityID) {
	blueprint, ok := s.BlueprintInteraction().StatePreview().Get(entity)
	if !ok {
		return
	}

	s.CoordsCursor().Set(entity, tile.NewCoordsCursor(blueprint.State.Entity, true))
	if object, ok := s.ObjectInteraction().StatePreview().Get(entity); ok {
		s.CoordsAnchor().Set(entity, tile.NewCoordsAnchor(object.State.Entity))
	}
}
