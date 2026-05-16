# audio
## Architecture
this module is reponsible for integrating audio via events

## Types
### type Service
Type: `engine/modules/audio.Service`

#### method Service Play
Type: `func(engine/modules/audio.Channel, engine/services/ecs.EntityID) error`

#### method Service Queue
Type: `func(engine/modules/audio.Channel, engine/services/ecs.EntityID) error`

#### method Service QueueEndless
Type: `func(engine/modules/audio.Channel, engine/services/ecs.EntityID) error`

#### method Service Register
Type: `func() error`

#### method Service SetChannelVolume
Type: `func(engine/modules/audio.Channel, engine/modules/audio.Volume) error`

#### method Service SetMasterVolume
Type: `func(engine/modules/audio.Volume) error`

#### method Service Stop
Type: `func(engine/modules/audio.Channel) error`

### type AudioAsset
Type: `engine/modules/audio.AudioAsset`

#### method AudioAsset Chunk
Type: `func() *github.com/veandco/go-sdl2/mix.Chunk`

#### method AudioAsset Release
Type: `func()`

### type PlayerService
Type: `engine/modules/audio.PlayerService`

#### method PlayerService Play
Type: `func(engine/modules/audio.Channel, engine/services/ecs.EntityID) error`

#### method PlayerService Queue
Type: `func(engine/modules/audio.Channel, engine/services/ecs.EntityID) error`

#### method PlayerService QueueEndless
Type: `func(engine/modules/audio.Channel, engine/services/ecs.EntityID) error`

#### method PlayerService Stop
Type: `func(engine/modules/audio.Channel) error`

### type VolumeService
Type: `engine/modules/audio.VolumeService`

#### method VolumeService SetChannelVolume
Type: `func(engine/modules/audio.Channel, engine/modules/audio.Volume) error`

#### method VolumeService SetMasterVolume
Type: `func(engine/modules/audio.Volume) error`

### type StopEvent
Type: `engine/modules/audio.StopEvent`

#### property StopEvent Channel
Type: `engine/modules/audio.Channel`

### type PlayEvent
Type: `engine/modules/audio.PlayEvent`

#### property PlayEvent Channel
Type: `engine/modules/audio.Channel`

#### property PlayEvent Asset
Type: `engine/services/ecs.EntityID`

### type QueueEvent
Type: `engine/modules/audio.QueueEvent`

#### property QueueEvent Channel
Type: `engine/modules/audio.Channel`

#### property QueueEvent Asset
Type: `engine/services/ecs.EntityID`

### type QueueEndlessEvent
Type: `engine/modules/audio.QueueEndlessEvent`

#### property QueueEndlessEvent Channel
Type: `engine/modules/audio.Channel`

#### property QueueEndlessEvent Asset
Type: `engine/services/ecs.EntityID`

### type SetMasterVolumeEvent
Type: `engine/modules/audio.SetMasterVolumeEvent`

#### property SetMasterVolumeEvent Volume
Type: `engine/modules/audio.Volume`

### type SetChannelVolumeEvent
Type: `engine/modules/audio.SetChannelVolumeEvent`

#### property SetChannelVolumeEvent Channel
Type: `engine/modules/audio.Channel`

#### property SetChannelVolumeEvent Volume
Type: `engine/modules/audio.Volume`

### type Channel
Type: `engine/modules/audio.Channel`

### type Volume
Type: `engine/modules/audio.Volume`

## Functions
### func NewAudioAsset
Type: `func(chunk *github.com/veandco/go-sdl2/mix.Chunk, source []byte) engine/modules/audio.AudioAsset`

### func NewStopEvent
Type: `func(channel engine/modules/audio.Channel) engine/modules/audio.StopEvent`

### func NewPlayEvent
Type: `func(channel engine/modules/audio.Channel, assetID engine/services/ecs.EntityID) engine/modules/audio.PlayEvent`

### func NewQueueEvent
Type: `func(channel engine/modules/audio.Channel, assetID engine/services/ecs.EntityID) engine/modules/audio.QueueEvent`

### func NewQueueEndlessEvent
Type: `func(channel engine/modules/audio.Channel, assetID engine/services/ecs.EntityID) engine/modules/audio.QueueEndlessEvent`

### func NewSetMasterVolumeEvent
Type: `func(volume engine/modules/audio.Volume) engine/modules/audio.SetMasterVolumeEvent`

### func NewSetChannelVolumeEvent
Type: `func(channel engine/modules/audio.Channel, volume engine/modules/audio.Volume) engine/modules/audio.SetChannelVolumeEvent`


## Dependencies
`engine/modules/assets`:
  - `engine/modules/assets.Asset`
  - `engine/modules/assets.GetAsset`
  - `engine/modules/assets.Path`
  - `engine/modules/assets.PathComponent`
  - `engine/modules/assets.Register`
  - `engine/modules/assets.Service`

`engine/modules/audio`:
  - `engine/modules/audio.Asset`
  - `engine/modules/audio.AudioAsset`
  - `engine/modules/audio.Channel`
  - `engine/modules/audio.Chunk`
  - `engine/modules/audio.NewAudioAsset`
  - `engine/modules/audio.Play`
  - `engine/modules/audio.PlayEvent`
  - `engine/modules/audio.PlayerService`
  - `engine/modules/audio.Queue`
  - `engine/modules/audio.QueueEndless`
  - `engine/modules/audio.QueueEndlessEvent`
  - `engine/modules/audio.QueueEvent`
  - `engine/modules/audio.Service`
  - `engine/modules/audio.SetChannelVolume`
  - `engine/modules/audio.SetChannelVolumeEvent`
  - `engine/modules/audio.SetMasterVolume`
  - `engine/modules/audio.SetMasterVolumeEvent`
  - `engine/modules/audio.Stop`
  - `engine/modules/audio.StopEvent`
  - `engine/modules/audio.Volume`
  - `engine/modules/audio.VolumeService`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

`engine/services/datastructures`:
  - `engine/services/datastructures.Get`
  - `engine/services/datastructures.GetIndices`
  - `engine/services/datastructures.NewSparseArray`
  - `engine/services/datastructures.SparseArray`

`engine/services/ecs`:
  - `engine/services/ecs.EntityID`
  - `engine/services/ecs.NewSystemRegister`
  - `engine/services/ecs.SystemRegister`

### Third Party
`github.com/ogiusek/events`
`github.com/ogiusek/ioc/v2`
`github.com/veandco/go-sdl2/mix`