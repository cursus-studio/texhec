// this is a generic package which is used to access entities by component (id) in O(1) time
package relation

import (
	"engine/services/ecs"
)

type Service[Key any] interface {
	Get(Key) (ecs.EntityID, bool)
}
