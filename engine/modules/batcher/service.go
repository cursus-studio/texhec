// this module allows us to write tasks and to progress them across frames without stuterring
package batcher

import "engine/modules/ecs"

type Batch struct {
	Steps   int
	Handler func(int)
}

func NewBatch(steps int, handler func(int)) Batch {
	return Batch{steps, handler}
}

type TaskFactory interface {
	AddOrderedBatch(Batch) TaskFactory
	AddConcurrentBatch(Batch) TaskFactory
	Build() Task
}

type Task interface {
	Step()
	Progress() float32

	Perform()
}

type Service interface {
	ecs.SystemRegister
	NewTask() TaskFactory

	Queue(Task)
	Tasks() []Task
	// progress of first task in queue
	// when there is no tasks in queue -1 is returned
	Progress() float32
}
