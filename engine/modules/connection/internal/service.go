package internal

import (
	"engine"
	"engine/modules/connection"
	"engine/modules/datastructures"
	"engine/modules/ecs"
	"net"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type service struct {
	engine.EngineWorld `inject:""`

	listenersDirtySet ecs.DirtySet
	listeners         datastructures.Set[net.Listener]
	listenersArray    ecs.ComponentArray[connection.ListenerComponent]

	connectionDirtySet ecs.DirtySet
	connections        datastructures.Set[connection.Conn]
	connectionArray    ecs.ComponentArray[connection.ConnectionComponent]
}

func NewService(c ioc.Dic) connection.Service {
	s := ioc.GetServices[*service](c)
	s.listenersDirtySet = ecs.NewDirtySet()
	s.listeners = datastructures.NewSet[net.Listener]()
	s.listenersArray = ecs.GetComponentArray[connection.ListenerComponent](s.World())

	s.connectionDirtySet = ecs.NewDirtySet()
	s.connections = datastructures.NewSet[connection.Conn]()
	s.connectionArray = ecs.GetComponentArray[connection.ConnectionComponent](s.World())

	s.listenersArray.AddDirtySet(s.listenersDirtySet)
	s.listenersArray.OnMod(s.OnListenerMod)

	s.connectionArray.AddDirtySet(s.connectionDirtySet)
	s.connectionArray.OnMod(s.OnConnectionMod)

	return s
}

func (s *service) OnListenerMod(ecs.EntityID) {
	if entities := s.listenersDirtySet.Get(); len(entities) == 0 {
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

func (s *service) OnConnectionMod(ecs.EntityID) {
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
		connection.Close()
	}
}

func (s *service) Register() error {
	events.Listen(s.EventsBuilder(), s.OnFrame)
	return nil
}

func (s *service) Component() ecs.ComponentArray[connection.ConnectionComponent] {
	return s.connectionArray
}
func (s *service) Listener() ecs.ComponentArray[connection.ListenerComponent] {
	return s.listenersArray
}

func (s *service) Host(entity ecs.EntityID, addr string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	s.AddListener(entity, listener)
	return nil
}

func (s *service) Connect(entity ecs.EntityID, addr string) error {
	rawConn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	s.AddConnection(entity, rawConn)
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
}

func (s *service) AddConnection(entity ecs.EntityID, rawConn net.Conn) connection.ConnectionComponent {
	conn := &conn{
		service: s,
		conn:    ConnBuffer{Conn: rawConn},
	}

	comp := connection.NewConnection(conn)
	s.connectionArray.Set(entity, comp)
	return comp
}
