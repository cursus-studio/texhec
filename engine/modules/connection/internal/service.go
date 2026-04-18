package internal

import (
	"encoding/binary"
	"engine"
	"engine/modules/connection"
	"engine/services/datastructures"
	"engine/services/ecs"
	"io"
	"net"

	"github.com/ogiusek/ioc/v2"
)

type service struct {
	engine.EngineWorld `inject:""`

	listenersDirtySet ecs.DirtySet
	listeners         datastructures.Set[net.Listener]
	listenersArray    ecs.ComponentsArray[connection.ListenerComponent]

	connectionDirtySet ecs.DirtySet
	connections        datastructures.Set[connection.Conn]
	connectionArray    ecs.ComponentsArray[connection.ConnectionComponent]
}

func NewService(c ioc.Dic) connection.Service {
	s := ioc.GetServices[*service](c)
	s.listenersDirtySet = ecs.NewDirtySet()
	s.listeners = datastructures.NewSet[net.Listener]()
	s.listenersArray = ecs.GetComponentsArray[connection.ListenerComponent](s.World())

	s.connectionDirtySet = ecs.NewDirtySet()
	s.connections = datastructures.NewSet[connection.Conn]()
	s.connectionArray = ecs.GetComponentsArray[connection.ConnectionComponent](s.World())

	s.listenersArray.AddDirtySet(s.listenersDirtySet)
	s.listenersArray.OnUpsert(s.BeforeListenerGet)

	s.connectionArray.AddDirtySet(s.connectionDirtySet)
	s.connectionArray.OnUpsert(s.BeforeConnectionGet)

	return s
}

func (s *service) BeforeListenerGet(ecs.EntityID) {
	if entities := s.connectionDirtySet.Get(); len(entities) == 0 {
		return
	}
	present := datastructures.NewSet[net.Listener]()
	for _, entity := range s.listenersArray.GetEntities() {
		comp, ok := s.listenersArray.Get(entity)
		if !ok {
			continue
		}
		conn := comp.Listener()
		if conn == nil {
			continue
		}
		present.Add(conn)
	}

	for _, listener := range s.listeners.Get() {
		_, ok := present.GetIndex(listener)
		if ok {
			continue
		}
		s.listeners.RemoveElements(listener)
		_ = listener.Close()
	}
}

func (s *service) BeforeConnectionGet(ecs.EntityID) {
	if entities := s.connectionDirtySet.Get(); len(entities) == 0 {
		return
	}
	present := datastructures.NewSet[connection.Conn]()
	for _, entity := range s.connectionArray.GetEntities() {
		comp, ok := s.connectionArray.Get(entity)
		if !ok {
			continue
		}
		conn := comp.Conn()
		if conn == nil {
			continue
		}
		present.Add(conn)
	}

	for _, connection := range s.connections.Get() {
		_, ok := present.GetIndex(connection)
		if ok {
			continue
		}
		s.connections.RemoveElements(connection)
		_ = connection.Close()
	}
}

func (s *service) Component() ecs.ComponentsArray[connection.ConnectionComponent] {
	return s.connectionArray
}
func (s *service) Listener() ecs.ComponentsArray[connection.ListenerComponent] {
	return s.listenersArray
}

func (s *service) Host(addr string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	s.AddListener(s.World().NewEntity(), listener)
	return nil
}

func (s *service) Connect(addr string) error {
	rawConn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	s.AddConnection(s.World().NewEntity(), rawConn)
	return nil
}

func (s *service) TransferConnection(entityFrom, entityTo ecs.EntityID) error {
	comp, ok := s.connectionArray.Get(entityFrom)
	if !ok {
		return nil
	}
	s.connectionArray.Remove(entityFrom)
	s.connectionArray.Set(entityTo, comp)
	return nil
}

func (s *service) AddListener(entity ecs.EntityID, rawListener net.Listener) {
	s.listeners.Add(rawListener)
	comp := connection.NewListener(rawListener)
	s.listenersArray.Set(entity, comp)

	go func() {
		for {
			rawConn, err := rawListener.Accept()
			if err != nil {
				break
			}
			clientEntity := s.World().NewEntity()
			s.Hierarchy().SetParent(clientEntity, entity)
			s.AddConnection(clientEntity, rawConn)
		}
		if comp, ok := s.listenersArray.Get(entity); ok && comp.Listener() == rawListener {
			s.World().RemoveEntity(entity)
		}

		_ = rawListener.Close()
	}()
}

func (s *service) AddConnection(entity ecs.EntityID, rawConn net.Conn) {
	conn := &conn{
		service:  s,
		conn:     rawConn,
		messages: make(chan any),
	}
	comp := connection.NewConnection(conn)
	s.connectionArray.Set(entity, comp)
	go func() {
		for {
			messageLengthInBytes := make([]byte, 4)
			if _, err := io.ReadFull(rawConn, messageLengthInBytes); err != nil {
				break
			}
			messageLength := binary.BigEndian.Uint32(messageLengthInBytes)
			messageBytes := make([]byte, messageLength)
			if _, err := io.ReadFull(rawConn, messageBytes); err != nil {
				break
			}

			message, err := s.Codec().Decode(messageBytes)
			if err != nil {
				s.Logger().Warn(err)
				continue
			}
			// f.logger.Info(fmt.Sprintf("received '***' type '%v'", reflect.TypeOf(message).String()))

			conn.messages <- message
		}
		if connComp, ok := s.connectionArray.Get(entity); ok && connComp.Conn() == conn {
			s.World().RemoveEntity(entity)
		}
		close(conn.messages)
		_ = rawConn.Close()
	}()
}
