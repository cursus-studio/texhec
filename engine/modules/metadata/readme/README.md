# metadata
## Architecture
allows to store data about objects.
its extremely useful when having in game documentation of objects

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               3             21              5             85
-------------------------------------------------------------------------------
SUM:                             3             21              5             85
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `engine/modules/metadata.Service`

#### method Service Description
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/metadata.DescriptionComponent]`

#### method Service Link
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/metadata.LinkComponent]`

#### method Service Name
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/metadata.NameComponent]`

### type NameComponent
Type: `engine/modules/metadata.NameComponent`

#### property NameComponent Name
Type: `string`

### type DescriptionComponent
Type: `engine/modules/metadata.DescriptionComponent`

#### property DescriptionComponent Description
Type: `string`

### type LinkComponent
Type: `engine/modules/metadata.LinkComponent`

#### property LinkComponent Entity
Type: `engine/modules/ecs.EntityID`

## Functions
### func NewName
Type: `func(name string) engine/modules/metadata.NameComponent`

### func NewDescription
Type: `func(description string) engine/modules/metadata.DescriptionComponent`

### func NewLink
Type: `func(entity engine/modules/ecs.EntityID) engine/modules/metadata.LinkComponent`


## Dependencies
`engine`:
  - `engine.EngineWorld`
  - `engine.World`

`engine/modules/ecs`:
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`

`engine/modules/entityregistry`:
  - `engine/modules/entityregistry.Register`
  - `engine/modules/entityregistry.Service`

`engine/modules/metadata`:
  - `engine/modules/metadata.Description`
  - `engine/modules/metadata.DescriptionComponent`
  - `engine/modules/metadata.Link`
  - `engine/modules/metadata.LinkComponent`
  - `engine/modules/metadata.Name`
  - `engine/modules/metadata.NameComponent`
  - `engine/modules/metadata.NewDescription`
  - `engine/modules/metadata.NewLink`
  - `engine/modules/metadata.NewName`
  - `engine/modules/metadata.Service`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

### Third Party
- `github.com/ogiusek/ioc/v2`