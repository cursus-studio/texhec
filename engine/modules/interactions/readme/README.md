# interactions
## Architecture
Allows to compose features (events) from multiple steps.
It is heavily inspired by wizzard pattern.

Feature is event emited after collecting multiple steps.
Step is filtered interaction. For example it isn't only unit click it is friendly unit click.
Interaction is selecting an single thing like object or coordinates.

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               8            131             62            672
Markdown                         2              1              0              6
-------------------------------------------------------------------------------
SUM:                            10            132             62            678
-------------------------------------------------------------------------------
```
## TODO
Add feature and interaction history

## Types
### type Service
Type: `engine/modules/interactions.Service`
service

#### method Service AvailableFeatures
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/interactions.AvailableFeaturesComponent]`

#### method Service Features
Type: `func() []engine/modules/interactions.FeatureKey`

### type FeatureKey
Type: `engine/modules/interactions.FeatureKey`

#### method FeatureKey Align
Type: `func() int`

#### method FeatureKey AssignableTo
Type: `func(u reflect.Type) bool`

#### method FeatureKey Bits
Type: `func() int`

#### method FeatureKey CanSeq
Type: `func() bool`

#### method FeatureKey CanSeq2
Type: `func() bool`

#### method FeatureKey ChanDir
Type: `func() reflect.ChanDir`

#### method FeatureKey Comparable
Type: `func() bool`

#### method FeatureKey ConvertibleTo
Type: `func(u reflect.Type) bool`

#### method FeatureKey Elem
Type: `func() reflect.Type`

#### method FeatureKey Field
Type: `func(i int) reflect.StructField`

#### method FeatureKey FieldAlign
Type: `func() int`

#### method FeatureKey FieldByIndex
Type: `func(index []int) reflect.StructField`

#### method FeatureKey FieldByName
Type: `func(name string) (reflect.StructField, bool)`

#### method FeatureKey FieldByNameFunc
Type: `func(match func(string) bool) (reflect.StructField, bool)`

#### method FeatureKey Implements
Type: `func(u reflect.Type) bool`

#### method FeatureKey In
Type: `func(i int) reflect.Type`

#### method FeatureKey IsVariadic
Type: `func() bool`

#### method FeatureKey Key
Type: `func() reflect.Type`

#### method FeatureKey Kind
Type: `func() reflect.Kind`

#### method FeatureKey Len
Type: `func() int`

#### method FeatureKey Method
Type: `func(int) reflect.Method`

#### method FeatureKey MethodByName
Type: `func(string) (reflect.Method, bool)`

#### method FeatureKey Name
Type: `func() string`

#### method FeatureKey NumField
Type: `func() int`

#### method FeatureKey NumIn
Type: `func() int`

#### method FeatureKey NumMethod
Type: `func() int`

#### method FeatureKey NumOut
Type: `func() int`

#### method FeatureKey Out
Type: `func(i int) reflect.Type`

#### method FeatureKey OverflowComplex
Type: `func(x complex128) bool`

#### method FeatureKey OverflowFloat
Type: `func(x float64) bool`

#### method FeatureKey OverflowInt
Type: `func(x int64) bool`

#### method FeatureKey OverflowUint
Type: `func(x uint64) bool`

#### method FeatureKey PkgPath
Type: `func() string`

#### method FeatureKey Size
Type: `func() uintptr`

#### method FeatureKey String
Type: `func() string`

### type StepKey
Type: `engine/modules/interactions.StepKey`

#### method StepKey Align
Type: `func() int`

#### method StepKey AssignableTo
Type: `func(u reflect.Type) bool`

#### method StepKey Bits
Type: `func() int`

#### method StepKey CanSeq
Type: `func() bool`

#### method StepKey CanSeq2
Type: `func() bool`

#### method StepKey ChanDir
Type: `func() reflect.ChanDir`

#### method StepKey Comparable
Type: `func() bool`

#### method StepKey ConvertibleTo
Type: `func(u reflect.Type) bool`

#### method StepKey Elem
Type: `func() reflect.Type`

#### method StepKey Field
Type: `func(i int) reflect.StructField`

