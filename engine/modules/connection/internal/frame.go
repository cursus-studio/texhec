package internal

import (
	"engine/modules/ecs"
	"engine/modules/loop"
	"errors"
	"net"
	"slices"
	"time"
)

func (s *service) OnFrame(loop.FrameEvent) {
	for _, entity := range slices.Clone(s.Connection().Component().GetEntities()) {
		s.PollMessages(entity)
	}

	for _, entity := range slices.Clone(s.Connection().Listener().GetEntities()) {
		s.PollListeners(entity)
	}
}

func (s *service) PollListeners(entity ecs.EntityID) {
	listener, ok := s.Connection().Listener().Get(entity)
	if !ok {
		return
	}
	rawListener := listener.Listener().(*net.TCPListener)
	for {
		_ = rawListener.SetDeadline(time.Now().Add(time.Microsecond))
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
