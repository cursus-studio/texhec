// uses bitmasks and allows us to group entities to do not collide despite shared position or
// to do not be visible for a camera despite being in its view
package groups

import (
	"engine/modules/ecs"
	"engine/modules/hierarchy"
)

type Service interface {
	Component() ecs.ComponentArray[GroupsComponent]
	Inherit() ecs.ComponentArray[hierarchy.InheritComponent[GroupsComponent]]
	InheritGroups(ecs.EntityID)
}
