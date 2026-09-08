package internal

import (
	"engine/modules/loop"
	"slices"
)

func (s *service) OnFrame(loop.FrameEvent) {
	// listen to messages
	for _, entity := range slices.Clone(s.Connection().Component().GetEntities()) {
		s.PollMessages(entity)
	}

	// listen to events
}
