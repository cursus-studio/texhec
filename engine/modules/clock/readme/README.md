# clock
## Architecture
defines clock and unified date format.
It's a simple time module wrapper.

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               3             15              0             45
Markdown                         1              0              0              2
-------------------------------------------------------------------------------
SUM:                             4             15              0             47
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `engine/modules/clock.Service`

#### method Service GetDateFormat
Type: `func() engine/modules/clock.DateFormat`

#### method Service Now
Type: `func() time.Time`

#### method Service SetDateFormat
Type: `func(engine/modules/clock.DateFormat)`

### type DateFormat
Type: `engine/modules/clock.DateFormat`

#### method DateFormat String
Type: `func() string`

#### method DateFormat Parse
Type: `func(date string) (time.Time, error)`

#### method DateFormat Format
Type: `func(date time.Time) string`

## Functions
### func NewDateFormat
Type: `func(date string) engine/modules/clock.DateFormat`


## Dependencies
`engine/modules/clock`:
  - `engine/modules/clock.DateFormat`
  - `engine/modules/clock.Service`

### Third Party
- `github.com/ogiusek/ioc/v2`