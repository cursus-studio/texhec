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
