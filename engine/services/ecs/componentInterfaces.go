package ecs

type ApplyEntityEvent interface {
	ApplyEntity(entityEmitting EntityID) (event any)
}
