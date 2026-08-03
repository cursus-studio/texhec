package actions

import (
	"engine/modules/ecs"
	"engine/modules/grid"
	"engine/modules/interactions"
	"errors"
)

var (
	ErrRequiresSpeed  error = errors.New("tile:requires speed")
	ErrRequiresDeploy error = errors.New("tile:requires deploy")
)

// components to configure interactions
type CanDeployComponent struct {
	Entity ecs.EntityID
}
type CoordsCursorComponent struct {
	PropertiesEntity ecs.EntityID
	// if true then entity is used as an image else default icon is used
	CustomImage bool
}
type AnchorComponent struct {
	Entity ecs.EntityID
}

func NewCanDeploy(canDeploy ecs.EntityID) CanDeployComponent {
	return CanDeployComponent{canDeploy}
}
func NewCoordsCursor(propertiesEntity ecs.EntityID, customImage bool) CoordsCursorComponent {
	return CoordsCursorComponent{propertiesEntity, customImage}
}
func NewAnchor(entity ecs.EntityID) AnchorComponent {
	return AnchorComponent{entity}
}

//

type CoordsInteraction struct{ Coords grid.Coords }
type EntityInteraction struct{ Entity ecs.EntityID }
type BlueprintInteraction struct{ Entity ecs.EntityID }

func NewCoordsInteraction(coords grid.Coords) CoordsInteraction  { return CoordsInteraction{coords} }
func NewEntityInteraction(entity ecs.EntityID) EntityInteraction { return EntityInteraction{entity} }
func NewBlueprintInteraction(entity ecs.EntityID) BlueprintInteraction {
	return BlueprintInteraction{entity}
}

//

type CoordsStep interactions.Step[CoordsInteraction]

type EntityStep interactions.Step[EntityInteraction]
type FriendlyEntityStep interactions.Step[EntityInteraction]
type FriendlyMobileEntityStep interactions.Step[EntityInteraction]
type FriendlyBuilderEntityStep interactions.Step[EntityInteraction]
type EnemyEntityStep interactions.Step[EntityInteraction]

type BlueprintStep interactions.Step[BlueprintInteraction]

//

type Service interface {
	CanDeploy() ecs.ComponentArray[CanDeployComponent]
	CoordsCursor() ecs.ComponentArray[CoordsCursorComponent]
	Anchor() ecs.ComponentArray[AnchorComponent]

	CoordsInteraction() interactions.InteractionService[CoordsInteraction]
	EntityInteraction() interactions.InteractionService[EntityInteraction]
	BlueprintInteraction() interactions.InteractionService[BlueprintInteraction]
}
