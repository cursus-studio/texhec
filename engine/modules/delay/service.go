// package responsible for delaying events
package delay

import (
	"engine/modules/ecs"
	"time"
)

type applyDelayWrapper struct{ Event any }

func (w applyDelayWrapper) ApplyDelay(time.Duration) any { return w.Event }

//

type ApplyDelay interface {
	ApplyDelay(time.Duration) any
}

// emits event after duration on frame
type DelayedEvent struct {
	Event    ApplyDelay
	Duration time.Duration
}

func NewDelayedEvent(
	event any,
	duration time.Duration,
) DelayedEvent {
	if event, ok := event.(ApplyDelay); ok {
		return DelayedEvent{
			Event:    event,
			Duration: duration,
		}
	}
	return DelayedEvent{
		Event:    applyDelayWrapper{event},
		Duration: duration,
	}
}

type Service interface {
	ecs.SystemRegister
	Delay(DelayedEvent)
}
