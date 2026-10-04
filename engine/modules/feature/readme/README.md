# feature
## Architecture
This is second attempt to create `feature` architecture, previous was in module called `interactions`.

It splits `feature` into:
- ### Fields
`FeatureEvent` structs are composed from fields and each can be either:
#### Provided
Field provided by code
#### Interactions
Field provided by interaction with GUI

- ### Rules
Rule wires two fields and runs test on them.
Rule can be anything (for e.g. `OwnsRule`, `CanAttackRule`).
It is sufficient to run all tests on feature.

- ### Features
Features are just an event composed from fields and rules between them.
There is a component which stores all compatible features with correct interactions.
There are few rules for it:
1. If no feature is selected it stores utmost one interaction and lists all features which first interaction matches
2. If feature is selected it'll ensure state matches and no interactions which fail tests are selected

### GUI Components
1. `InteractionsComponent`
It is global and stores whole state.
All following components will store this state or its consequences.
It stores:
- selected feature or feature isn't selected then which features are compatible
- preview interaction
- interactions selected

2. `FieldRenderComponent[Field]`
Specialized component for rendering fields which are:
- missing for selected feature to be complete
- preview field which do or do not pass rule
- selected interactions

3. `RuleRenderComponent[Rule]`
Specialized component for rendering rules for preview field

## Benchmarks
```
$ go test ./... -bench=.
PASS
ok  	engine/modules/feature/test	0.014s
```
## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                              11            171             96           1011
Markdown                         1              6              0             33
-------------------------------------------------------------------------------
SUM:                            12            177             96           1044
-------------------------------------------------------------------------------
```
## Types
### type Service
Type: `engine/modules/feature.Service`

#### method Service Features
Type: `func() []engine/modules/feature.FeatureKey`

#### method Service Interactions
Type: `func() engine/modules/feature.InteractionsComponent`

#### method Service OnInteractionsMod
Type: `func(func(engine/modules/ecs.EntityID))`

#### method Service PreviewInteraction
Type: `func(field any) error`

#### method Service Register
Type: `func() error`

#### method Service SubmitPreviewInteraction
Type: `func(fieldType engine/modules/feature.FieldKey)`

### type FieldKey
Type: `engine/modules/feature.FieldKey`

#### method FieldKey Align
Type: `func() int`

#### method FieldKey AssignableTo
Type: `func(u reflect.Type) bool`

#### method FieldKey Bits
Type: `func() int`

#### method FieldKey CanSeq
Type: `func() bool`

#### method FieldKey CanSeq2
Type: `func() bool`

#### method FieldKey ChanDir
Type: `func() reflect.ChanDir`

#### method FieldKey Comparable
Type: `func() bool`

#### method FieldKey ConvertibleTo
Type: `func(u reflect.Type) bool`

#### method FieldKey Elem
Type: `func() reflect.Type`

#### method FieldKey Field
Type: `func(i int) reflect.StructField`

#### method FieldKey FieldAlign
Type: `func() int`

#### method FieldKey FieldByIndex
Type: `func(index []int) reflect.StructField`

#### method FieldKey FieldByName
Type: `func(name string) (reflect.StructField, bool)`

#### method FieldKey FieldByNameFunc
Type: `func(match func(string) bool) (reflect.StructField, bool)`

#### method FieldKey Implements
Type: `func(u reflect.Type) bool`

#### method FieldKey In
Type: `func(i int) reflect.Type`

#### method FieldKey IsVariadic
Type: `func() bool`

#### method FieldKey Key
Type: `func() reflect.Type`

#### method FieldKey Kind
Type: `func() reflect.Kind`

#### method FieldKey Len
Type: `func() int`

#### method FieldKey Method
Type: `func(int) reflect.Method`

#### method FieldKey MethodByName
Type: `func(string) (reflect.Method, bool)`

#### method FieldKey Name
Type: `func() string`

#### method FieldKey NumField
Type: `func() int`

#### method FieldKey NumIn
Type: `func() int`

#### method FieldKey NumMethod
Type: `func() int`

#### method FieldKey NumOut
Type: `func() int`

#### method FieldKey Out
Type: `func(i int) reflect.Type`

#### method FieldKey OverflowComplex
Type: `func(x complex128) bool`

#### method FieldKey OverflowFloat
Type: `func(x float64) bool`

#### method FieldKey OverflowInt
Type: `func(x int64) bool`

#### method FieldKey OverflowUint
Type: `func(x uint64) bool`

#### method FieldKey PkgPath
Type: `func() string`

#### method FieldKey Size
Type: `func() uintptr`

#### method FieldKey String
Type: `func() string`

### type RuleKey
Type: `engine/modules/feature.RuleKey`

#### method RuleKey Align
Type: `func() int`

#### method RuleKey AssignableTo
Type: `func(u reflect.Type) bool`

#### method RuleKey Bits
Type: `func() int`

#### method RuleKey CanSeq
Type: `func() bool`

#### method RuleKey CanSeq2
Type: `func() bool`

#### method RuleKey ChanDir
Type: `func() reflect.ChanDir`

#### method RuleKey Comparable
Type: `func() bool`

#### method RuleKey ConvertibleTo
Type: `func(u reflect.Type) bool`

#### method RuleKey Elem
Type: `func() reflect.Type`

#### method RuleKey Field
Type: `func(i int) reflect.StructField`

#### method RuleKey FieldAlign
Type: `func() int`

#### method RuleKey FieldByIndex
Type: `func(index []int) reflect.StructField`

#### method RuleKey FieldByName
Type: `func(name string) (reflect.StructField, bool)`

#### method RuleKey FieldByNameFunc
Type: `func(match func(string) bool) (reflect.StructField, bool)`

#### method RuleKey Implements
Type: `func(u reflect.Type) bool`

#### method RuleKey In
Type: `func(i int) reflect.Type`

#### method RuleKey IsVariadic
Type: `func() bool`

#### method RuleKey Key
Type: `func() reflect.Type`

#### method RuleKey Kind
Type: `func() reflect.Kind`

#### method RuleKey Len
Type: `func() int`

#### method RuleKey Method
Type: `func(int) reflect.Method`

#### method RuleKey MethodByName
Type: `func(string) (reflect.Method, bool)`

#### method RuleKey Name
Type: `func() string`

#### method RuleKey NumField
Type: `func() int`

#### method RuleKey NumIn
Type: `func() int`

#### method RuleKey NumMethod
Type: `func() int`

#### method RuleKey NumOut
Type: `func() int`

#### method RuleKey Out
Type: `func(i int) reflect.Type`

#### method RuleKey OverflowComplex
Type: `func(x complex128) bool`

#### method RuleKey OverflowFloat
Type: `func(x float64) bool`

#### method RuleKey OverflowInt
Type: `func(x int64) bool`

#### method RuleKey OverflowUint
Type: `func(x uint64) bool`

#### method RuleKey PkgPath
Type: `func() string`

#### method RuleKey Size
Type: `func() uintptr`

#### method RuleKey String
Type: `func() string`

### type FeatureKey
Type: `engine/modules/feature.FeatureKey`

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

### type AnyFieldService
Type: `engine/modules/feature.AnyFieldService`

#### method AnyFieldService NewAnyRenderComponent
Type: `func(value any, fieldRender engine/modules/feature.FieldRender) any`

#### method AnyFieldService RenderAny
Type: `func() engine/modules/ecs.AnyComponentArray`

### type FieldService
Type: `engine/modules/feature.FieldService[FieldType any]`

#### method FieldService NewAnyRenderComponent
Type: `func(value any, fieldRender engine/modules/feature.FieldRender) any`

#### method FieldService Render
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/feature.FieldRenderComponent[FieldType]]`

