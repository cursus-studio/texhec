// defines connection and stores it in component
package connection

import (
	"engine/modules/ecs"
	"net"
)

// types

// singular connection interface
type Conn interface {
	// send has block behavior
	Send(message any) error

	// returns messages from last call
	Messages() []any
	Close() error
}

// components

type ListenerComponent struct {
	listener net.Listener
}

func NewListener(listener net.Listener) ListenerComponent {
	return ListenerComponent{listener}
}

func (comp *ListenerComponent) Listener() net.Listener {
	return comp.listener
}

//

type ConnectionComponent struct {
	conn Conn
}

func NewConnection(conn Conn) ConnectionComponent {
	return ConnectionComponent{conn}
}

func (comp *ConnectionComponent) Conn() Conn {
	return comp.conn
}

type Service interface {
	ecs.SystemRegister
	Component() ecs.ComponentArray[ConnectionComponent]
	Listener() ecs.ComponentArray[ListenerComponent]

	Host(entity ecs.EntityID, addr string) error
	Connect(entity ecs.EntityID, addr string) error

	TransferConnection(fromEntity, toEntity ecs.EntityID) error
}
