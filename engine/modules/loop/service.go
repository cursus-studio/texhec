// defines game loop
package loop

import (
	"time"
)

// stops game loop
// can listen to it to clean up
type StopEvent struct{}

// changes game loop configuration
type ConfigureEvent struct {
	FPS,
	TPS int
}

func NewStopEvent() StopEvent { return StopEvent{} }
func NewConfigureEvent(fps, tps int) ConfigureEvent {
	return ConfigureEvent{fps, tps}
}

//

// tick has fixed delta to ensure determinism
// tick is triggered before frame as many times as many ticks passed between frames
type TickEvent struct{ Delta time.Duration }
type FrameEvent struct{ Delta time.Duration }

//

type Stats interface {
	FrameBudget() time.Duration
	FrameBudgetLeft() time.Duration
}

//

type EmitOnTickComponent struct {
	Event any
}

type EmitOnTickEvent struct {
	Event any
}

func NewEmitOnTickEvent(event any) EmitOnTickEvent {
	return EmitOnTickEvent{event}
}

//

type Service interface {
	// Starts the game loop if it isn't started.
	// Waits until game loop stops.
	Run(initialConfiguration ConfigureEvent)
	Stop()                    // emits stop event
	Configure(ConfigureEvent) // emits confugure event

	// also re-ticks
	SyncToUnixNano(int64)
	LastTickUnixNano() int64

	Stats() Stats

	EmitOnTick(event any)
}
