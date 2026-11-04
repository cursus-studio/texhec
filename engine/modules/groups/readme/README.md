# groups
## Architecture
uses bitmasks and allows us to group entities to do not collide despite shared position or
to do not be visible for a camera despite being in its view

## Benchmarks
```
$ go test ./... -bench=.
PASS
ok  	engine/modules/groups/test	0.013s
```
## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               6             39              4            152
-------------------------------------------------------------------------------
SUM:                             6             39              4            152
-------------------------------------------------------------------------------
```
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
  - `engine/modules/groups.DefaultGroups`
  - `engine/modules/groups.Groups`
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