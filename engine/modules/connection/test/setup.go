package test

import (
	"encoding/binary"
	"engine"
	typeregistrypkg "engine/modules/typeregistry/pkg"
	enginepkg "engine/pkg"
	"net"
	"sync"
	"time"

	"github.com/ogiusek/ioc/v2"
)

type Message struct {
	Content string
}

var mutex sync.Mutex

type Setup struct {
	engine.EngineWorld `inject:""`

	Message Message
	Network string
	Addr    string
}

func NewSetup() Setup {
	c := ioc.NewContainer(
		enginepkg.Pkg,
		typeregistrypkg.PkgT[Message],
	)
	s := ioc.GetServices[Setup](c)
	s.Message.Content = "example message"
	s.Network = "tcp"
	s.Addr = "localhost:9999"
	return s
}

func (s *Setup) Connect() (net.Conn, error)  { return net.Dial(s.Network, s.Addr) }
func (s *Setup) Host() (net.Listener, error) { return net.Listen(s.Network, s.Addr) }

func (s *Setup) Sleep() {
	time.Sleep(time.Millisecond)
}

func (s *Setup) Send(conn net.Conn, message Message) error {
	bytes, err := s.Codec().Encode(message)
	if err != nil {
		return err
	}

	length := uint32(len(bytes))
	lengthInByes := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthInByes, length)
	if _, err := conn.Write(append(lengthInByes, bytes...)); err != nil {
		return err
	}

	return nil
}
