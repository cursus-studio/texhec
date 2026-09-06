package collisions

import (
	"engine"
	"engine/modules/assets"
	"engine/modules/collider"
	"engine/modules/datastructures"
	"engine/modules/ecs"
	"engine/modules/groups"
	"errors"
	"slices"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ogiusek/ioc/v2"
)

type service struct {
	// shared
	engine.EngineWorld `inject:""`

	collider ecs.ComponentArray[collider.Component]

	// tracking
	collidersDirtySet ecs.DirtySet
	chunkSize         float32
	chunks            map[mgl32.Vec2]datastructures.Set[ecs.EntityID]
	entitiesPositions map[ecs.EntityID][]mgl32.Vec2

	rayFallTroughPolicies []collider.FallTroughPolicy
}

func NewService(c ioc.Dic,
	chunkSize float32,
) collider.Service {
	s := ioc.GetServices[*service](c)

	s.collider = ecs.GetComponentArray[collider.Component](s.World())

	s.collidersDirtySet = ecs.NewDirtySet()
	s.Transform().AddDirtySet(s.collidersDirtySet)
	s.collider.AddDirtySet(s.collidersDirtySet)

	s.chunkSize = chunkSize
	s.chunks = make(map[mgl32.Vec2]datastructures.Set[ecs.EntityID])
	s.entitiesPositions = make(map[ecs.EntityID][]mgl32.Vec2)
	s.rayFallTroughPolicies = make([]collider.FallTroughPolicy, 0)

	return s
}

func floorF32ToInt(num float32) int {
	// return int(math.Floor(float64(num)))
	return int(num)
}

func (s *service) getPositions(aabb collider.AABB) []mgl32.Vec2 {
	minPos, maxPos := aabb.Min.Vec2(), aabb.Max.Vec2()
	minGridX := floorF32ToInt(minPos.X() / s.chunkSize)
	minGridY := floorF32ToInt(minPos.Y() / s.chunkSize)
	maxGridX := floorF32ToInt(maxPos.X() / s.chunkSize)
	maxGridY := floorF32ToInt(maxPos.Y() / s.chunkSize)

	var positions []mgl32.Vec2
	for x := minGridX; x <= maxGridX; x++ {
		for y := minGridY; y <= maxGridY; y++ {
			tileX := float32(x) * s.chunkSize
			tileY := float32(y) * s.chunkSize
			tileCenter := mgl32.Vec2{tileX, tileY}

			positions = append(positions, tileCenter)
		}
	}

	return positions
}

func (s *service) ChunkSize() float32                                      { return s.chunkSize }
func (s *service) Chunks() map[mgl32.Vec2]datastructures.Set[ecs.EntityID] { return s.chunks }

// tracking

func (s *service) ApplyChanges() {
	entities := s.collidersDirtySet.Get()
	s.Remove(entities...)
	for _, entity := range entities {
		if _, ok := s.collider.Get(entity); !ok {
			continue
		}
		aabb := TransformAABB(s.Transform(), entity)
		positions := s.getPositions(aabb)
		s.entitiesPositions[entity] = positions
		for _, position := range positions {
			arr, ok := s.chunks[position]
			if !ok {
				arr = datastructures.NewSet[ecs.EntityID]()
			}
			arr.Add(entity)
			s.chunks[position] = arr
		}
	}
}

func (s *service) Remove(entities ...ecs.EntityID) {
	for _, entity := range entities {
		positions, ok := s.entitiesPositions[entity]
		if !ok {
			continue
		}
		delete(s.entitiesPositions, entity)
		for _, position := range positions {
			arr, ok := s.chunks[position]
			if !ok {
				continue
			}
			arr.RemoveElements(entity)
			if len(arr.Get()) == 0 {
				delete(s.chunks, position)
				continue
			}
		}
	}
}

//

func (s *service) Component() ecs.ComponentArray[collider.Component] { return s.collider }

