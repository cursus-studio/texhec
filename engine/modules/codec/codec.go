package codec

import "errors"

var (
	ErrCannotDecodeBytes error = errors.New("cannot decode bytes")
	ErrCannotEncodeType  error = errors.New("cannot encode type")
)

type Service interface {
	Encode(any) ([]byte, error)

	// can return:
	// ErrInvalidInput
	Decode([]byte) (any, error)
}
