# feature
## Architecture
This is second attempt to create `feature` architecture, previous was in module called `interactions`.

This module relies on its own `DSL`:
- ### `Field`
`FeatureEvent` structs are composed from fields and each can be either:
#### `Provided`
Field provided by code
#### `Interaction`
Field provided by interaction with GUI

- ### `Rule`
Rule wires two fields and runs test on them.
Rule can be anything (for e.g. `OwnsRule`, `CanAttackRule`).
It is sufficient to run all tests on feature.

- ### `Feature`
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
ok  	engine/modules/feature/test	0.015s
```
## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                              13            200            109           1298
Markdown                         1              6              0             33
-------------------------------------------------------------------------------
SUM:                            14            206            109           1331
-------------------------------------------------------------------------------
```
## Dependencies
`engine`:
  - `engine.EngineWorld`
  - `engine.Events`
  - `engine.EventsBuilder`
  - `engine.Feature`
  - `engine.Logger`
  - `engine.World`

`engine/modules/ecs`:
  - `engine/modules/ecs.AnyComponentArray`
  - `engine/modules/ecs.ComponentArray`
  - `engine/modules/ecs.ComponentArrayListeners`
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.GetComponentArray`
  - `engine/modules/ecs.Register`
  - `engine/modules/ecs.SystemRegister`

`engine/modules/feature`:
  - `engine/modules/feature.AllFields`
  - `engine/modules/feature.AnyFieldService`
  - `engine/modules/feature.AnyRuleService`
  - `engine/modules/feature.AppliedToField`
  - `engine/modules/feature.AppliesAny`
  - `engine/modules/feature.Equal`
  - `engine/modules/feature.ErrInvalidType`
  - `engine/modules/feature.Feature`
  - `engine/modules/feature.FeatureKey`
  - `engine/modules/feature.Features`
  - `engine/modules/feature.FieldGetter`
  - `engine/modules/feature.FieldKey`
  - `engine/modules/feature.FieldRender`
  - `engine/modules/feature.FieldRenderComponent`
  - `engine/modules/feature.FieldRenderInteraction`
  - `engine/modules/feature.FieldRenderMissing`
  - `engine/modules/feature.FieldRenderPreviewRulesFail`
  - `engine/modules/feature.FieldRenderPreviewRulesPass`
  - `engine/modules/feature.FieldService`
  - `engine/modules/feature.From`
  - `engine/modules/feature.Index`
  - `engine/modules/feature.Interactions`
  - `engine/modules/feature.InteractionsComponent`
  - `engine/modules/feature.InteractionsService`
  - `engine/modules/feature.LockFeatureEvent`
  - `engine/modules/feature.LockedFeature`
  - `engine/modules/feature.NewAnyRenderComponent`
  - `engine/modules/feature.NewFieldGetter`
  - `engine/modules/feature.NewFieldRender`
  - `engine/modules/feature.NewInteractions`
  - `engine/modules/feature.NewRender`
  - `engine/modules/feature.NewRuleRender`
  - `engine/modules/feature.PreviewInteraction`
  - `engine/modules/feature.ProvideFieldService`
  - `engine/modules/feature.Render`
  - `engine/modules/feature.RenderAny`
  - `engine/modules/feature.RenderComponent`
  - `engine/modules/feature.Rule`
  - `engine/modules/feature.RuleKey`
  - `engine/modules/feature.RuleRender`
  - `engine/modules/feature.RuleRenderCompletePreview`
  - `engine/modules/feature.RuleRenderComponent`
  - `engine/modules/feature.RuleRenderPreviewRulesFail`
  - `engine/modules/feature.RuleRenderPreviewRulesPass`
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

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`