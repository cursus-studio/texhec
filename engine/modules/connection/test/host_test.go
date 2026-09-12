package test

import (
	"testing"
)

func TestHost(t *testing.T) {
	mutex.Lock()
	defer mutex.Unlock()
	s := NewSetup()

	if _, err := s.Connect(); err == nil {
		t.Fatalf("\"%v\" is occupied and cannot be tested", s.Addr)
	}

	if listeners := len(s.Connection().Listener().GetEntities()); listeners != 0 {
		t.Fatalf("Expected 0 listener, got %v", listeners)
	}

	// host
	if err := s.Connection().Host(s.World().NewEntity(), s.Addr); err != nil {
		t.Fatalf("Unexpected error when hosting: \"%v\"", err)
	}

	if listeners := len(s.Connection().Listener().GetEntities()); listeners != 1 {
		t.Fatalf("Expected 1 listener, got %v", listeners)
	}

	// connect
	conn, err := s.Connect()
	if err != nil {
		t.Fatalf("Unexpected error when connecting: \"%v\"", err)
	}

	// Signal connection is established before polling
	connectedSignal := make(chan any, 1)
	connectedSignal <- struct{}{}
	s.PollUntil(connectedSignal)

	if connections := len(s.Connection().Component().GetEntities()); connections != 1 {
		t.Fatalf("Expected 1 connection, got %v", connections)
	}

	// close client connection
	if err := conn.Close(); err != nil {
		t.Fatalf("Unexpected error closing connection: \"%v\"", err)
	}

	// Signal disconnect event before polling cleanup
	closedSignal := make(chan any, 1)
	closedSignal <- struct{}{}
	s.PollUntil(closedSignal)

	if connections := len(s.Connection().Component().GetEntities()); connections != 0 {
		t.Fatalf("Expected 0 connection, got %v", connections)
	}

	listener, _ := s.Connection().Listener().Get(s.Connection().Listener().GetEntities()[0])
	_ = listener.Listener().Close()
}
