package tileservice

import (
	"core/modules/definitions"
	"core/modules/deploy"
	"core/modules/pathfind"
	"core/modules/tile"
	"engine/modules/ecs"
	"engine/modules/inputs"
	"engine/modules/interactions"
	"engine/modules/render"
	"engine/modules/text"
	"engine/modules/transform"
	"errors"
	"fmt"
)

func (s *service) ObjectInteraction() interactions.InteractionService[tile.ObjectInteraction] {
	return s.ObjectInteractionService
}

// create a list of object and show it (TODO later. Currently just on object click select it)
func (s *service) OnMissingObjectInteractionUpsert(entity ecs.EntityID) {
}
func (s *service) OnMissingObjectInteractionRemove(entity ecs.EntityID) {
}

func (s *service) OnClickEntitySelect(event tile.ClickEntityEvent) {
	featureEntity := s.Interactions().FeatureEntity()
	comp := interactions.NewInteraction(tile.NewObjectInteraction(event.Entity))
	s.ObjectInteractionService.Interaction().Set(featureEntity, comp)
	s.ObjectInteractionService.MissingInteraction().Remove(featureEntity)

	link, ok := s.Metadata().Link().Get(event.Entity)
	if !ok {
		s.Logger().Warn(fmt.Errorf("cannot click entity which doesn't have original entity"))
		return
	}

	s.CoordsCursor().Set(featureEntity, tile.NewCoordsCursor(link.Entity, false))
}

func (s *service) OnClickEntityRenderFeatures(e tile.ClickEntityEvent) {
	link, ok := s.Metadata().Link().Get(e.Entity)
	if !ok {
		s.Logger().Log(errors.New("expected entity to have link component"))
		return
	}
	name, ok := s.Metadata().Name().Get(link.Entity)
	if !ok {
		s.Logger().Log(errors.New("expected link to have name component"))
		return
	}
	owner, ok := s.Player().Owner().Get(e.Entity)
	if !ok {
		s.Logger().Log(errors.New("object without owner cannot build"))
		return
	}
	playerName, ok := s.Metadata().Name().Get(owner.Owner)
	if !ok {
		s.Logger().Log(errors.New("expected player to have player component"))
		return
	}

	type Button struct {
		text  string
		event any
	}
	btns := []Button{
		{fmt.Sprintf("%v's %v", playerName.Name, name.Name), nil},
	}

	if deployed, _ := s.Deploy().Component().Get(link.Entity); len(deployed.Deployable) != 0 {
		btns = append(btns, Button{"Deploy", deploy.NewDeployFeature()})
	}
	if _, ok := s.Pathfind().Speed().Get(link.Entity); ok {
		btns = append(btns, Button{"Move", pathfind.NewFindPathFeature()})
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

func (s *service) OnObjectInteractionUpsert(ecs.EntityID) {
	worldEntity, ok := s.Seed().WorldSeed()
	if !ok {
		return
	}

	placeholders := s.ObjectInteractionService.InteractionGUI().GetEntities()
	if len(placeholders) > 1 {
		for _, entity := range s.CoordsInteractionService.InteractionGUI().GetEntities()[1:] {
			s.World().RemoveEntity(entity)
		}
	}
	featureEntity := s.Interactions().FeatureEntity()
	object, ok := s.ObjectInteractionService.Interaction().Get(featureEntity)
	if !ok {
		return
	}

	var placeholderEntity ecs.EntityID
	if len(placeholders) == 0 {
		placeholderEntity = s.World().NewEntity()
	} else {
		placeholderEntity = placeholders[0]
	}
	s.Hierarchy().SetParent(placeholderEntity, worldEntity)

	if pos, ok := s.Tile().Pos().Get(object.State.Entity); ok {
		s.Tile().Pos().Set(placeholderEntity, pos)
	}
	s.Tile().Rot().Set(placeholderEntity, tile.NewRot(0))
	if size, ok := s.Tile().Size().Get(object.State.Entity); ok {
		s.Tile().Size().Set(placeholderEntity, size)
	}
	s.Transform().Parent().Set(placeholderEntity, transform.NewParent(transform.Absolute))
	s.Tile().Layer().Set(placeholderEntity, tile.NewLayer(definitions.ObjectSelectionPlaceholderLayer))

	s.Render().Mesh().Set(placeholderEntity, render.NewMesh(s.Definitions().Assets().SquareMesh))
	s.Render().Texture().Set(placeholderEntity, render.NewTexture(s.Definitions().Hud().Can))
	s.Groups().InheritGroups(placeholderEntity)

	s.Obstruction().Deployed().Remove(placeholderEntity)
	s.Collider().Component().Remove(placeholderEntity)

	s.ObjectInteractionService.InteractionGUI().Set(placeholderEntity, interactions.InteractionGUIComponent[tile.ObjectInteraction]{})
}
