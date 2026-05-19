// uses bitmasks and allows us to group entities to do not collide despite shared position or
// to do not be visible for a camera despite being in its view
package groups

import (
	"engine/services/ecs"
)

type Service interface {
	Component() ecs.ComponentsArray[GroupsComponent]
	Inherit() ecs.ComponentsArray[InheritGroupsComponent]
}
