package prediction_test

import (
	"engine"
	"engine/modules/ecs"
	"engine/modules/netsync"
	"engine/modules/netsync/internal/prediction"
	netsyncpkg "engine/modules/netsync/pkg"
	"engine/modules/record"
	enginepkg "engine/pkg"
	"slices"
	"testing"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type ValueComponent struct {
	Value int
}
type IncreaseToNextEvenEvent struct{}
type IncreaseToNextOddEvent struct{}
type IncreaseByTenEvent struct{}

func NewValue(value int) ValueComponent { return ValueComponent{value} }

func (c ValueComponent) NextEven() ValueComponent {
	c.Value += 2 - (c.Value & 1)
	return c
}
func (c ValueComponent) NextOdd() ValueComponent {
	c.Value += 1 + (c.Value & 1)
	return c
}
func (c ValueComponent) IncrementByTen() ValueComponent {
	c.Value += 10
	return c
}

//

type World struct {
	engine.EngineWorld `inject:""`
	Prediction         prediction.Service   `inject:""`
	PredictionID       prediction.IDService `inject:""`
	T                  *testing.T

	componentArray ecs.ComponentArray[ValueComponent]
}

func (w World) AssertSuccess(err error) {
	w.T.Helper()
	if err == nil {
		return
	}
	w.T.Fatalf("unexpected error: %v", err)
}

type Component int
type Queued int
type Flushed []prediction.Event

func (c Component) NextEven() Component       { return c + 2 - (c & 1) }
func (c Component) NextOdd() Component        { return c + 1 + (c & 1) }
func (c Component) IncrementByTen() Component { return c + 10 }

func (w World) AssertWorldState(expectedValue Component, expectedQueued Queued, expectedFlushed Flushed) []prediction.ClientEventID {
	w.T.Helper()
	w.AssertState(NewValue(int(expectedValue)))
	w.AssertQueued(int(expectedQueued))
	return w.AssertFlushed(expectedFlushed...)
}
func (w World) AssertState(expected ValueComponent) {
	w.T.Helper()
	comp := w.GetComp()
	if comp != expected {
		w.T.Fatalf("unexpected state. expected %v and got %v", expected, comp)
	}
}
func (w World) AssertQueued(n int) {
	w.T.Helper()
	queued := w.Prediction.QueuedPredictions()
	if n != queued {
		w.T.Fatalf("queued unexpected predictions. queued %v and expected %v", queued, n)
	}
}
func (w World) AssertFlushed(events ...prediction.Event) []prediction.ClientEventID {
	w.T.Helper()
	newPredictions := w.Prediction.FlushNewPredictions()
	var flushedIDs []prediction.ClientEventID
	var flushedEvents []prediction.Event
	for _, prediction := range newPredictions {
		flushedIDs = append(flushedIDs, prediction.ID)
		flushedEvents = append(flushedEvents, prediction.Event)
	}

	if !slices.Equal(flushedEvents, events) {
		w.T.Fatalf("unexpected events got flushed. expected %v to be flushed but got %v", events, flushedEvents)
	}
	return flushedIDs
}

func (w World) GetEntity() ecs.EntityID {
	if len(w.componentArray.GetEntities()) == 1 {
		return w.componentArray.GetEntities()[0]
	}
	for _, entity := range slices.Clone(w.componentArray.GetEntities()) {
		w.componentArray.Remove(entity)
	}
	entity := w.World().NewEntity()
	value := NewValue(0)
	w.componentArray.Set(entity, value)
	return entity
}
func (w World) GetComp() ValueComponent {
	entity := w.GetEntity()
	comp, _ := w.componentArray.Get(entity)
	return comp
}

var Pkg = ioc.NewPkgT(func(b ioc.Builder, t *testing.T) {
	ioc.Wrap(b, func(c ioc.Dic, config netsyncpkg.Config) {
		record.AddToConfig[ValueComponent](config.RecordConfig())
		netsyncpkg.AddEvent[IncreaseToNextEvenEvent](config)
		netsyncpkg.AddEvent[IncreaseToNextOddEvent](config)
		netsyncpkg.AddTickEvent[IncreaseByTenEvent](config)
	})
	ioc.Register(b, func(c ioc.Dic) World {
		w := ioc.GetServices[World](c)
		w.T = t
		w.componentArray = ecs.GetComponentArray[ValueComponent](w.World())

		// enables predictions
		w.NetSync().Server().Set(w.World().NewEntity(), netsync.ServerComponent{})

		w.AssertSuccess(w.NetSync().Start().Register())
		events.Listen(w.EventsBuilder(), func(IncreaseToNextEvenEvent) {
			entity := w.GetEntity()
			comp, _ := w.componentArray.Get(entity)
			w.componentArray.Set(entity, comp.NextEven())
		})
		events.Listen(w.EventsBuilder(), func(IncreaseToNextOddEvent) {
			entity := w.GetEntity()
			comp, _ := w.componentArray.Get(entity)
			w.componentArray.Set(entity, comp.NextOdd())
		})
		events.Listen(w.EventsBuilder(), func(IncreaseByTenEvent) {
			entity := w.GetEntity()
			comp, _ := w.componentArray.Get(entity)
			w.componentArray.Set(entity, comp.IncrementByTen())
		})
		w.AssertSuccess(w.NetSync().Stop().Register())
		return w
	})
})

//

func NewWorld(t *testing.T) World {
	c := ioc.NewContainer(enginepkg.Pkg, Pkg(t))
	return ioc.Get[World](c)
}
