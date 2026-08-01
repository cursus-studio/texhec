package internal

import (
	"engine"
	"engine/modules/delay"
	"engine/modules/ecs"
	"engine/modules/loop"
	"engine/modules/record"
	"engine/modules/smooth"
	"engine/modules/transition"
	"time"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type UpdateWaypointEvent[Component any] struct {
	Entity ecs.EntityID
	Delay  time.Duration
}

func NewUpdateWaypointEvent[Component any](entity ecs.EntityID) UpdateWaypointEvent[Component] {
	return UpdateWaypointEvent[Component]{Entity: entity}
}

func (event UpdateWaypointEvent[Component]) ApplyDelay(delay time.Duration) any {
	event.Delay = delay
	return event
}

//

type NextWaypointsComponent[Component any] struct {
	Waypoints []Component
}
type WaypointsComponent[Component any] struct {
	CurrentWaypoint int
	WaypointDelta   time.Duration
	Waypoints       []Component
}

func (c *NextWaypointsComponent[Component]) Next(tickDelta time.Duration,
	first, last Component) WaypointsComponent[Component] {
	waypointDelta := tickDelta / time.Duration(len(c.Waypoints)+1)
	return WaypointsComponent[Component]{
		CurrentWaypoint: 0,
		WaypointDelta:   waypointDelta,
		Waypoints:       append([]Component{first}, append(c.Waypoints, last)...),
	}
}
func (c *WaypointsComponent[Component]) Progress() (
	state Component, lerp transition.TransitionComponent[Component], hasWaypointLeft bool) {
	state = c.Waypoints[c.CurrentWaypoint]
	lerp = transition.NewTransition(
		c.Waypoints[c.CurrentWaypoint],
		c.Waypoints[c.CurrentWaypoint+1],
		c.WaypointDelta)
	c.CurrentWaypoint++
	hasWaypointLeft = len(c.Waypoints) > c.CurrentWaypoint+1
	return state, lerp, hasWaypointLeft
}

// type ServiceT[Component smooth.SmoothConstraint[Component]] struct {
type ServiceT[Component any] struct {
	engine.EngineWorld `inject:""`
	recordingID        record.RecordingID
	config             record.Config

	componentArray     ecs.ComponentArray[Component]
	nextWaypointsArray ecs.ComponentArray[NextWaypointsComponent[Component]]
	waypointsArray     ecs.ComponentArray[WaypointsComponent[Component]]
	lerpArray          ecs.ComponentArray[transition.TransitionComponent[Component]]
}

// func NewServiceT[Component smooth.SmoothConstraint[Component]](c ioc.Dic) *Service[Component] {
func NewServiceT[Component any](c ioc.Dic) *ServiceT[Component] {
	config := record.NewConfig()
	record.AddToConfig[Component](config)

	s := ioc.GetServices[*ServiceT[Component]](c)

	s.recordingID = 0
	s.config = config
	s.componentArray = ecs.GetComponentArray[Component](s.World())
	s.nextWaypointsArray = ecs.GetComponentArray[NextWaypointsComponent[Component]](s.World())
	s.waypointsArray = ecs.GetComponentArray[WaypointsComponent[Component]](s.World())
	s.lerpArray = ecs.GetComponentArray[transition.TransitionComponent[Component]](s.World())
	events.Listen(s.EventsBuilder(), s.AddWaypointListener)
	return s
}

func (s *ServiceT[Component]) AddWaypointListener(e smooth.AddWaypointEvent[Component]) {
	s.AddWaypoint(e.Entity, e.State)
}
func (s *ServiceT[Component]) AddWaypoint(entity ecs.EntityID, component Component) {
	waypoints, _ := s.nextWaypointsArray.Get(entity)
	waypoints.Waypoints = append(waypoints.Waypoints, component)
	s.nextWaypointsArray.Set(entity, waypoints)
}
func (s *ServiceT[Component]) SaveWaypoint(entity ecs.EntityID, tickDelta time.Duration, first, last Component) {
	nextWaypoints, _ := s.nextWaypointsArray.Get(entity)
	waypoints := nextWaypoints.Next(tickDelta, first, last)
	s.nextWaypointsArray.Remove(entity)
	s.waypointsArray.Set(entity, waypoints)
}

//

// type system[Component smooth.SmoothConstraint[Component]] struct {
type system[Component any] struct {
	engine.EngineWorld `inject:""`
	Service            *ServiceT[Component] `inject:""`
}

type FirstEvent loop.TickEvent
type LastEvent loop.TickEvent

func (s *system[Component]) First(event FirstEvent) {
	for _, entity := range s.Service.waypointsArray.GetEntities() {
		waypoint, ok := s.Service.waypointsArray.Get(entity)
		if !ok {
			continue
		}
		s.Service.lerpArray.Remove(entity)
		s.Service.waypointsArray.Remove(entity)
		s.Service.componentArray.Set(entity, waypoint.Waypoints[len(waypoint.Waypoints)-1])
	}

	s.Service.recordingID = s.Record().Entity().StartBackwardsRecording(s.Service.config)
}
func (s *system[Component]) Last(event LastEvent) {
	r, ok := s.Record().Entity().Stop(s.Service.recordingID)
	if !ok {
		return
	}
	for _, entity := range r.Entities.GetIndices() {
		beforeComponents, ok := r.Entities.Get(entity)
		if !ok || beforeComponents == nil {
			continue
		}
		before, ok := beforeComponents[0].(Component)
		if !ok {
			continue
		}
		after, ok := s.Service.componentArray.Get(entity)
		if !ok {
			continue
		}
		s.Service.SaveWaypoint(entity, event.Delta, before, after)
		s.ListenUpdateWaypoint(NewUpdateWaypointEvent[Component](entity))
	}
}
func (s *system[Component]) ListenUpdateWaypoint(event UpdateWaypointEvent[Component]) {
	waypoint, ok := s.Service.waypointsArray.Get(event.Entity)
	if !ok {
		return
	}
	state, lerpComponent, hasWaypointsLeft := waypoint.Progress()
	s.Service.waypointsArray.Set(event.Entity, waypoint)
	s.Service.lerpArray.Set(event.Entity, lerpComponent)
	s.Service.componentArray.Set(event.Entity, state)
	if !hasWaypointsLeft {
		return
	}

	updateWaypointEvent := NewUpdateWaypointEvent[Component](event.Entity)
	delayEvent := delay.NewDelayedEvent(
		updateWaypointEvent,
		waypoint.WaypointDelta-event.Delay,
	)
	s.Delay().Delay(delayEvent)
}

// func NewSystems[Component smooth.SmoothConstraint[Component]](c ioc.Dic) {
func NewSystems[Component any](c ioc.Dic) {
	s := ioc.GetServices[*system[Component]](c)
	events.Listen(s.EventsBuilder(), s.First)
	events.Listen(s.EventsBuilder(), s.Last)
	events.Listen(s.EventsBuilder(), s.ListenUpdateWaypoint)
}