#### method FieldService RenderAny
Type: `func() engine/modules/ecs.AnyComponentArray`

### type Rule
Type: `engine/modules/feature.Rule[DataIn, AppliedTo any]`

### type AnyRuleService
Type: `engine/modules/feature.AnyRuleService`

#### method AnyRuleService AppliesToAny
Type: `func(dataIn any, appliedTo any) error`

#### method AnyRuleService NewAnyRenderComponent
Type: `func(in any, appliedTo any) any`

#### method AnyRuleService RenderAny
Type: `func() engine/modules/ecs.AnyComponentArray`

### type RuleService
Type: `engine/modules/feature.RuleService[RuleType engine/modules/feature.Rule[DataIn, AppliedTo], DataIn, AppliedTo any]`

#### method RuleService Applies
Type: `func(DataIn, AppliedTo) error`

#### method RuleService AppliesToAny
Type: `func(dataIn any, appliedTo any) error`

#### method RuleService NewAnyRenderComponent
Type: `func(in any, appliedTo any) any`

#### method RuleService Render
Type: `func() engine/modules/ecs.ComponentArray[engine/modules/feature.RuleRenderComponent[RuleType, DataIn, AppliedTo]]`

#### method RuleService RenderAny
Type: `func() engine/modules/ecs.AnyComponentArray`

### type InteractionsService
Type: `engine/modules/feature.InteractionsService`

#### method InteractionsService Interactions
Type: `func() engine/modules/feature.InteractionsComponent`

#### method InteractionsService OnInteractionsMod
Type: `func(func(engine/modules/ecs.EntityID))`

#### method InteractionsService PreviewInteraction
Type: `func(field any) error`

#### method InteractionsService SubmitPreviewInteraction
Type: `func(fieldType engine/modules/feature.FieldKey)`

### type FieldRender
Type: `engine/modules/feature.FieldRender`

### type ProvideFieldService
Type: `engine/modules/feature.ProvideFieldService[FieldType any]`

### type FieldRenderComponent
Type: `engine/modules/feature.FieldRenderComponent[FieldType any]`

#### property FieldRenderComponent Value
Type: `FieldType`

#### property FieldRenderComponent Render
Type: `engine/modules/feature.FieldRender`

### type RuleRenderComponent
Type: `engine/modules/feature.RuleRenderComponent[RuleType engine/modules/feature.Rule[DataIn, AppliedTo], DataIn, AppliedTo any]`

