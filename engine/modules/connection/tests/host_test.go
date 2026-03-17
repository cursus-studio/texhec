package test

import (
	"testing"
)

func TestHost(t *testing.T) {
	mutex.Lock()
	defer mutex.Unlock()
	s := NewSetup()

	if _, err := s.Connect(); err == nil {
		t.Errorf("\"%v\" is occupied and cannot be tested", s.Addr)
		return
	}

	if listeners := len(s.Connection.Listener().GetEntities()); listeners != 0 {
		t.Errorf("Expected 0 listener not %v", listeners)
		return
	}

	// host
	if err := s.Connection.Host(s.Addr); err != nil {
		t.Errorf("Unexpected error when hosting: \"%v\"", err)
		return
	}

	if listeners := len(s.Connection.Listener().GetEntities()); listeners != 1 {
		t.Errorf("Expected 1 listener not %v", listeners)
		return
	}

	// connect
	conn, err := s.Connect()
	if err != nil {
		t.Errorf("Unexpected error when connecting: \"%v\"", err)
		return
	}

	s.Sleep()

	if connections := len(s.Connection.Component().GetEntities()); connections != 1 {
		t.Errorf("Expected 1 connection not %v", connections)
		return
	}

	// can add here communication tests

	// close
	if err := conn.Close(); err != nil {
		t.Errorf("Unexpected error \"%v\"", err)
		return
	}

	s.Sleep()

	if connections := len(s.Connection.Component().GetEntities()); connections != 0 {
		t.Errorf("Expected 0 connection not %v", connections)
		return
	}

	listener, _ := s.Connection.Listener().Get(s.Connection.Listener().GetEntities()[0])
	_ = listener.Listener().Close()
}