#### method StepKey FieldAlign
Type: `func() int`

#### method StepKey FieldByIndex
Type: `func(index []int) reflect.StructField`

#### method StepKey FieldByName
Type: `func(name string) (reflect.StructField, bool)`

#### method StepKey FieldByNameFunc
Type: `func(match func(string) bool) (reflect.StructField, bool)`

#### method StepKey Implements
Type: `func(u reflect.Type) bool`

#### method StepKey In
Type: `func(i int) reflect.Type`

#### method StepKey IsVariadic
Type: `func() bool`

#### method StepKey Key
Type: `func() reflect.Type`

#### method StepKey Kind
Type: `func() reflect.Kind`

#### method StepKey Len
Type: `func() int`

#### method StepKey Method
Type: `func(int) reflect.Method`

#### method StepKey MethodByName
Type: `func(string) (reflect.Method, bool)`

#### method StepKey Name
Type: `func() string`

#### method StepKey NumField
Type: `func() int`

#### method StepKey NumIn
Type: `func() int`

#### method StepKey NumMethod
Type: `func() int`

#### method StepKey NumOut
Type: `func() int`

#### method StepKey Out
Type: `func(i int) reflect.Type`

#### method StepKey OverflowComplex
Type: `func(x complex128) bool`

#### method StepKey OverflowFloat
Type: `func(x float64) bool`

#### method StepKey OverflowInt
Type: `func(x int64) bool`

#### method StepKey OverflowUint
Type: `func(x uint64) bool`

#### method StepKey PkgPath
Type: `func() string`

#### method StepKey Size
Type: `func() uintptr`

#### method StepKey String
Type: `func() string`

### type InteractionKey
Type: `engine/modules/interactions.InteractionKey`

#### method InteractionKey Align
Type: `func() int`

#### method InteractionKey AssignableTo
Type: `func(u reflect.Type) bool`

#### method InteractionKey Bits
Type: `func() int`

#### method InteractionKey CanSeq
Type: `func() bool`

#### method InteractionKey CanSeq2
Type: `func() bool`

#### method InteractionKey ChanDir
Type: `func() reflect.ChanDir`

#### method InteractionKey Comparable
Type: `func() bool`

#### method InteractionKey ConvertibleTo
Type: `func(u reflect.Type) bool`

#### method InteractionKey Elem
Type: `func() reflect.Type`

#### method InteractionKey Field
Type: `func(i int) reflect.StructField`

#### method InteractionKey FieldAlign
Type: `func() int`

#### method InteractionKey FieldByIndex
Type: `func(index []int) reflect.StructField`

#### method InteractionKey FieldByName
Type: `func(name string) (reflect.StructField, bool)`

#### method InteractionKey FieldByNameFunc
Type: `func(match func(string) bool) (reflect.StructField, bool)`

#### method InteractionKey Implements
Type: `func(u reflect.Type) bool`

#### method InteractionKey In
Type: `func(i int) reflect.Type`

#### method InteractionKey IsVariadic
Type: `func() bool`

#### method InteractionKey Key
Type: `func() reflect.Type`

#### method InteractionKey Kind
Type: `func() reflect.Kind`

#### method InteractionKey Len
Type: `func() int`

#### method InteractionKey Method
Type: `func(int) reflect.Method`

#### method InteractionKey MethodByName
Type: `func(string) (reflect.Method, bool)`

#### method InteractionKey Name
Type: `func() string`

#### method InteractionKey NumField
Type: `func() int`

#### method InteractionKey NumIn
Type: `func() int`

#### method InteractionKey NumMethod
Type: `func() int`

#### method InteractionKey NumOut
Type: `func() int`

#### method InteractionKey Out
Type: `func(i int) reflect.Type`

#### method InteractionKey OverflowComplex
Type: `func(x complex128) bool`

#### method InteractionKey OverflowFloat
Type: `func(x float64) bool`

#### method InteractionKey OverflowInt
Type: `func(x int64) bool`

#### method InteractionKey OverflowUint
Type: `func(x uint64) bool`

#### method InteractionKey PkgPath
Type: `func() string`

#### method InteractionKey Size
Type: `func() uintptr`

