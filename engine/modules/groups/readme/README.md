# groups
## Architecture
uses bitmasks and allows us to group entities to do not collide despite shared position or
to do not be visible for a camera despite being in its view

## Types
### type Service
Type: `engine/modules/groups.Service`

#### method Service Component
Type: `func() engine/services/ecs.ComponentsArray[engine/modules/groups.GroupsComponent]`

#### method Service Inherit
Type: `func() engine/services/ecs.ComponentsArray[engine/modules/groups.InheritGroupsComponent]`

### type Group
Type: `engine/modules/groups.Group`

### type GroupsComponent
Type: `engine/modules/groups.GroupsComponent`

#### property GroupsComponent Mask
Type: `uint32`
this can be swapped to uint64 etc (remember to swap all uint32 occurencies)

#### method GroupsComponent Ptr
Type: `func() *engine/modules/groups.GroupsComponent`

#### method GroupsComponent Val
Type: `func() engine/modules/groups.GroupsComponent`

#### method GroupsComponent Enable
Type: `func(g engine/modules/groups.Group) *engine/modules/groups.GroupsComponent`

#### method GroupsComponent Enabled
Type: `func(g engine/modules/groups.Group) bool`

#### method GroupsComponent Disable
Type: `func(g engine/modules/groups.Group) *engine/modules/groups.GroupsComponent`

#### method GroupsComponent GetSharedWith
Type: `func(g2 engine/modules/groups.GroupsComponent) engine/modules/groups.GroupsComponent`

#### method GroupsComponent SharesAnyGroup
Type: `func(g2 engine/modules/groups.GroupsComponent) bool`

### type InheritGroupsComponent
Type: `engine/modules/groups.InheritGroupsComponent`

## Variables
### var Groupless
Type: `engine/modules/groups.Group`

## Functions
### func EmptyGroups
Type: `func() engine/modules/groups.GroupsComponent`

### func DefaultGroups
Type: `func() engine/modules/groups.GroupsComponent`


# Benchmarks
```
$ go test ./... -bench=.
PASS
ok  	engine/modules/groups/test	0.007s

```
## Dependencies
`engine`:
  - `engine.EngineWorld`
  - `engine.Hierarchy`
  - `engine.World`

`engine/modules/groups`:
  - `engine/modules/groups.Component`
  - `engine/modules/groups.DefaultGroups`
  - `engine/modules/groups.GroupsComponent`
  - `engine/modules/groups.InheritGroupsComponent`
  - `engine/modules/groups.Service`

`engine/modules/hierarchy`:
  - `engine/modules/hierarchy.Children`
  - `engine/modules/hierarchy.Component`
  - `engine/modules/hierarchy.Parent`
  - `engine/modules/hierarchy.Service`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

`engine/pkg`:
  - `engine/pkg.Pkg`

`engine/services/ecs`:
  - `engine/services/ecs.AddDependency`
  - `engine/services/ecs.AddDirtySet`
  - `engine/services/ecs.BeforeGet`
  - `engine/services/ecs.Clear`
  - `engine/services/ecs.ComponentsArray`
  - `engine/services/ecs.EntityID`
  - `engine/services/ecs.Get`
  - `engine/services/ecs.GetComponentsArray`
  - `engine/services/ecs.NewDirtySet`
  - `engine/services/ecs.Set`
  - `engine/services/ecs.SetEmpty`
  - `engine/services/ecs.World`

### Third Party
`github.com/ogiusek/ioc/v2`