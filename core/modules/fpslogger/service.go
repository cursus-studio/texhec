// logs frames per second
package fpslogger

import "engine/modules/ecs"

type Service interface {
	ecs.SystemRegister
}
