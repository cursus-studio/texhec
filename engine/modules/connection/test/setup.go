package test

import (
	"encoding/binary"
	"engine/modules/connection"
	connectionpkg "engine/modules/connection/pkg"
	hierarchypkg "engine/modules/hierarchy/pkg"
	"engine/services/clock"
	"engine/services/codec"
	"engine/services/ecs"
	"engine/services/logger"
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
	World      ecs.World          `inject:"1"`
	Connection connection.Service `inject:"1"`
	Codec      codec.Codec        `inject:"1"`

	Message Message
	Network string
	Addr    string
}

func NewSetup() Setup {
	c := ioc.NewContainer(
		logger.Pkg(logger.NewConfig(true, func(c ioc.Dic, message string) { print(message) })),
		clock.Pkg(time.RFC3339Nano),
		ecs.Pkg,
		codec.Pkg,
		hierarchypkg.Pkg,
		connectionpkg.Pkg,
		func(b ioc.Builder) {
			ioc.Wrap(b, func(c ioc.Dic, builder codec.Builder) {
				builder.Register(
					Message{},
				)
			})
		},
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
	bytes, err := s.Codec.Encode(message)
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
