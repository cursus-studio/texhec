# interactions
## Architecture
Allows to compose features (events) from multiple user interactions

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               5             53             15            285
Markdown                         2              0              0              2
-------------------------------------------------------------------------------
SUM:                             7             53             15            287
-------------------------------------------------------------------------------
```
## TODO
Add feature and interaction history

## Types
### type Service
Type: `engine/modules/interactions.Service`

#### method Service FeatureEntity
Type: `func() engine/modules/ecs.EntityID`

#### method Service FeatureEvent
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/interactions.FeatureEventComponent]`

#### method Service Instance
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/interactions.InstanceComponent]`
entity with this stores all components with interactions and selected feature

#### method Service Proceed
Type: `func(engine/modules/ecs.EntityID)`
entity here isn't used
it's here just to make calling it easier by OnUpsert because only there it should be used

### type Event
Type: `engine/modules/interactions.Event`

### type AnyInteractionService
Type: `engine/modules/interactions.AnyInteractionService`

#### method AnyInteractionService InteractionAny
Type: `func() engine/modules/ecs.AnyComponentArray`

#### method AnyInteractionService Measure
Type: `func() (alreadyMeasured bool)`
saves [MissingInteractionComponent] if it there is no [InteractionComponent]

#### method AnyInteractionService MissingInteractionAny
Type: `func() engine/modules/ecs.AnyComponentArray`

#### method AnyInteractionService Name
Type: `func() engine/modules/interactions.Name`

#### method AnyInteractionService StateType
Type: `func() reflect.Type`

### type InteractionService
Type: `engine/modules/interactions.InteractionService[State any]`

#### method InteractionService FinishMeasurement
Type: `func(engine/modules/interactions.FinishMeasurementEvent[State])`

#### method InteractionService Interaction
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/interactions.InteractionComponent[State]]`

#### method InteractionService InteractionAny
Type: `func() engine/modules/ecs.AnyComponentArray`

#### method InteractionService InteractionGUI
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/interactions.InteractionGUIComponent[State]]`
elements are removed when interaction is removed.
they can be used to indicate that element is used.

#### method InteractionService Measure
Type: `func() (alreadyMeasured bool)`
saves [MissingInteractionComponent] if it there is no [InteractionComponent]

#### method InteractionService MissingInteraction
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/interactions.MissingInteractionComponent[State]]`

#### method InteractionService MissingInteractionAny
Type: `func() engine/modules/ecs.AnyComponentArray`

#### method InteractionService Name
Type: `func() engine/modules/interactions.Name`

#### method InteractionService StateType
Type: `func() reflect.Type`

### type AnyFeatureService
Type: `engine/modules/interactions.AnyFeatureService`

#### method AnyFeatureService EventType
Type: `func() reflect.Type`

#### method AnyFeatureService Interactions
Type: `func() []engine/modules/interactions.AnyInteractionService`

#### method AnyFeatureService Name
Type: `func() engine/modules/interactions.Name`

### type FeatureService
Type: `engine/modules/interactions.FeatureService[Event any]`
listens to [FeatureEvent]

#### method FeatureService EventType
Type: `func() reflect.Type`

#### method FeatureService Interactions
Type: `func() []engine/modules/interactions.AnyInteractionService`

#### method FeatureService Name
Type: `func() engine/modules/interactions.Name`

### type MissingInteractionComponent
Type: `engine/modules/interactions.MissingInteractionComponent[State any]`

### type InteractionGUIComponent
Type: `engine/modules/interactions.InteractionGUIComponent[State any]`

### type InteractionComponent
Type: `engine/modules/interactions.InteractionComponent[State any]`

#### property InteractionComponent State
Type: `State`

### type FinishMeasurementEvent
Type: `engine/modules/interactions.FinishMeasurementEvent[State any]`

#### property FinishMeasurementEvent State
Type: `State`

### type FeatureEvent
Type: `engine/modules/interactions.FeatureEvent[Event any]`

#### property FeatureEvent Event
Type: `Event`

#### method FeatureEvent Component
Type: `func() engine/modules/interactions.FeatureEventComponent`

### type FeatureEventComponent
Type: `engine/modules/interactions.FeatureEventComponent`

#### property FeatureEventComponent Event
Type: `engine/modules/interactions.Event`

### type InstanceComponent
Type: `engine/modules/interactions.InstanceComponent`

## Functions
### func NewInteraction
Type: `func[State any](state State) engine/modules/interactions.InteractionComponent[State]`

### func NewFinishMeasurementEvent
Type: `func[State any](state State) engine/modules/interactions.FinishMeasurementEvent[State]`

### func NewFeatureEvent
Type: `func[Event any](event Event) engine/modules/interactions.FeatureEvent[Event]`


## Dependencies
`engine`:
  - `engine.EngineWorld`
  - `engine.Events`
  - `engine.EventsBuilder`
  - `engine.Interactions`
  - `engine.Logger`
  - `engine.World`

`engine/modules/ecs`:
  - `engine/modules/ecs.AnyComponentArray`
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`

`engine/modules/interactions`:
  - `engine/modules/interactions.AnyFeatureService`
  - `engine/modules/interactions.AnyInteractionService`
  - `engine/modules/interactions.Component`
  - `engine/modules/interactions.Event`
  - `engine/modules/interactions.EventType`
  - `engine/modules/interactions.FeatureEntity`
  - `engine/modules/interactions.FeatureEvent`
  - `engine/modules/interactions.FeatureEventComponent`
  - `engine/modules/interactions.FeatureService`
  - `engine/modules/interactions.FinishMeasurementEvent`
  - `engine/modules/interactions.InstanceComponent`
  - `engine/modules/interactions.InteractionAny`
  - `engine/modules/interactions.InteractionComponent`
  - `engine/modules/interactions.InteractionGUIComponent`
  - `engine/modules/interactions.InteractionService`
  - `engine/modules/interactions.Interactions`
  - `engine/modules/interactions.MissingInteractionAny`
  - `engine/modules/interactions.MissingInteractionComponent`
  - `engine/modules/interactions.Name`
  - `engine/modules/interactions.NewInteraction`
  - `engine/modules/interactions.Proceed`
  - `engine/modules/interactions.Service`
  - `engine/modules/interactions.State`
  - `engine/modules/interactions.StateType`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`