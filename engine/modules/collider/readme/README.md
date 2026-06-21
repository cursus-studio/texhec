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

## Types
### type Service
Type: `engine/modules/collider.Service`

#### method Service AddRayFallThroughPolicy
Type: `func(engine/modules/collider.FallTroughPolicy)`

#### method Service CollidesWithObject
Type: `func(engine/modules/ecs.EntityID, engine/modules/ecs.EntityID) *engine/modules/collider.ObjectObjectCollision`

#### method Service CollidesWithRay
Type: `func(engine/modules/ecs.EntityID, engine/modules/collider.Ray) *engine/modules/collider.ObjectRayCollision`
todo add collision groups
narrow

#### method Service Component
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/collider.Component]`

#### method Service NarrowCollisions
Type: `func(engine/modules/ecs.EntityID) []engine/modules/ecs.EntityID`

#### method Service Raycast
Type: `func(engine/modules/collider.Ray) *engine/modules/collider.ObjectRayCollision`
broad

#### method Service RaycastAll
Type: `func(engine/modules/collider.Ray) []engine/modules/collider.ObjectRayCollision`

### type ColliderAsset
Type: `engine/modules/collider.ColliderAsset`

#### method ColliderAsset AABBs
Type: `func() []engine/modules/collider.AABB`
first aabb is the entry point

#### method ColliderAsset Polygons
Type: `func() []engine/modules/collider.Polygon`

#### method ColliderAsset Ranges
Type: `func() []engine/modules/collider.Range`
[]Range element index corresponds to []AABB element index

#### method ColliderAsset Release
Type: `func()`

### type FallTroughPolicy
Type: `engine/modules/collider.FallTroughPolicy`

#### method FallTroughPolicy FallThrough
Type: `func(target engine/modules/collider.ObjectRayCollision) bool`
position is normalized to be between -1 and 1

### type RangeTarget
Type: `engine/modules/collider.RangeTarget`

### type Range
Type: `engine/modules/collider.Range`

#### property Range Target
Type: `engine/modules/collider.RangeTarget`

#### property Range First
Type: `uint32`

#### property Range Count
Type: `uint32`

### type Polygon
Type: `engine/modules/collider.Polygon`
todo add normals and store aabb

#### property Polygon A
Type: `github.com/go-gl/mathgl/mgl32.Vec3`

#### property Polygon B
Type: `github.com/go-gl/mathgl/mgl32.Vec3`

#### property Polygon C
Type: `github.com/go-gl/mathgl/mgl32.Vec3`

### type ObjectRayCollision
Type: `engine/modules/collider.ObjectRayCollision`

#### property ObjectRayCollision Entity
Type: `engine/modules/ecs.EntityID`

#### property ObjectRayCollision Hit
Type: `engine/modules/collider.RayHit`

### type ObjectObjectCollision
Type: `engine/modules/collider.ObjectObjectCollision`

#### property ObjectObjectCollision PolygonPairs
Type: `[][2]engine/modules/collider.Polygon`

### type Component
Type: `engine/modules/collider.Component`

#### property Component ID
Type: `engine/modules/ecs.EntityID`

### type AABB
Type: `engine/modules/collider.AABB`

#### property AABB Min
Type: `github.com/go-gl/mathgl/mgl32.Vec3`

#### property AABB Max
Type: `github.com/go-gl/mathgl/mgl32.Vec3`

### type Ray
Type: `engine/modules/collider.Ray`

#### property Ray Pos
Type: `github.com/go-gl/mathgl/mgl32.Vec3`

#### property Ray Direction
Type: `github.com/go-gl/mathgl/mgl32.Vec3`

#### property Ray MaxDistance
Type: `float32`
max length is either 0 symbolizing infinity or a potive number

#### property Ray Groups
Type: `engine/modules/groups.GroupsComponent`
collision mask

#### method Ray Apply
Type: `func(transform github.com/go-gl/mathgl/mgl32.Mat4)`

#### method Ray HitPoint
Type: `func() github.com/go-gl/mathgl/mgl32.Vec3`

### type RayHit
Type: `engine/modules/collider.RayHit`

#### property RayHit Point
Type: `github.com/go-gl/mathgl/mgl32.Vec3`

#### property RayHit Normal
Type: `github.com/go-gl/mathgl/mgl32.Vec3`

#### property RayHit Distance
Type: `float32`

## Variables
### var Leaf
Type: `engine/modules/collider.RangeTarget`

### var Branch
Type: `engine/modules/collider.RangeTarget`

## Functions
### func NewRange
Type: `func(target engine/modules/collider.RangeTarget, first uint32, count uint32) engine/modules/collider.Range`

### func NewPolygon
Type: `func(a github.com/go-gl/mathgl/mgl32.Vec3, b github.com/go-gl/mathgl/mgl32.Vec3, c github.com/go-gl/mathgl/mgl32.Vec3) engine/modules/collider.Polygon`

### func NewColliderAsset
Type: `func(aabbs []engine/modules/collider.AABB, ranges []engine/modules/collider.Range, polygons []engine/modules/collider.Polygon) engine/modules/collider.ColliderAsset`

### func NewObjectRayCollision
Type: `func(entity engine/modules/ecs.EntityID, hit engine/modules/collider.RayHit) engine/modules/collider.ObjectRayCollision`

### func NewObjectObjectCollision
Type: `func(pairs [][2]engine/modules/collider.Polygon) engine/modules/collider.ObjectObjectCollision`

### func NewCollider
Type: `func(id engine/modules/ecs.EntityID) engine/modules/collider.Component`

### func NewAABB
Type: `func(min github.com/go-gl/mathgl/mgl32.Vec3, max github.com/go-gl/mathgl/mgl32.Vec3) engine/modules/collider.AABB`

### func NewRay
Type: `func(pos github.com/go-gl/mathgl/mgl32.Vec3, direction github.com/go-gl/mathgl/mgl32.Vec3, maxDistance float32, groups engine/modules/groups.GroupsComponent) engine/modules/collider.Ray`

### func NewRayHit
Type: `func(ray engine/modules/collider.Ray, normal github.com/go-gl/mathgl/mgl32.Vec3) engine/modules/collider.RayHit`


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
  - `engine/modules/collider.Component`
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
  - `engine/modules/groups.Component`
  - `engine/modules/groups.DefaultGroups`
  - `engine/modules/groups.GetSharedWith`
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