#### method InteractionKey String
Type: `func() string`

### type InteractionService
Type: `engine/modules/interactions.InteractionService[State any]`

#### method InteractionService MissingPreview
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/interactions.MissingPreviewComponent[State]]`

#### method InteractionService Save
Type: `func(propertiesEntity engine/modules/ecs.EntityID, state State)`
saves state in entity with [IsMissingComponent] or resets interactions

#### method InteractionService StatePreview
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/interactions.StatePreviewComponent[State]]`

### type Step
Type: `engine/modules/interactions.Step[State any]`
step

#### method Step State
Type: `func() State`

### type Event
Type: `engine/modules/interactions.Event`
feature

### type StatePreviewComponent
Type: `engine/modules/interactions.StatePreviewComponent[State any]`
interaction

#### property StatePreviewComponent State
Type: `State`

### type MissingPreviewComponent
Type: `engine/modules/interactions.MissingPreviewComponent[State any]`

### type AvailableFeaturesComponent
Type: `engine/modules/interactions.AvailableFeaturesComponent`

#### property AvailableFeaturesComponent Features
Type: `[]engine/modules/interactions.FeatureKey`

#### property AvailableFeaturesComponent Selected
Type: `bool`

#### method AvailableFeaturesComponent Equal
Type: `func(other engine/modules/interactions.AvailableFeaturesComponent) bool`

### type SelectFeatureEvent
Type: `engine/modules/interactions.SelectFeatureEvent`

#### property SelectFeatureEvent FeatureKey
Type: `engine/modules/interactions.FeatureKey`

## Functions
### func NewStatePreview
Type: `func[State any](state State) engine/modules/interactions.StatePreviewComponent[State]`

### func NewMissingPreview
Type: `func[State any]() engine/modules/interactions.MissingPreviewComponent[State]`

### func NewStep
Type: `func[State any](state State) engine/modules/interactions.Step[State]`

### func NewSelectedFeature
Type: `func(feature engine/modules/interactions.FeatureKey) engine/modules/interactions.AvailableFeaturesComponent`

### func NewAvailableFeatures
Type: `func(features ...engine/modules/interactions.FeatureKey) engine/modules/interactions.AvailableFeaturesComponent`

### func NewSelectFeatureEvent
Type: `func(featureKey engine/modules/interactions.FeatureKey) engine/modules/interactions.SelectFeatureEvent`

### func NewDeselectFeatureEvent
Type: `func() engine/modules/interactions.SelectFeatureEvent`


## Dependencies
`engine`:
  - `engine.EngineWorld`
  - `engine.Events`
  - `engine.EventsBuilder`
  - `engine.Hierarchy`
  - `engine.Logger`
  - `engine.Prototype`
  - `engine.World`

`engine/modules/ecs`:
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`
  - `engine/modules/ecs.World`

`engine/modules/grid`:
  - `engine/modules/grid.Coords`

`engine/modules/interactions`:
  - `engine/modules/interactions.AvailableFeaturesComponent`
  - `engine/modules/interactions.FeatureKey`
  - `engine/modules/interactions.Features`
  - `engine/modules/interactions.InteractionKey`
  - `engine/modules/interactions.InteractionService`
  - `engine/modules/interactions.MissingPreviewComponent`
  - `engine/modules/interactions.NewAvailableFeatures`
  - `engine/modules/interactions.NewMissingPreview`
  - `engine/modules/interactions.NewSelectedFeature`
  - `engine/modules/interactions.NewStatePreview`
  - `engine/modules/interactions.NewStep`
  - `engine/modules/interactions.SelectFeatureEvent`
  - `engine/modules/interactions.Selected`
  - `engine/modules/interactions.Service`
  - `engine/modules/interactions.StatePreviewComponent`
  - `engine/modules/interactions.Step`
  - `engine/modules/interactions.StepKey`

`engine/modules/interactions/pkg`:
  - `engine/modules/interactions/pkg.FeaturePkg`
  - `engine/modules/interactions/pkg.InteractionPkg`
  - `engine/modules/interactions/pkg.NewRelation`
  - `engine/modules/interactions/pkg.StepPkg`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`