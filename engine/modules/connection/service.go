// defines connection and stores it in component
package connection

import (
	"engine/modules/ecs"
	"engine/modules/uuid"
	"net"
)

// types

type MsgCtx struct {
	ConnEntity ecs.EntityID
}

func NewMsgCtx(connEntity ecs.EntityID) MsgCtx {
	return MsgCtx{connEntity}
}

func (ptr *MsgCtx) SetCtx(tgtMsg MsgCtx) { *ptr = tgtMsg }

type MsgCtxSetter interface {
	SetCtx(tgtMsg MsgCtx)
}

//

// singular connection interface
type Conn interface {
	Close()
	// messages received are emited as event

	// send has block behavior
	Send(message any) error
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

//

type SetConnUUIDDTO struct {
	MsgCtx
	UUID uuid.UUID
}

func NewSetConnUUIDDTO(uuid uuid.UUID) SetConnUUIDDTO {
	return SetConnUUIDDTO{UUID: uuid}
}

//

type Service interface {
	ecs.SystemRegister
	Component() ecs.ComponentArray[ConnectionComponent]
	Listener() ecs.ComponentArray[ListenerComponent]

	Host(entity ecs.EntityID, addr string) error
	Connect(entity ecs.EntityID, addr string) error

	TransferConnection(fromEntity, toEntity ecs.EntityID) error
}
