package test

import (
	"testing"
)

func TestClient(t *testing.T) {
	s := NewSetup()

	listener, err := s.Host()
	if err != nil {
		t.Fatalf("\"%v\" is occupied and cannot be tested", s.Addr)
	}

	if connections := len(s.Connection().Component().GetEntities()); connections != 0 {
		t.Fatalf("Expected 0 connection not %v", connections)
		return
	}

	if err := s.Connection().Connect(s.World().NewEntity(), s.Addr); err != nil {
		t.Fatalf("Unexpected error when hosting: \"%v\"", err)
		return
	}

	// ensure connection request is sent
	s.Poll()

	connFromServer, err := listener.Accept()
	if err != nil {
		return
	}

	// ensure connection request is accepted
	s.Poll()

	if connections := len(s.Connection().Component().GetEntities()); connections != 1 {
		t.Fatalf("Expected 1 connection not %v", connections)
	}

	if err := s.Send(connFromServer, s.Message); err != nil {
		t.Fatalf("unexpected error sending message: %v", err)
	}

	// ensure message is sent
	s.Poll()

	connection, _ := s.Connection().Component().Get(s.Connection().Component().GetEntities()[0])
	messages := connection.Conn().Messages()
	if len(messages) != 1 {
		t.Fatalf("expected \"%v\" but received no messgaes", s.Message)
	} else if messages[0] != s.Message {
		t.Fatalf("expected \"%v\" but got \"%v\"", s.Message, messages[0])
	}

	_ = listener.Close()
	_ = connFromServer.Close()

	// ensure connection is removed
	s.Poll()

	if connections := len(s.Connection().Component().GetEntities()); connections != 0 {
		t.Fatalf("Expected 0 connection not %v", connections)
	}
}
