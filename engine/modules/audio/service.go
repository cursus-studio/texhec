// this module is reponsible for integrating audio via events
package audio

import (
	"engine/modules/ecs"
)

type Channel int
type Volume float32 // volume is normalized

type Service interface {
	ecs.SystemRegister
	PlayerService
	VolumeService
}

type PlayerService interface {
	Stop(Channel) error
	Play(Channel, ecs.EntityID) error
	Queue(Channel, ecs.EntityID) error
	QueueEndless(Channel, ecs.EntityID) error
}

type VolumeService interface {
	SetMasterVolume(Volume) error
	SetChannelVolume(Channel, Volume) error
}
