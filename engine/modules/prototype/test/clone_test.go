package test

import "testing"

func TestClone(t *testing.T) {
	s := NewSetup()

	cloned := s.World().NewEntity()
	cloned1Comp := Cloned1Component{1}
	cloned2Comp := Cloned2Component{1}
	notClonedComp := NotClonedComponent{1}
	s.Cloned1.Set(cloned, cloned1Comp)
	s.Cloned2.Set(cloned, cloned2Comp)
	s.NotCloned.Set(cloned, notClonedComp)

	clone := s.Prototype().Clone(cloned)
	if clonedComp1Cp, ok := s.Cloned1.Get(clone); !ok {
		t.Errorf("clonedComponent didn't got cloned ")
		return
	} else if cloned1Comp != clonedComp1Cp {
		t.Errorf("expected clonedComponent to be \"%v\" but got \"%v\"", cloned1Comp, clonedComp1Cp)
		return
	}
	if clonedComp2Cp, ok := s.Cloned2.Get(clone); !ok {
		t.Errorf("clonedComponent didn't got cloned ")
		return
	} else if cloned2Comp != clonedComp2Cp {
		t.Errorf("expected clonedComponent to be \"%v\" but got \"%v\"", cloned2Comp, clonedComp2Cp)
		return
	}
	if _, ok := s.NotCloned.Get(clone); ok {
		t.Errorf("expected not cloned component to do not be cloned")
	}
}
