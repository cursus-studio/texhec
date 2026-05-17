# codec
## Architecture
this module allows us to encode and decode golang objects.

## Types
### type Service
Type: `engine/modules/codec.Service`

#### method Service Decode
Type: `func([]byte) (any, error)`
can return:
ErrInvalidInput

#### method Service Encode
Type: `func(any) ([]byte, error)`

## Variables
### var ErrCannotDecodeBytes
Type: `error`

### var ErrCannotEncodeType
Type: `error`


## Benchmarks
```
$ go test ./... -bench=.
PASS
ok  	engine/modules/codec/test	0.006s
```
## Dependencies
`engine`:
  - `engine.EngineWorld`

`engine/modules/codec`:
  - `engine/modules/codec.ErrCannotDecodeBytes`
  - `engine/modules/codec.ErrCannotEncodeType`
  - `engine/modules/codec.Service`

### Third Party
- `github.com/ogiusek/ioc/v2`