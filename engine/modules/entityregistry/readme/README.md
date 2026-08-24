# entityregistry
## Architecture
`entityregistry` allows us to define entities and components using struct tags
example:
```go

	type OurEntities struct {
	  OurEntity ecs.EntityID `registered_component:"its_value"`
	}

```

## Benchmarks
```
$ go test ./... -bench=.
PASS
ok  	engine/modules/entityregistry/test	0.008s
```
## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               5             31             10            164
-------------------------------------------------------------------------------
SUM:                             5             31             10            164
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `engine/modules/entityregistry.Service`

#### method Service Populate
Type: `func(any) error`
can return ErrExpectedPointerToAStruct

#### method Service Register
Type: `func(structTagKey string, handler func(entity engine/modules/ecs.EntityID, structTagValue string))`

## Variables
### var ErrExpectedPointerToAStruct
Type: `error`

## Functions
### func GetRegistry
Type: `func[Registry any](s engine/modules/entityregistry.Service) (Registry, error)`


## Dependencies
`engine`:
  - `engine.EngineWorld`
  - `engine.Logger`
  - `engine.UUID`
  - `engine.World`

`engine/modules/ecs`:
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`
  - `engine/modules/ecs.World`

`engine/modules/entityregistry`:
  - `engine/modules/entityregistry.ErrExpectedPointerToAStruct`
  - `engine/modules/entityregistry.Register`
  - `engine/modules/entityregistry.Service`

`engine/modules/uuid`:
  - `engine/modules/uuid.Component`
  - `engine/modules/uuid.New`
  - `engine/modules/uuid.NewUUID`

`engine/pkg`:
  - `engine/pkg.Pkg`

### Third Party
- `github.com/ogiusek/ioc/v2`