package internal

import (
	"core/modules/actions"
	"core/modules/tile"
	"engine/modules/ecs"
	"engine/modules/inputs"
	"engine/modules/text"
	"engine/modules/uuid"
	"fmt"
)

func (s *service) OnBlueprintClick(event tile.ClickBlueprintEvent) {
	uuidComp, ok := s.UUID().Component().Get(event.Entity)
	if !ok {
		return
	}
	s.Feature().SubmitInteraction(actions.Blueprint(uuidComp.ID))
}

func (s *service) OnCanBuildBlueprintUpsert(compEntity ecs.EntityID) {
	comp, ok := s.CanBuildBlueprint().Render().Get(compEntity)
	if !ok {
		return
	}
	type Button struct {
		text  string
		event any
	}

	entity, ok := s.UUID().Entity(uuid.UUID(comp.In))
	if !ok {
		s.Logger().Warn(s.CanBuildBlueprint().Error())
		return
	}
	entityBlueprint, ok := s.Tile().Blueprint().Get(entity)
	if !ok {
		s.Logger().Warn(s.CanBuildBlueprint().Error())
		return
	}
	deployed, _ := s.Deploy().Component().Get(entityBlueprint)
	if len(deployed.Deployable) == 0 {
		s.Logger().Warn(s.CanBuildBlueprint().Error())
		return
	}

	btns := make([]Button, 0, len(deployed.Deployable))
	for _, deployableEntity := range deployed.Deployable {
		name, ok := s.Tile().Name().Get(deployableEntity)
		if !ok {
			s.Logger().Warn(s.CanBuildBlueprint().Error())
			continue
		}
		btns = append(btns, Button{
			text:  fmt.Sprintf("%v", name.Name),
			event: tile.NewClickBlueprintEvent(deployableEntity),
		})
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

// import (
// 	"core/modules/actions"
// 	"core/modules/tile"
// 	"engine/modules/ecs"
// 	"engine/modules/inputs"
// 	"engine/modules/interactions"
// 	"engine/modules/text"
// 	"errors"
// 	"fmt"
// )
//
// func (s *service) BlueprintInteraction() interactions.InteractionService[actions.Blueprint] {
// 	return s.BlueprintInteractionService
// }
//
// func (s *service) OnClickBlueprint(event tile.ClickBlueprintEvent) {
// 	uuidComp, ok := s.UUID().Component().Get(event.Entity)
// 	if !ok {
// 		return
// 	}
// 	propertiesEntity := s.World().NewEntity()
// 	s.CoordsCursor().Set(propertiesEntity, actions.NewCoordsCursor(event.Entity, true))
// 	s.Anchor().Set(propertiesEntity, actions.NewAnchor(event.Entity))
//
// 	s.BlueprintInteraction().Save(propertiesEntity, actions.NewBlueprint(uuidComp.ID))
// }
//
// func (s *service) OnBlueprintMissingUpsert(entity ecs.EntityID) {
// 	type Button struct {
// 		text  string
// 		event any
// 	}
// 	btns := []Button{}
// 	canDeploy, ok := s.CanDeploy().Get(entity)
// 	if !ok {
// 		s.Logger().Warn(fmt.Errorf("cannot select blueprint without CanDeploy component"))
// 		return
// 	}
//
// 	deployed, _ := s.Deploy().Component().Get(canDeploy.Entity)
// 	if len(deployed.Deployable) == 0 {
// 		return
// 	}
// 	for _, deployed := range deployed.Deployable {
// 		name, ok := s.Tile().Name().Get(deployed)
// 		if !ok {
// 			s.Logger().Log(errors.New("expected entity to have name component"))
// 			continue
// 		}
// 		btn := Button{fmt.Sprintf("%v", name.Name), tile.NewClickBlueprintEvent(deployed)}
// 		btns = append(btns, btn)
// 	}
//
// 	for _, p := range s.Ui().ShowMenu() {
// 		for _, btn := range btns {
// 			btnEntity := s.Prototype().Clone(s.Definitions().Hud().Btn)
// 			s.Inputs().LeftClick().Set(btnEntity, inputs.NewLeftClick(btn.event))
// 			s.Hierarchy().SetParent(btnEntity, p)
// 			s.Text().Content().Set(btnEntity, text.NewText(btn.text))
// 		}
// 	}
// }
// func (s *service) OnBlueprintMissingRemove(ecs.EntityID) {
// 	s.Ui().HideMenu()
// }
//
// func (s *service) OnBlueprintStateUpsert(entity ecs.EntityID) {
// 	blueprint, ok := s.BlueprintInteraction().StatePreview().Get(entity)
// 	if !ok {
// 		return
// 	}
// 	blueprintEntity, ok := s.UUID().Entity(blueprint.State.UUID)
// 	if !ok {
// 		return
// 	}
//
// 	s.CoordsCursor().Set(entity, actions.NewCoordsCursor(blueprintEntity, true))
// 	if object, ok := s.EntityInteraction().StatePreview().Get(entity); !ok {
// 	} else if objectEntity, ok := s.UUID().Entity(object.State.UUID); ok {
// 		s.Anchor().Set(entity, actions.NewAnchor(objectEntity))
// 	}
// }
