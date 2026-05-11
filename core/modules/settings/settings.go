package settings

import (
	"engine/services/ecs"
)

type Service interface {
	ecs.SystemRegister
}

type EnterSettingsEvent struct{}

type EnterSettingsForParentEvent struct {
	Parent ecs.EntityID
}
