package frames

import (
	"engine/services/clock"
	"errors"
	"time"

	"github.com/ogiusek/events"
)

// emmited to service
type StopEvent struct{}
type StartEvent struct{}

// emmited by service
type CleanUpEvent struct{}

var (
	ErrAlreadyRunning error = errors.New("already running")
)

type Frames interface {
	FrameBudget() time.Duration
	FrameBudgetLeft() time.Duration

	Run() error
	Stop()
}

type frames struct {
	Running bool
	TPS,
	FPS int

	TickDuration,
	FrameDuration time.Duration

	LastFrameTime time.Time

	TickProgress time.Duration
	Events       events.Events
	Clock        clock.Clock
}

func (frames *frames) FrameBudget() time.Duration {
	return frames.FrameDuration
}
func (frames *frames) FrameBudgetLeft() time.Duration {
	return max(0, frames.LastFrameTime.Add(frames.FrameDuration).Sub(frames.Clock.Now()))
}

func (frames *frames) StartLoop() {
	ticker := time.NewTicker(frames.FrameDuration)
	defer ticker.Stop()

	frames.LastFrameTime = frames.Clock.Now()

	for frames.Running {
		<-ticker.C
		currentTime := frames.Clock.Now()

		delta := currentTime.Sub(frames.LastFrameTime)
		frames.LastFrameTime = currentTime

		frameEvent := NewFrameEvent(delta)
		frames.TickProgress += delta
		for frames.TickProgress > frames.TickDuration {
			frames.TickProgress -= frames.TickDuration
			events.Emit(frames.Events, TickEvent{frames.TickDuration})
		}
		events.Emit(frames.Events, frameEvent)
	}
	events.Emit(frames.Events, CleanUpEvent{})
}

func (frames *frames) Run() error {
	if frames.Running {
		return ErrAlreadyRunning
	}

	frames.Running = true
	frames.StartLoop()

	return nil
}

func (frames *frames) Stop() {
	frames.Running = false
}
