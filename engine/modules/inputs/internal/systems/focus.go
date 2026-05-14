package systems

import "engine/modules/inputs"

func (s *inputsSystem) OnDefaultFocus(e inputs.DefaultFocusEvent) {
	for _, entity := range s.Inputs().DefaultFocused().GetEntities() {
		s.Inputs().DefaultFocused().Remove(entity)
	}
	s.Inputs().DefaultFocused().Set(e.Entity, inputs.NewDefaultFocused())
}
func (s *inputsSystem) OnFocus(e inputs.FocusEvent) {
	for _, focusedEntity := range s.Inputs().Focused().GetEntities() {
		s.Inputs().Focused().Remove(focusedEntity)
	}
	s.Inputs().Focused().Set(e.Entity, inputs.NewFocused())
}
func (s *inputsSystem) OnUnfocus(inputs.UnfocusEvent) {
	for _, focusedEntity := range s.Inputs().Focused().GetEntities() {
		s.Inputs().Focused().Remove(focusedEntity)
	}
}
