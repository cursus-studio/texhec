// renders settings GUI
package settings

import (
	"engine/modules/ecs"
)

type Service interface {
	ecs.SystemRegister
}

type EnterSettingsEvent struct{}

type EnterSettingsForParentEvent struct {
	Parent ecs.EntityID
}
