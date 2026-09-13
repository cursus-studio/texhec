package internal

import (
	"engine/modules/ecs"
	"engine/modules/loop"
	"errors"
	"net"
	"slices"
	"time"
)

func getDeadline() time.Time {
	// 3 microseconds are tuned to pass tests with command `go -C engine/modules/connection test -count=1000 ./...`
	// while being lowest possible.
	// if this will be to big than game will slow down and if to small than messages will arrive across to many frames
	//
	// here loop isn't used because this would make fetching to hard to achive and there are to few gains in this scenario.
	return time.Now().Add(time.Microsecond * 3)
}

func (s *service) OnFrame(loop.FrameEvent) {
	for _, entity := range slices.Clone(s.Connection().Listener().GetEntities()) {
		s.PollListeners(entity)
	}
	for _, entity := range slices.Clone(s.Connection().Component().GetEntities()) {
		s.PollMessages(entity)
	}
}

func (s *service) PollListeners(entity ecs.EntityID) {
	listener, ok := s.Connection().Listener().Get(entity)
	if !ok {
		return
	}
	rawListener := listener.Listener().(*net.TCPListener)
	for {
		_ = rawListener.SetDeadline(getDeadline())
		rawConn, err := rawListener.Accept()
		if err == nil {
			clientEntity := s.World().NewEntity()
			s.Hierarchy().SetParent(clientEntity, entity)
			s.AddConnection(clientEntity, rawConn)
			continue
		}
		var netErr net.Error
		if errors.As(err, &netErr) && netErr.Timeout() {
			return
		}
		break
	}

	if comp, ok := s.listenersArray.Get(entity); ok && comp.Listener() == rawListener {
		s.World().RemoveEntity(entity)
	}
	_ = rawListener.Close()
}