func (s *service) CollidesWithRay(entity ecs.EntityID, ray collider.Ray) *collider.ObjectRayCollision {
	s.ApplyChanges()
	entityGroups, ok := s.Groups().Component().Get(entity)
	if !ok {
		entityGroups = groups.DefaultGroups()
	}
	if entityGroups.GetSharedWith(ray.Groups).Mask == 0 {
		return nil
	}

	aabb := TransformAABB(s.Transform(), entity)
	if ok, _ := RayAABBIntersect(ray, aabb); !ok {
		return nil
	}

	colliderComponent, ok := s.collider.Get(entity)
	if !ok {
		return nil
	}
	colliderAsset, err := assets.GetAsset[collider.ColliderAsset](s.Assets(), colliderComponent.ID)
	if err != nil {
		// invalid internal state
		s.Logger().Log(err)
		return nil
	}

	//

	ray.Apply(s.Transform().Mat4(entity).Inv())

	aabbs := colliderAsset.AABBs()
	ranges := colliderAsset.Ranges()
	polygons := colliderAsset.Polygons()

	rangesToVisit := []collider.Range{}
	if len(ranges) > 0 {
		rangesToVisit = append(rangesToVisit, collider.NewRange(collider.Branch, 0, 1))
	}

	var closestHit *collider.RayHit

	for len(rangesToVisit) > 0 {
		currentRange := rangesToVisit[len(rangesToVisit)-1]
		rangesToVisit = rangesToVisit[:len(rangesToVisit)-1]

		if currentRange.Target == collider.Branch {
			for i := currentRange.First; i < currentRange.First+currentRange.Count; i++ {
				aabb := aabbs[i]
				intersects, _ := RayAABBIntersect(ray, aabb)
				if !intersects {
					continue
				}
				rangesToVisit = append(rangesToVisit, ranges[i])
			}
		} else if currentRange.Target == collider.Leaf {
			polygons := polygons[currentRange.First : currentRange.First+currentRange.Count]
			for _, polygon := range polygons {
				intersect, dist := RayTriangleIntersect(ray, polygon)
				if !intersect {
					continue
				}
				if closestHit != nil && closestHit.Distance < dist {
					continue
				}

				ray := ray
				ray.MaxDistance = dist

				normal := polygon.B.Sub(polygon.A).Cross(polygon.C.Sub(polygon.A)).Normalize()
				hit := collider.NewRayHit(ray, normal)
				closestHit = &hit
			}
		}
	}

	if closestHit == nil {
		return nil
	}

	collision := collider.NewObjectRayCollision(entity, *closestHit)

	for _, rayFallTroughPolicy := range s.rayFallTroughPolicies {
		if fallThrough := rayFallTroughPolicy.FallThrough(collision); fallThrough {
			return nil
		}
	}
	return &collision
}

func (s *service) CollidesWithObject(entityA ecs.EntityID, entityB ecs.EntityID) *collider.ObjectObjectCollision {
	s.ApplyChanges()
	s.Logger().Log(errors.New("501"))
	return nil
}

func (s *service) Raycast(ray collider.Ray) *collider.ObjectRayCollision {
	s.ApplyChanges()
	chunkSize := s.ChunkSize()
	gridX := floorF32ToInt(ray.Pos[0] / chunkSize)
	gridY := floorF32ToInt(ray.Pos[1] / chunkSize)
	chunkCoord := mgl32.Vec2{
		float32(gridX) * chunkSize,
		float32(gridY) * chunkSize,
	}

	chunk, ok := s.Chunks()[chunkCoord]
	if !ok {
		return nil
	}

	var closestHit *collider.RayHit
	var closestEntity ecs.EntityID

	for _, entity := range chunk.Get() {
		collision := s.CollidesWithRay(entity, ray)
		if collision == nil {
			continue
		}

		if closestHit != nil && closestHit.Distance < collision.Hit.Distance {
			continue
		}

		hit := collision.Hit

		closestHit = &hit
		closestEntity = entity
	}

	if closestHit == nil {
		return nil
	}

	collision := collider.NewObjectRayCollision(closestEntity, *closestHit)
	return &collision
}

func (s *service) RaycastAll(ray collider.Ray) []collider.ObjectRayCollision {
	s.ApplyChanges()
	chunkSize := s.ChunkSize()
	gridX := floorF32ToInt(ray.Pos[0] / chunkSize)
	gridY := floorF32ToInt(ray.Pos[1] / chunkSize)
	chunkCoord := mgl32.Vec2{
		float32(gridX) * chunkSize,
		float32(gridY) * chunkSize,
	}

	chunk, ok := s.Chunks()[chunkCoord]
	if !ok {
		return nil
	}

	collisions := []collider.ObjectRayCollision{}

	for _, entity := range chunk.Get() {
		collision := s.CollidesWithRay(entity, ray)
		if collision == nil {
			continue
		}

		collisions = append(collisions, collider.NewObjectRayCollision(entity, collision.Hit))
	}

	slices.SortFunc(collisions, func(a, b collider.ObjectRayCollision) int {
		if a.Hit.Distance < b.Hit.Distance {
			return -1
		}
		if a.Hit.Distance > b.Hit.Distance {
			return 1
		}
		return 0
	})

	return collisions
}

func (s *service) NarrowCollisions(entity ecs.EntityID) []ecs.EntityID {
	s.ApplyChanges()
	s.Logger().Log(errors.New("501"))
	return nil
}
func (s *service) AddRayFallThroughPolicy(rayFallTroughPolicy collider.FallTroughPolicy) {
	s.rayFallTroughPolicies = append(s.rayFallTroughPolicies, rayFallTroughPolicy)
}
