package internal

import (
	"engine"
	"engine/modules/batcher"
	"engine/modules/loop"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type Service struct {
	engine.EngineWorld `inject:""`

	workers int
	tasks   []batcher.Task
}

func NewService(
	c ioc.Dic,
	workers int,
) *Service {
	s := ioc.GetServices[*Service](c)
	s.workers = workers
	return s
}

func (s *Service) NewTask() batcher.TaskFactory {
	f := NewTaskFactory(s.workers)
	f.AddOrderedBatch(batcher.NewBatch(1, func(int) { s.WarmUp().WarmUp() }))
	return f
}
func (s *Service) Queue(task batcher.Task) { s.tasks = append(s.tasks, task) }
func (s *Service) Progress() float32 {
	if len(s.tasks) != 0 {
		return s.tasks[0].Progress()
	}
	return -1
}

func (s *Service) Register() error {
	events.Listen(s.EventsBuilder(), s.Listen)
	return nil
}

func (s *Service) Listen(loop.FrameEvent) {
	if len(s.tasks) == 0 {
		return
	}
	task := s.tasks[0]
	for s.Loop().Stats().FrameBudgetLeft() > 0 {
		task.Step()
		if task.Progress() != 1 {
			continue
		}
		s.tasks = s.tasks[1:]
		if len(s.tasks) == 0 {
			break
		}
		task = s.tasks[0]
	}
}
