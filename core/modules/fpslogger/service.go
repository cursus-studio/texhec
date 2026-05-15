package fpslogger

import "engine/services/ecs"

type Service interface {
	ecs.SystemRegister
}