#### property RuleRenderComponent In
Type: `DataIn`

#### property RuleRenderComponent AppliedTo
Type: `AppliedTo`

### type InteractionsComponent
Type: `engine/modules/feature.InteractionsComponent`

#### property InteractionsComponent Interactions
Type: `[]any`

#### property InteractionsComponent PreviewInteraction
Type: `any`

#### property InteractionsComponent Features
Type: `[]engine/modules/feature.FeatureKey`

#### property InteractionsComponent LockedFeature
Type: `bool`

#### method InteractionsComponent Equal
Type: `func(other engine/modules/feature.InteractionsComponent) bool`

### type LockFeatureEvent
Type: `engine/modules/feature.LockFeatureEvent`

#### property LockFeatureEvent Feature
Type: `engine/modules/feature.FeatureKey`
to unlock use empty value

## Variables
### var ErrInvalidType
Type: `error`

### var MissingField
Type: `engine/modules/feature.FieldRender`
is used when value is empty and we need to render it

### var PreviewRulesFail
Type: `engine/modules/feature.FieldRender`
is used when [InteractionsComponent].PreviewInteraction rules fail

### var PreviewRulesPass
Type: `engine/modules/feature.FieldRender`
is used when [InteractionsComponent].PreviewInteraction rules pass

### var Interaction
Type: `engine/modules/feature.FieldRender`
is used for each Interaction in [InteractionsComponent].Interactions

## Functions
### func NewFieldRender
Type: `func[FieldType any](value FieldType, render engine/modules/feature.FieldRender) engine/modules/feature.FieldRenderComponent[FieldType]`

### func NewRuleRender
Type: `func[RuleType engine/modules/feature.Rule[DataIn, AppliedTo], DataIn, AppliedTo any](in DataIn, appliedTo AppliedTo) engine/modules/feature.RuleRenderComponent[RuleType, DataIn, AppliedTo]`

### func NewInteractions
Type: `func(interactions []any, previewInteraction any, features []engine/modules/feature.FeatureKey, lockFeature bool) engine/modules/feature.InteractionsComponent`

### func NewLockFeatureEvent
Type: `func(featureKey engine/modules/feature.FeatureKey) engine/modules/feature.LockFeatureEvent`

### func NewUnlockFeatureEvent
Type: `func() engine/modules/feature.LockFeatureEvent`


## Dependencies
`engine`:
  - `engine.EngineWorld`
  - `engine.Events`
  - `engine.EventsBuilder`
  - `engine.Feature`
  - `engine.Hierarchy`
  - `engine.Logger`
  - `engine.World`

`engine/modules/ecs`:
  - `engine/modules/ecs.AnyComponentArray`
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`
  - `engine/modules/ecs.Register`
  - `engine/modules/ecs.SystemRegister`

`engine/modules/feature`:
  - `engine/modules/feature.AnyFieldService`
  - `engine/modules/feature.AnyRuleService`
  - `engine/modules/feature.AppliesToAny`
  - `engine/modules/feature.Equal`
  - `engine/modules/feature.Feature`
  - `engine/modules/feature.FeatureKey`
  - `engine/modules/feature.Features`
  - `engine/modules/feature.FieldKey`
  - `engine/modules/feature.FieldRender`
  - `engine/modules/feature.FieldRenderComponent`
  - `engine/modules/feature.FieldService`
  - `engine/modules/feature.Interaction`
  - `engine/modules/feature.Interactions`
  - `engine/modules/feature.InteractionsComponent`
  - `engine/modules/feature.InteractionsService`
  - `engine/modules/feature.LockFeatureEvent`
  - `engine/modules/feature.LockedFeature`
  - `engine/modules/feature.MissingField`
  - `engine/modules/feature.NewAnyRenderComponent`
  - `engine/modules/feature.NewFieldRender`
  - `engine/modules/feature.NewInteractions`
  - `engine/modules/feature.NewRuleRender`
  - `engine/modules/feature.PreviewInteraction`
  - `engine/modules/feature.PreviewRulesFail`
  - `engine/modules/feature.PreviewRulesPass`
  - `engine/modules/feature.ProvideFieldService`
  - `engine/modules/feature.Render`
  - `engine/modules/feature.RenderAny`
  - `engine/modules/feature.Rule`
  - `engine/modules/feature.RuleKey`
  - `engine/modules/feature.RuleRenderComponent`
  - `engine/modules/feature.RuleService`
  - `engine/modules/feature.Service`

`engine/modules/feature/pkg`:
  - `engine/modules/feature/pkg.Config`
  - `engine/modules/feature/pkg.FieldPkg`
  - `engine/modules/feature/pkg.RegisterFeature`
  - `engine/modules/feature/pkg.RegisterFieldProvider`
  - `engine/modules/feature/pkg.RegisterRuleToField`
  - `engine/modules/feature/pkg.RegisterSystems`
  - `engine/modules/feature/pkg.RulePkg`

`engine/modules/loop`:
  - `engine/modules/loop.TickEvent`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`