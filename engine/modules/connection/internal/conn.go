package internal

import (
	"encoding/binary"
	"fmt"
	"math"
	"net"
)

type conn struct {
	*service
	conn     net.Conn
	messages chan any
}

func (conn *conn) Close() error       { return conn.conn.Close() }
func (conn *conn) Messages() chan any { return conn.messages }

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
