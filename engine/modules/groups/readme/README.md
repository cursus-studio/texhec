# groups
## Architecture
uses bitmasks and allows us to group entities to do not collide despite shared position or
to do not be visible for a camera despite being in its view

## Benchmarks
```
$ go test ./... -bench=.
PASS
ok  	engine/modules/groups/test	0.006s
```
## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               7             42             75            156
-------------------------------------------------------------------------------
SUM:                             7             42             75            156
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `engine/modules/groups.Service`

#### method Service Component
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/groups.GroupsComponent]`

#### method Service Inherit
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/hierarchy.InheritComponent[engine/modules/groups.GroupsComponent]]`

#### method Service InheritGroups
Type: `func(engine/modules/ecs.EntityID)`

### type Group
Type: `engine/modules/groups.Group`

### type GroupsComponent
Type: `engine/modules/groups.GroupsComponent`

#### property GroupsComponent Mask
Type: `uint32`
this can be swapped to uint64 etc (remember to swap all uint32 occurencies)

#### method GroupsComponent Enable
Type: `func(g engine/modules/groups.Group) engine/modules/groups.GroupsComponent`

#### method GroupsComponent Enabled
Type: `func(g engine/modules/groups.Group) bool`

#### method GroupsComponent Disable
Type: `func(g engine/modules/groups.Group) engine/modules/groups.GroupsComponent`

#### method GroupsComponent GetSharedWith
Type: `func(g2 engine/modules/groups.GroupsComponent) engine/modules/groups.GroupsComponent`

#### method GroupsComponent SharesAnyGroup
Type: `func(g2 engine/modules/groups.GroupsComponent) bool`

## Variables
### var Groupless
Type: `engine/modules/groups.Group`

## Functions
### func EmptyGroups
Type: `func() engine/modules/groups.GroupsComponent`

### func DefaultGroups
Type: `func() engine/modules/groups.GroupsComponent`


## Dependencies
`engine`:
  - `engine.EngineWorld`
  - `engine.World`

`engine/modules/ecs`:
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`
  - `engine/modules/ecs.World`

`engine/modules/groups`:
  - `engine/modules/groups.Component`
  - `engine/modules/groups.DefaultGroups`
  - `engine/modules/groups.GroupsComponent`
  - `engine/modules/groups.Service`

`engine/modules/hierarchy`:
  - `engine/modules/hierarchy.Inherit`
  - `engine/modules/hierarchy.InheritComponent`
  - `engine/modules/hierarchy.Service`
  - `engine/modules/hierarchy.ServiceT`

`engine/modules/hierarchy/pkg`:
  - `engine/modules/hierarchy/pkg.PkgT`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

`engine/pkg`:
  - `engine/pkg.Pkg`

### Third Party
- `github.com/ogiusek/ioc/v2`