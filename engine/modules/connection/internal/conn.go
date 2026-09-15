package internal

import (
	"encoding/binary"
	"engine/modules/ecs"
	"errors"
	"fmt"
	"math"
	"net"
	"os"
)

type ConnBuffer struct {
	net.Conn
	buffer []byte
}

func (conn *ConnBuffer) Buffer(buffer []byte) {
	conn.buffer = append(conn.buffer, buffer...)
}

func (conn *ConnBuffer) Read(dst []byte) (n int, err error) {
	if len(conn.buffer) > 0 {
		n = copy(dst, conn.buffer)
		conn.buffer = conn.buffer[n:]
		if n == len(dst) {
			return n, nil
		}
	}
	_ = conn.SetReadDeadline(getDeadline())
	netN, err := conn.Conn.Read(dst[n:])
	n += netN
	return n, err
}

//

type conn struct {
	*service
	conn     ConnBuffer
	messages []any
}

func (conn *conn) Close() { _ = conn.conn.Close() }
func (conn *conn) Messages() []any {
	messages := conn.messages
	conn.messages = nil
	return messages
}

func (conn *conn) Send(message any) error {
	bytes, err := conn.Codec().Encode(message)
	if err != nil {
		return err
	}

	bytesLength := len(bytes)
	if bytesLength < 0 || bytesLength > math.MaxUint32 {
		return fmt.Errorf("message length exceeded maximal %v length", bytesLength)
	}

	// conn.logger.Info(fmt.Sprintf("sending '***' of type '%v'", reflect.TypeOf(message).String()))
	length := uint32(bytesLength)
	lengthInByes := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthInByes, length)
	if _, err := conn.conn.Write(append(lengthInByes, bytes...)); err != nil {
		return err
	}

	return nil
}

//

func (s *service) CloseConn(entity ecs.EntityID) {
	connComp, ok := s.Connection().Component().Get(entity)
	if !ok {
		return
	}
	conn := connComp.Conn().(*conn)
	conn.Close()

	s.World().RemoveEntity(entity)
}

func (s *service) PollMessages(entity ecs.EntityID) {
	connComp, ok := s.Connection().Component().Get(entity)
	if !ok {
		return
	}
	conn := connComp.Conn().(*conn)

	messageLengthInBytes := make([]byte, 4)
	if n, err := conn.conn.Read(messageLengthInBytes); n < int(len(messageLengthInBytes)) && errors.Is(err, os.ErrDeadlineExceeded) {
		conn.conn.Buffer(messageLengthInBytes[:n])
		return
	} else if err != nil {
		s.CloseConn(entity)
		return
	}

	messageLength := binary.BigEndian.Uint32(messageLengthInBytes)
	messageBytes := make([]byte, messageLength)
	if n, err := conn.conn.Read(messageBytes); n < int(len(messageLengthInBytes)) && errors.Is(err, os.ErrDeadlineExceeded) {
		conn.conn.Buffer(messageLengthInBytes)
		conn.conn.Buffer(messageBytes[:n])
		return
	} else if err != nil {
		s.CloseConn(entity)
		return
	}

	message, err := s.Codec().Decode(messageBytes)
	if err != nil {
		s.Logger().Log(err)
		return
	}
	// f.logger.Info(fmt.Sprintf("received '***' type '%v'", reflect.TypeOf(message).String()))
	conn.messages = append(conn.messages, message)
}
