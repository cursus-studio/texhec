package datastructures_test

import (
	"engine/services/datastructures"
	"testing"
)

func TestSpraseSet(t *testing.T) {
	set := datastructures.NewSparseSet[uint8]()
	v1, v2 := uint8(2), uint8(3)

	if !set.Add(v1) {
		t.Errorf("expected v1 to be added")
	}
	if set.Add(v1) {
		t.Errorf("expected v1 to not be added")
	}
	if !set.Add(v2) {
		t.Errorf("expected v2 to be added")
	}

	values := set.GetIndices()
	if len(values) != 2 || min(values[0], values[1]) != v1 || max(values[0], values[1]) != v2 {
		t.Errorf("sparse set has invalid values. expected [%v, %v] in any order but got %v", v1, v2, values)
	}
	if !set.Get(v1) {
		t.Errorf("expected to have v1")
	}
	if !set.Get(v2) {
		t.Errorf("expected to have v2")
	}

	if !set.Remove(v1) {
		t.Errorf("expected v1 be removed")
	}

	values = set.GetIndices()
	if len(values) != 1 || values[0] != v2 {
		t.Errorf("sparse set has invalid values. expected [%v] in any order but got %v", v2, values)
	}
	if set.Get(v1) {
		t.Errorf("expected to not have v1")
	}
	if !set.Get(v2) {
		t.Errorf("expected to have v2")
	}
}
