package internal

import (
	"encoding/binary"
	"engine/modules/connection"
	"engine/modules/ecs"
	"engine/modules/uuid"
	"errors"
	"fmt"
	"math"
	"net"
	"os"
	"reflect"

	"github.com/ogiusek/events"
)

func setMsgCtx(msg any, ctx connection.MsgCtx) (any, bool) {
	msgValuePointer := reflect.New(reflect.TypeOf(msg))
	msgValuePointer.Elem().Set(reflect.ValueOf(msg))
	msgSetter, ok := msgValuePointer.Interface().(connection.MsgCtxSetter)
	if !ok {
		return nil, false
	}
	msgSetter.SetCtx(ctx)
	msgAny := msgValuePointer.Elem().Interface()
	return msgAny, ok
}

func implementsMsgCtxSetter(msg any) bool {
	t := reflect.PointerTo(reflect.TypeOf(msg))
	return t.Implements(reflect.TypeFor[connection.MsgCtxSetter]())
}

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
	conn ConnBuffer
}

func (conn *conn) Close() { _ = conn.conn.Close() }

func (conn *conn) Send(message any) error {
	if !implementsMsgCtxSetter(message) {
		return fmt.Errorf("message doesn't inherit connection.MsgCtx")
	}

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

	msg, err := s.Codec().Decode(messageBytes)
	if err != nil {
		s.Logger().Log(err)
		return
	}
	msg, ok = setMsgCtx(msg, connection.NewMsgCtx(entity))
	if !ok {
		s.Logger().Warn(fmt.Errorf("received message doesn't implement engine.connection.msgCtxSetter"))
		return
	}
	// f.logger.Info(fmt.Sprintf("received '***' type '%v'", reflect.TypeOf(message).String()))
	events.EmitAny(s.Events(), msg)
}

func (s *service) OnSetUUID(setUUID connection.SetConnUUIDDTO) {
	s.UUID().Component().Set(setUUID.ConnEntity, uuid.New(setUUID.UUID))
}
