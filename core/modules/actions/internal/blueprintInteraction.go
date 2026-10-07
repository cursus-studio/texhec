package internal

import (
	"core/modules/actions"
	"core/modules/tile"
	"engine/modules/ecs"
	"engine/modules/feature"
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

func (s *service) OnCanDeployBlueprintUpsert(entity ecs.EntityID) {
	comp, ok := s.CanDeployBlueprint().Render().Get(entity)
	if ok && (comp.Render == feature.RuleRenderPreviewRulesFail ||
		comp.Render == feature.RuleRenderPreviewRulesPass) {
		goto render
	}
	return
render:
	type Button struct {
		text  string
		event any
	}

	canEntityBuild, ok := s.UUID().Entity(uuid.UUID(comp.In))
	if !ok {
		s.Logger().Warn(s.CanDeployBlueprint().Error())
		return
	}
	canEntityBuildBlueprint, ok := s.Tile().Blueprint().Get(canEntityBuild)
	if !ok {
		s.Logger().Warn(s.CanDeployBlueprint().Error())
		return
	}
	deployed, _ := s.Deploy().Component().Get(canEntityBuildBlueprint)
	if len(deployed.Deployable) == 0 {
		s.Logger().Warn(s.CanDeployBlueprint().Error())
		return
	}

	btns := make([]Button, 0, len(deployed.Deployable))
	for _, deployableEntity := range deployed.Deployable {
		name, ok := s.Tile().Name().Get(deployableEntity)
		if !ok {
			s.Logger().Warn(s.CanDeployBlueprint().Error())
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
