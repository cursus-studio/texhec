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
There is a component which stores all compatible features with currect interactions.
There are few rules for it:
1. If no feature is selected it stores upmost one interaction and lists all features which first interaction matches
2. If feature is selected it'll ensure state matches and no interactions which fail tests are selected

### GUI Components
1. `InteractionsComponent`
It is global and stores whole state.
All following components will store this state or its consequences.
It stores:
- selected feature or feature isn't selected then which features are compatible
- preview interaction
- interacitons selected

2. `FieldRenderComponent[Field]`
Specialized component for rendering fields which are:
- missing for selected feature to be complete
- preview field which do or do not pass rule
- selected interactions

3. `RuleRenderComponent[Rule]`
Specialized component for rendering rules for preview field
