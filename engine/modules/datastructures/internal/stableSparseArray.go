package internal

import (
	"engine/modules/datastructures/internal/types"

	"golang.org/x/exp/constraints"
)

type stableSparseArray[Index constraints.Integer, Value any] struct {
	EmptyValue    Index
	ValuesIndices []Index // here some indices have special meaning (read constants above)

	// both arrays below indices correspond
	Values  []Value
	Indices []Index // here value means index in sparse array
}

func NewStableSparseArray[Index constraints.Integer, Value any]() types.SparseArray[Index, Value] {
	var zero Index
	return &stableSparseArray[Index, Value]{
		EmptyValue: ^zero,
	}
}

func (a *stableSparseArray[Index, Value]) Get(index Index) (Value, bool) {
	if int(index) >= len(a.ValuesIndices) {
		var zero Value
		return zero, false
	}

	valueIndex := a.ValuesIndices[index]
	if valueIndex == a.EmptyValue {
		var zero Value
		return zero, false
	}

	value := a.Values[valueIndex]
	return value, true
}

func (a *stableSparseArray[Index, Value]) GetValues() []Value  { return a.Values }
func (a *stableSparseArray[Index, Value]) GetIndices() []Index { return a.Indices }

func (a *stableSparseArray[Index, Value]) Size() int { return len(a.Indices) }

func (a *stableSparseArray[Index, Value]) Set(index Index, value Value) bool {
	for int(index) >= len(a.ValuesIndices) {
		a.ValuesIndices = append(a.ValuesIndices, a.EmptyValue)
	}

	valueIndex := a.ValuesIndices[index]

	if valueIndex == a.EmptyValue {
		a.ValuesIndices[index] = Index(len(a.Values))
		a.Values = append(a.Values, value)
		a.Indices = append(a.Indices, index)
		return true
	}

	a.Values[valueIndex] = value
	a.Indices[valueIndex] = index
	return false
}

func (a *stableSparseArray[Index, Value]) Remove(index Index) bool {
	if int(index) >= len(a.ValuesIndices) {
		return false
	}

	valueIndex := a.ValuesIndices[index]
	if valueIndex == a.EmptyValue {
		return false
	}

	a.ValuesIndices[index] = a.EmptyValue

	for i := int(valueIndex); i < len(a.Values)-1; i++ {
		a.Values[i] = a.Values[i+1]
		a.Indices[i] = a.Indices[i+1]

		a.ValuesIndices[a.Indices[i]] = Index(i)
	}

	a.Values = a.Values[:len(a.Values)-1]
	a.Indices = a.Indices[:len(a.Indices)-1]
	return true
}
