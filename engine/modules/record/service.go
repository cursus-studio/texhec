// This module splits into two separate parts.\
// One to record by `EntityID` which should be used to perform things localy.\
// Second to record by `UUID` which should be used to record things for external machines.
//
// Backwards recording returns state on recording start.\
// Backwards recording can be used to record state before to smoothen changes.
//
// Forwards recording returns state on recording end.\
// Forwards recording can be used to record changes to send them somewhere else to replicate them.
package record

import (
	"engine/modules/datastructures"
	"engine/modules/ecs"
	"engine/modules/uuid"
)

type Service interface {
	// starts opened recording (opened recording is recorded until stopped)
	// applying it on previous state will create current state
	StartRecording(Config) RecordingID
	// starts opened recording (opened recording is recorded until stopped)
	// applying it rewinds state.
	StartBackwardsRecording(Config) RecordingID
	// finishes recording if open (false is returned if recording isn't started)
	Stop(RecordingID) (r Recording, ok bool)
	// returns recording without closing it
	Peek(RecordingID) (r Recording, ok bool)

	Apply(Config, ...Recording)

	// doesn't work for removing entities
	// is it meant to be used for recordings with service lifetime
	UUIDState(Recording) UUIDState
	// applies state and purges all entities which have any of config components
	// it is used to apply whole world state therefor it ensures it is synchronized
	ApplyUUIDState(Config, UUIDState)
}

type RecordingID uint16
type EntityComponents = []any

// recording cannot be encoded
type Recording struct {
	// [componentArrayLayoutID]any component
	// nil for removed entity
	Entities datastructures.SparseArray[ecs.EntityID, EntityComponents]
}

//

type UUIDState struct {
	Keys   []uuid.UUID
	Values []EntityComponents
}
