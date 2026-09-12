package test

import (
	"encoding/binary"
	"engine"
	"engine/modules/ecs"
	"engine/modules/loop"
	typeregistrypkg "engine/modules/typeregistry/pkg"
	enginepkg "engine/pkg"
	"fmt"
	"math"
	"net"
	"sync"
	"time"

	"github.com/ogiusek/events"
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
	_ = ecs.RegisterSystems(
		s.Connection(),
	)
	s.Message.Content = "example message"
	s.Network = "tcp"

	ln, err := net.Listen(s.Network, "127.0.0.1:0")
	if err != nil {
		panic(err)
	}
	s.Addr = ln.Addr().String()
	_ = ln.Close()

	return s
}

func (s *Setup) Connect() (net.Conn, error)  { return net.Dial(s.Network, s.Addr) }
func (s *Setup) Host() (net.Listener, error) { return net.Listen(s.Network, s.Addr) }

func (s *Setup) Poll() {
	time.Sleep(time.Millisecond * 1) // wait for message to be delivered
	events.Emit(s.Events(), loop.FrameEvent{})
}

func (s *Setup) PollUntil(await chan any) {
	<-await
	events.Emit(s.Events(), loop.FrameEvent{})
}

func (s *Setup) Send(conn net.Conn, message Message) error {
	bytes, err := s.Codec().Encode(message)
	if err != nil {
		return err
	}

	bytesLength := len(bytes)
	if bytesLength < 0 || bytesLength > math.MaxUint32 {
		return fmt.Errorf("message length exceeded maximal %v length", bytesLength)
	}

	length := uint32(bytesLength)
	lengthInByes := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthInByes, length)
	if _, err := conn.Write(append(lengthInByes, bytes...)); err != nil {
		return err
	}

	return nil
}
