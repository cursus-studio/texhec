package types

import "golang.org/x/exp/constraints"

// set
type SetReader[Stored comparable] interface {
	Get() []Stored
	GetStored(index int) (element Stored, ok bool)
	GetIndex(element Stored) (index int, ok bool)
}
type Set[Stored comparable] interface {
	SetReader[Stored]
	Add(elements ...Stored)
	Set(index int, e Stored)
	Remove(indices ...int)
	RemoveElements(elements ...Stored)
}

// sparse array
type SparseArray[Index constraints.Integer, Value any] interface {
	Get(index Index) (value Value, ok bool)
	GetValues() []Value
	GetIndices() []Index
	Size() int
	// if false then updated
	Set(index Index, value Value) (added bool)
	Remove(index Index) (removed bool)
}

// sparse set
type SparseSetReader[Index constraints.Integer] interface {
	Get(index Index) (ok bool)
	GetIndices() []Index
}
type SparseSet[Index constraints.Integer] interface {
	SparseSetReader[Index]
	Add(index Index) (added bool)
	Remove(index Index) (removed bool)
}

// tracking array
type Change[Stored comparable] struct {
	Index int
	From  *Stored
}
type TrackingArray[Stored comparable] interface {
	Get() []Stored
	Add(elements ...Stored)
	Set(index int, e Stored)
	Remove(indices ...int)

	Changes() []Change[Stored]
	ClearChanges()
}

func NewChange[Stored comparable](index int, from *Stored) Change[Stored] {
	return Change[Stored]{index, from}
}
