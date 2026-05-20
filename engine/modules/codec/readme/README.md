# codec
## Architecture
this module allows us to encode and decode golang objects.

## Benchmarks
```
$ go test ./... -bench=.
PASS
ok  	engine/modules/codec/test	0.006s
```
## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               6             29              3            165
-------------------------------------------------------------------------------
SUM:                             6             29              3            165
-------------------------------------------------------------------------------
```
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


## Dependencies
`engine`:
  - `engine.EngineWorld`

`engine/modules/codec`:
  - `engine/modules/codec.ErrCannotDecodeBytes`
  - `engine/modules/codec.ErrCannotEncodeType`
  - `engine/modules/codec.Service`

### Third Party
- `github.com/ogiusek/ioc/v2`