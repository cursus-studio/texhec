package metadata

import "engine/services/ecs"

type NameComponent struct {
	Name string
}

func NewName(name string) NameComponent {
	return NameComponent{Name: name}
}

//

type DescriptionComponent struct {
	Description string
}

func NewDescription(description string) DescriptionComponent {
	return DescriptionComponent{Description: description}
}

//

type Service interface {
	Name() ecs.ComponentsArray[NameComponent]
	Description() ecs.ComponentsArray[DescriptionComponent]
}
