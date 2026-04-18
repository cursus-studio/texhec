package test

import (
	"net"
	"testing"
)

func TestClient(t *testing.T) {
	mutex.Lock()
	defer mutex.Unlock()
	s := NewSetup()

	// host
	listener, err := s.Host()
	if err != nil {
		t.Errorf("\"%v\" is occupied and cannot be tested", s.Addr)
		return
	}

	var connections []net.Conn

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			if err := s.Send(conn, s.Message); err != nil {
				t.Errorf("unexpected error sending message: %v", err)
			}
			connections = append(connections, conn)
		}
	}()

	if connections := len(s.Connection().Component().GetEntities()); connections != 0 {
		t.Errorf("Expected 0 connection not %v", connections)
		return
	}

	// connect
	if err := s.Connection().Connect(s.Addr); err != nil {
		t.Errorf("Unexpected error when hosting: \"%v\"", err)
		return
	}

	if connections := len(s.Connection().Component().GetEntities()); connections != 1 {
		t.Errorf("Expected 1 connection not %v", connections)
		return
	}

	// communication
	s.Sleep()
	connection, _ := s.Connection().Component().Get(s.Connection().Component().GetEntities()[0])
	var message any
	select {
	case message = <-connection.Conn().Messages():
	default:
	}
	if message != s.Message {
		t.Errorf("expected \"%v\" but got \"%v\"", s.Message, message)
		return
	}

	// close
	_ = listener.Close()
	for _, conn := range connections {
		_ = conn.Close()
	}

	s.Sleep()

	if connections := len(s.Connection().Component().GetEntities()); connections != 0 {
		t.Errorf("Expected 0 connection not %v", connections)
		return
	}
}
