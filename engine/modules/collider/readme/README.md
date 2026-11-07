# collider
## Architecture
this module allows us to check collision between objects and rays
current algorithm used under the hood is spatial algorithm.

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                              10            127             22            530
Markdown                         1              1              0              3
-------------------------------------------------------------------------------
SUM:                            11            128             22            533
-------------------------------------------------------------------------------
```
## TODO
Implement `CollidesWithObject`

Change main algorithm from spatial algorithm to tree algorithm.
Create methods to perform only shallow comparisons or only deep comparisons.

## Dependencies
`engine`:
  - `engine.Assets`
  - `engine.EngineWorld`
  - `engine.Groups`
  - `engine.Logger`
  - `engine.Transform`
  - `engine.World`

`engine/modules/assets`:
  - `engine/modules/assets.Asset`
  - `engine/modules/assets.GetAsset`

`engine/modules/collider`:
  - `engine/modules/collider.A`
  - `engine/modules/collider.AABB`
  - `engine/modules/collider.AABBs`
  - `engine/modules/collider.Apply`
  - `engine/modules/collider.B`
  - `engine/modules/collider.Branch`
  - `engine/modules/collider.C`
  - `engine/modules/collider.ColliderAsset`
  - `engine/modules/collider.ColliderComponent`
  - `engine/modules/collider.Count`
  - `engine/modules/collider.Direction`
  - `engine/modules/collider.Distance`
  - `engine/modules/collider.FallThrough`
  - `engine/modules/collider.FallTroughPolicy`
  - `engine/modules/collider.First`
  - `engine/modules/collider.Groups`
  - `engine/modules/collider.Hit`
  - `engine/modules/collider.ID`
  - `engine/modules/collider.Leaf`
  - `engine/modules/collider.Max`
  - `engine/modules/collider.MaxDistance`
  - `engine/modules/collider.Min`
  - `engine/modules/collider.NewAABB`
  - `engine/modules/collider.NewObjectRayCollision`
  - `engine/modules/collider.NewRange`
  - `engine/modules/collider.NewRayHit`
  - `engine/modules/collider.ObjectObjectCollision`
  - `engine/modules/collider.ObjectRayCollision`
  - `engine/modules/collider.Polygon`
  - `engine/modules/collider.Polygons`
  - `engine/modules/collider.Pos`
  - `engine/modules/collider.Range`
  - `engine/modules/collider.Ranges`
  - `engine/modules/collider.Ray`
  - `engine/modules/collider.RayHit`
  - `engine/modules/collider.Service`
  - `engine/modules/collider.Target`

`engine/modules/datastructures`:
  - `engine/modules/datastructures.NewSet`
  - `engine/modules/datastructures.Set`

`engine/modules/ecs`:
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.DirtySet`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`
  - `engine/modules/ecs.NewDirtySet`

`engine/modules/groups`:
  - `engine/modules/groups.DefaultGroups`
  - `engine/modules/groups.GetSharedWith`
  - `engine/modules/groups.Groups`
  - `engine/modules/groups.GroupsComponent`
  - `engine/modules/groups.Mask`

`engine/modules/transform`:
  - `engine/modules/transform.AbsolutePos`
  - `engine/modules/transform.AbsoluteRotation`
  - `engine/modules/transform.AbsoluteSize`
  - `engine/modules/transform.AddDirtySet`
  - `engine/modules/transform.Mat4`
  - `engine/modules/transform.Pos`
  - `engine/modules/transform.Rotation`
  - `engine/modules/transform.Service`
  - `engine/modules/transform.Size`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

### Third Party
- `github.com/go-gl/mathgl/mgl32`
- `github.com/ogiusek/ioc/v2`