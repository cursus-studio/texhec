# logger
## Architecture
defines loging abstraction which can be easily expanded to log information in GUI or sent to developer

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               3             25             12            121
-------------------------------------------------------------------------------
SUM:                             3             25             12            121
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `engine/modules/logger.Service`

#### method Service Fatal
Type: `func(error)`
wrapper around Log(errors.Join(ErrFatal, err))

#### method Service Info
Type: `func(error)`
wrapper around Log(errors.Join(ErrInfo, err))

#### method Service Log
Type: `func(error)`
error is composed from:
- multiple meta tags which can contain audience or severity (optional)
- error message

#### method Service Warn
Type: `func(error)`
wrapper around Log(err)

## Variables
### var ErrInfo
Type: `error`

### var ErrFatal
Type: `error`
warn is default

## Functions
### func IsWarning
Type: `func(meta error) bool`


## Dependencies
`engine/modules/logger`:
  - `engine/modules/logger.ErrFatal`
  - `engine/modules/logger.ErrInfo`
  - `engine/modules/logger.Service`

### Third Party
- `github.com/ogiusek/ioc/v2`