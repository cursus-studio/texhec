package test

import (
	"net"
	"sync"
	"testing"
)

func TestClient(t *testing.T) {
	mutex.Lock()
	defer mutex.Unlock()
	s := NewSetup()

	// host
	listener, err := s.Host()
	if err != nil {
		t.Fatalf("\"%v\" is occupied and cannot be tested", s.Addr)
	}

	var connMutex sync.Mutex
	var connections []net.Conn
	messageSent := make(chan any, 1)

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			connMutex.Lock()
			connections = append(connections, conn)
			connMutex.Unlock()

			if err := s.Send(conn, s.Message); err != nil {
				t.Errorf("unexpected error sending message: %v", err)
			}

			messageSent <- struct{}{}
		}
	}()

	if count := len(s.Connection().Component().GetEntities()); count != 0 {
		t.Fatalf("Expected 0 connections, got %v", count)
	}

	if err := s.Connection().Connect(s.World().NewEntity(), s.Addr); err != nil {
		t.Fatalf("Unexpected error when connecting: \"%v\"", err)
	}

	if count := len(s.Connection().Component().GetEntities()); count != 1 {
		t.Fatalf("Expected 1 connection, got %v", count)
	}

	s.PollUntil(messageSent)

	connection, _ := s.Connection().Component().Get(s.Connection().Component().GetEntities()[0])
	messages := connection.Conn().Messages()
	if len(messages) != 1 {
		t.Fatalf("expected 1 message but got %v", len(messages))
	} else if messages[0] != s.Message {
		t.Fatalf("expected \"%v\" but got \"%v\"", s.Message, messages[0])
	}

	_ = listener.Close()

	connMutex.Lock()
	for _, conn := range connections {
		_ = conn.Close()
	}
	connMutex.Unlock()

	s.Poll()

	if count := len(s.Connection().Component().GetEntities()); count != 0 {
		t.Fatalf("Expected 0 connections, got %v", count)
	}
}
