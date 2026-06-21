package datastructures

import (
	"engine/modules/datastructures/internal"
	"engine/modules/datastructures/internal/types"
	"sync"

	"golang.org/x/exp/constraints"
)

// set
type SetReader[Stored comparable] = types.SetReader[Stored]
type Set[Stored comparable] = types.Set[Stored]

func NewSet[Stored comparable]() Set[Stored] {
	return internal.NewSet[Stored]()
}

// sparse array
type SparseArray[Index constraints.Integer, Value any] = types.SparseArray[Index, Value]

func NewSparseArray[Index constraints.Integer, Value any]() SparseArray[Index, Value] {
	return internal.NewSparseArray[Index, Value]()
}

// sparse set
type SparseSetReader[Index constraints.Integer] = types.SparseSetReader[Index]
type SparseSet[Index constraints.Integer] = types.SparseSet[Index]

func NewSparseSet[Index constraints.Integer]() SparseSet[Index] {
	return internal.NewSparseSet[Index]()
}
func NewSparseSetWithPaging[Index constraints.Integer]() SparseSet[Index] {
	return internal.NewSparseSetWithPaging[Index]()
}

type Change[Stored comparable] = types.Change[Stored]
type TrackingArray[Stored comparable] = types.TrackingArray[Stored]

func NewTrackingArray[Stored comparable]() TrackingArray[Stored] {
	return internal.NewTrackingArray[Stored]()
}
func NewThreadSafeTrackingArray[Stored comparable](mutex sync.Locker) TrackingArray[Stored] {
	return internal.NewThreadSafeTrackingArray[Stored](mutex)
}
