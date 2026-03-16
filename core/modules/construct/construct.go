package construct

import "engine/services/ecs"

type System ecs.SystemRegister

//

type ConstructComponent struct {
	Construct ecs.EntityID
}

func NewConstruct(construct ecs.EntityID) ConstructComponent { return ConstructComponent{construct} }

//

type ClickEvent struct {
	Construct ecs.EntityID
}

func NewClickEvent(construct ecs.EntityID) ClickEvent {
	return ClickEvent{construct}
}

// type BlueprintComponent struct {
// 	Construct string
// 	// Size int
//
// 	// complexity 1:
// 	// texture
// 	// click event
// 	// size (1x1 or 2x2 for example)
//
// 	// complexity 2:
// 	// healh
// 	// profits
// 	// other features like defense
// }
//
// func NewBlueprint(construct string) BlueprintComponent {
// 	return BlueprintComponent{construct}
// }

//

type Service interface {
	// adds mesh, texture, mouse click event, adds to grid
	Construct() ecs.ComponentsArray[ConstructComponent]
}
