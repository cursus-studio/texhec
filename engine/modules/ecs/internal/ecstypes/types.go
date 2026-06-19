package ecstypes

import "reflect"

// dirty set

type DirtySet interface {
	// get also clears
	Get() []EntityID
	Dirty(EntityID)
	Clear()

	Ok() bool
	Release()
}

// world

type EntityID uint32

func (id EntityID) Index() int { return int(id) }

type World interface {
	// entities
	GetEntities() []EntityID
	EntityExists(EntityID) bool

	NewEntity() EntityID
	EnsureExists(EntityID)
	RemoveEntity(EntityID)

	// components
	// there is a public method [GetComponentArray] to call generic method in Go

	// tooling
	WarmUp()
}

// components getters

type Component = any

// component arrays

type BeforeGet = func()
type OnMod = func(EntityID)

type AnyComponentArray interface {
	GetAny(entity EntityID) (Component, bool)
	GetEntities() []EntityID

	// when type doesn't match error is returned
	SetAny(EntityID, Component)
	Remove(EntityID)

	// configuration
	// on dependency change its also applied here
	AddDependency(AnyComponentArray)
	AddDirtySet(DirtySet)
	BeforeGet(BeforeGet)

	OnUpsert(OnMod)
	OnRemove(OnMod)
	// is called OnUpsert and OnRemove
	OnMod(OnMod)
}
type ComponentArray[Component any] interface {
	AnyComponentArray
	Get(entity EntityID) (Component, bool)

	Set(EntityID, Component)

	// configuration
	SetEmpty(Component)
	// it gets called imidiately with current empty
	GetEmpty() Component
}

//

func ComponentComparator[Component any]() func(c1, c2 Component) bool {
	equal := func(Component, Component) bool { return false }
	if reflect.TypeFor[Component]().Comparable() {
		equal = func(c1, c2 Component) bool { return any(c1) == any(c2) }
	}
	return equal
}
