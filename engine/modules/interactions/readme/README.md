# interactions
## Architecture
Allows to compose events from multiple user interactions with GUI

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               5             50             15            258
Markdown                         2              0              0              2
-------------------------------------------------------------------------------
SUM:                             7             50             15            260
-------------------------------------------------------------------------------
```
## TODO
Add feature and interaction history

## Types
### type Service
Type: `engine/modules/interactions.Service`

#### method Service Feature
Type: `func() engine/services/ecs.ComponentsArray[engine/modules/interactions.FeatureComponent]`
entity with this stores all components with interactions and selected feature

#### method Service FeatureEntity
Type: `func() engine/services/ecs.EntityID`

#### method Service FeatureEvent
Type: `func() engine/services/ecs.ComponentsArray[engine/modules/interactions.FeatureEventComponent]`

#### method Service Proceed
Type: `func(engine/services/ecs.EntityID)`
entity here isn't used
it's here just to make calling it easier by OnUpsert because only there it should be used

### type Event
Type: `engine/modules/interactions.Event`

### type AnyInteractionService
Type: `engine/modules/interactions.AnyInteractionService`

#### method AnyInteractionService InteractionAny
Type: `func() engine/services/ecs.AnyComponentArray`

#### method AnyInteractionService Measure
Type: `func() (alreadyMeasured bool)`
saves [MissingInteractionComponent] if it there is no [InteractionComponent]

#### method AnyInteractionService MissingInteractionAny
Type: `func() engine/services/ecs.AnyComponentArray`

#### method AnyInteractionService Name
Type: `func() engine/modules/interactions.Name`

#### method AnyInteractionService StateType
Type: `func() reflect.Type`

### type InteractionService
Type: `engine/modules/interactions.InteractionService[State any]`

#### method InteractionService FinishMeasurement
Type: `func(State) error`

#### method InteractionService Interaction
Type: `func() engine/services/ecs.ComponentsArray[engine/modules/interactions.InteractionComponent[State]]`

#### method InteractionService InteractionAny
Type: `func() engine/services/ecs.AnyComponentArray`

#### method InteractionService Measure
Type: `func() (alreadyMeasured bool)`
saves [MissingInteractionComponent] if it there is no [InteractionComponent]

#### method InteractionService MissingInteraction
Type: `func() engine/services/ecs.ComponentsArray[engine/modules/interactions.MissingInteractionComponent[State]]`

#### method InteractionService MissingInteractionAny
Type: `func() engine/services/ecs.AnyComponentArray`

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

### type InteractionComponent
Type: `engine/modules/interactions.InteractionComponent[State any]`

#### property InteractionComponent State
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

### type FeatureComponent
Type: `engine/modules/interactions.FeatureComponent`

## Functions
### func NewInteraction
Type: `func[State any](state State) engine/modules/interactions.InteractionComponent[State]`

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

`engine/modules/interactions`:
  - `engine/modules/interactions.AnyFeatureService`
  - `engine/modules/interactions.AnyInteractionService`
  - `engine/modules/interactions.Event`
  - `engine/modules/interactions.EventType`
  - `engine/modules/interactions.Feature`
  - `engine/modules/interactions.FeatureComponent`
  - `engine/modules/interactions.FeatureEntity`
  - `engine/modules/interactions.FeatureEvent`
  - `engine/modules/interactions.FeatureEventComponent`
  - `engine/modules/interactions.FeatureService`
  - `engine/modules/interactions.InteractionAny`
  - `engine/modules/interactions.InteractionComponent`
  - `engine/modules/interactions.InteractionService`
  - `engine/modules/interactions.Interactions`
  - `engine/modules/interactions.MissingInteractionAny`
  - `engine/modules/interactions.MissingInteractionComponent`
  - `engine/modules/interactions.Name`
  - `engine/modules/interactions.NewInteraction`
  - `engine/modules/interactions.Proceed`
  - `engine/modules/interactions.Service`
  - `engine/modules/interactions.StateType`

`engine/services/ecs`:
  - `engine/services/ecs.AnyComponentArray`
  - `engine/services/ecs.ComponentsArray`
  - `engine/services/ecs.EntityID`
  - `engine/services/ecs.Get`
  - `engine/services/ecs.GetAny`
  - `engine/services/ecs.GetComponentsArray`
  - `engine/services/ecs.GetEmpty`
  - `engine/services/ecs.GetEntities`
  - `engine/services/ecs.NewEntity`
  - `engine/services/ecs.OnUpsert`
  - `engine/services/ecs.Remove`
  - `engine/services/ecs.RemoveEntity`
  - `engine/services/ecs.Set`
  - `engine/services/ecs.SetAny`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`