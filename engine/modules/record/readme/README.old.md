# Record
## Architecture
This module splits into two separate parts.\
One to record by `EntityID` which should be used to perform things localy.\
Second to record by `UUID` which should be used to record things for external machines.

Backwards recording returns state on recording start.\
Backwards recording can be used to record state before to smoothen changes.

Forwards recording returns state on recording end.\
Forwards recording can be used to record changes to send them somewhere else to replicate them.

## Benchmarks

```bash
$ go test . -bench=. -benchmem
goos: linux
goarch: amd64
pkg: engine/modules/record/test
cpu: Intel(R) Core(TM) i5-8350U CPU @ 1.70GHz
BenchmarkEntityCodecEncode-8                	  155301	      7093 ns/op	    3281 B/op	      37 allocs/op
BenchmarkEntityCodecDecode-8                	   47541	     24805 ns/op	   10768 B/op	     245 allocs/op
BenchmarkUUIDCodecEncode-8                  	  209726	      5588 ns/op	    2505 B/op	      35 allocs/op
BenchmarkUUIDCodecDecode-8                  	   60512	     20209 ns/op	    9600 B/op	     214 allocs/op
BenchmarkEntityRecording-8                  	  974848	      1270 ns/op	     746 B/op	      13 allocs/op
BenchmarkCreateNEntitiesEntityRecording-8   	10272049	       120.9 ns/op	     219 B/op	       1 allocs/op
BenchmarkEntityApply1Entities-8             	 6980020	       143.6 ns/op	      80 B/op	       2 allocs/op
BenchmarkEntityApply10Entities-8            	 3908955	       310.5 ns/op	      80 B/op	       2 allocs/op
BenchmarkUUIDRecording-8                    	  956442	      1332 ns/op	     714 B/op	      13 allocs/op
BenchmarkCreateNEntitiesUUIDRecording-8     	 2665312	       478.9 ns/op	     270 B/op	       3 allocs/op
BenchmarkUUIDApply-8                        	 5510028	       202.3 ns/op	      72 B/op	       2 allocs/op
BenchmarkUUIDApply10Entities-8              	 1794826	       659.4 ns/op	      72 B/op	       2 allocs/op
```

## Usage examples
```go
type Service interface {
	Entity() EntityKeyedRecorder
	UUID() UUIDKeyedRecorder
}

type EntityKeyedRecorder interface {
	// gets state as finished recording
	GetState(Config) Recording

	// starts opened recording (opened recording is recorded until stopped)
	// applying it on previous state will create current state
	StartRecording(Config) RecordingID
	// starts opened recording (opened recording is recorded until stopped)
	// applying it rewinds state.
	StartBackwardsRecording(Config) RecordingID
	// finishes recording if open (false is returned if recording isn't started)
	Stop(RecordingID) (r Recording, ok bool)

	Apply(Config, ...Recording)
}


type Recording struct {
	// [componentArrayLayoutID]any component
	// nil for removed entity
	Entities datastructures.SparseArray[ecs.EntityID, []any]
}

type UUIDKeyedRecorder interface {
	// gets state as finished recording
	GetState(Config) UUIDRecording

	// starts opened recording (opened recording is recorded until stopped)
	// applying it on previous state will create current state
	StartRecording(Config) UUIDRecordingID
	// starts opened recording (opened recording is recorded until stopped)
	// applying it rewinds state.
	StartBackwardsRecording(Config) UUIDRecordingID
	// finishes recording if open (false is returned if recording isn't started)
	Stop(UUIDRecordingID) (r UUIDRecording, ok bool)

	Apply(Config, ...UUIDRecording)
}

type UUIDRecording struct {
	// map[componentUUID][componentArrayLayoutID]any component
	// map[componentUUID]nil is when entity is removed
	Entities map[uuid.UUID][]any
}
```

### Entity recording
```go
func _(r record.Service) {
    config := record.NewConfig()
    record.AddToConfig[MyComponent](config)

    recordingID := r.Entity().StartRecording(config)
    // do something
    recording, ok := r.Entity().Stop(recordingID)

    // now we can browse entities.
    // nil for a whole array means that entity doesn't exist
    // nil for a component means that component doesn't exists
}
```

### Entity backwards recording
```go
func _(r record.Service) {
    config := record.NewConfig()
    record.AddToConfig[MyComponent](config)

    recordingID := r.Entity().StartBackwardsRecording(config)
    // do something
    recording, ok := r.Entity().Stop(recordingID)

    // now we can browse entities.
    // nil for a whole array means that entity didn't exist
    // nil for a component means that component didn't exists
}
```

### UUID recording
```go
func _(r record.Service) {
    config := record.NewConfig()
    record.AddToConfig[MyComponent](config)

    recordingID := r.UUID().StartRecording(config)
    // do something
    recording, ok := r.UUID().Stop(recordingID)

    // now we can browse entities.
    // nil for a whole array means that entity doesn't exist
    // nil for a component means that component doesn't exists
}
```

### UUID backwards recording
```go
func _(r record.Service) {
    config := record.NewConfig()
    record.AddToConfig[MyComponent](config)

    recordingID := r.UUID().StartBackwardsRecording(config)
    // do something
    recording, ok := r.UUID().Stop(recordingID)

    // now we can browse entities.
    // nil for a whole array means that entity didn't exist
    // nil for a component means that component didn't exists
}
```

## Dependencies
- [codec](/engine/modules/codec/readme/README.md)
- [datastructures](/engine/modules/datastructures/readme/README.md)
- [ecs](/engine/modules/ecs/readme/README.md)
- [logger](/engine/modules/logger/readme/README.md)
- [uuid](/engine/modules/uuid/readme/README.md)
