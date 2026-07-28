package internal

import (
	"engine/modules/datastructures"
	"engine/modules/ecs/internal/ecstypes"
	"reflect"
)

// impl

type componentArray[Component any] struct {
	entities   entitiesImpl
	equal      func(Component, Component) bool
	empty      Component
	components datastructures.SparseArray[ecstypes.EntityID, Component]

	dependencies []ecstypes.AnyComponentArray
	dirtySets    datastructures.Set[ecstypes.DirtySet]
	beforeGets   []ecstypes.BeforeGet
	onUpsert     []ecstypes.OnMod
	onRemove     []ecstypes.OnMod
}

type Equaler[T any] interface {
	Equal(other T) bool
}

func ComponentComparator[Component any]() func(c1, c2 Component) bool {
	var zero Component
	if _, ok := any(zero).(Equaler[Component]); ok {
		return func(c1, c2 Component) bool {
			return any(c1).(Equaler[Component]).Equal(c2)
		}
	}

	if reflect.TypeFor[Component]().Comparable() {
		return func(c1, c2 Component) bool {
			return any(c1) == any(c2)
		}
	}

	return func(c1, c2 Component) bool { return false }
}

func newComponentArray[Component any](entities entitiesImpl) *componentArray[Component] {
	array := &componentArray[Component]{
		entities: entities,
		equal:    ComponentComparator[Component](),
		// empty: default,
		components: datastructures.NewSparseArray[ecstypes.EntityID, Component](),

		dependencies: nil,
		dirtySets:    datastructures.NewSet[ecstypes.DirtySet](),
		beforeGets:   nil,
		onUpsert:     nil,
		onRemove:     nil,
	}
	return array
}

func (c *componentArray[Component]) Set(entity ecstypes.EntityID, component Component) {
	value, ok := c.components.Get(entity)
	if ok && c.equal(value, component) {
		return
	}
	c.entities.EnsureExists(entity)
	c.components.Set(entity, component)
	for _, onMod := range c.onUpsert {
		onMod(entity)
	}
	for _, dirtySet := range c.dirtySets.Get() {
		if !dirtySet.Ok() {
			c.dirtySets.RemoveElements(dirtySet)
			continue
		}
		dirtySet.Dirty(entity)
	}
}

func (c *componentArray[Component]) SetAny(entity ecstypes.EntityID, anyComponent any) {
	// we use zero in case of invalid call
	component, _ := anyComponent.(Component)
	c.Set(entity, component)
}

func (c *componentArray[Component]) SetEmpty(empty Component) { c.empty = empty }
func (c *componentArray[Component]) GetEmpty() Component      { return c.empty }

func (c *componentArray[Component]) Remove(entity ecstypes.EntityID) {
	if _, ok := c.components.Get(entity); !ok {
		return
	}
	c.components.Remove(entity)
	for _, onMod := range c.onRemove {
		onMod(entity)
	}
	for _, dirtySet := range c.dirtySets.Get() {
		if !dirtySet.Ok() {
			c.dirtySets.RemoveElements(dirtySet)
			continue
		}
		dirtySet.Dirty(entity)
	}
}

func (c *componentArray[Component]) Get(entity ecstypes.EntityID) (Component, bool) {
	for _, beforeGet := range c.beforeGets {
		beforeGet()
	}
	if value, ok := c.components.Get(entity); !ok {
		return c.empty, false
	} else {
		return value, true
	}
}

func (c *componentArray[Component]) GetEntities() []ecstypes.EntityID {
	for _, beforeGet := range c.beforeGets {
		beforeGet()
	}
	return c.components.GetIndices()
}

func (c *componentArray[Component]) GetAny(entity ecstypes.EntityID) (any, bool) {
	return c.Get(entity)
}

//

func (c *componentArray[Component]) AddDependency(dependency ecstypes.AnyComponentArray) {
	c.dependencies = append(c.dependencies, dependency)
	for _, dirtySet := range c.dirtySets.Get() {
		if !dirtySet.Ok() {
			c.dirtySets.RemoveElements(dirtySet)
			continue
		}
		dependency.AddDirtySet(dirtySet)
	}
}
func (c *componentArray[Component]) AddDirtySet(dirtySet ecstypes.DirtySet) {
	if !dirtySet.Ok() {
		c.dirtySets.RemoveElements(dirtySet)
		return
	}
	if _, ok := c.dirtySets.GetIndex(dirtySet); ok {
		return
	}
	for _, entity := range c.GetEntities() {
		dirtySet.Dirty(entity)
	}
	for _, dependency := range c.dependencies {
		dependency.AddDirtySet(dirtySet)
	}
	c.dirtySets.Add(dirtySet)
}
func (c *componentArray[Component]) BeforeGet(listener ecstypes.BeforeGet) {
	// we prepend listener so they are triggered first.
	// if they are truely dependent they will call get again
	//   and BeforeGet will trigger again triggering other listeners
	// else if they won't be called again
	//   then nothing will change
	c.beforeGets = append([]ecstypes.BeforeGet{listener}, c.beforeGets...)
}

func (c *componentArray[Component]) OnUpsert(onUpsert ecstypes.OnMod) {
	c.onUpsert = append(c.onUpsert, onUpsert)
}

func (c *componentArray[Component]) OnRemove(onRemove ecstypes.OnMod) {
	c.onRemove = append(c.onRemove, onRemove)
}

func (c *componentArray[Component]) OnMod(onMod ecstypes.OnMod) {
	c.OnUpsert(onMod)
	c.OnRemove(onMod)
}
