# audio
## Architecture
this module is reponsible for integrating audio via events

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               6             51              6            260
-------------------------------------------------------------------------------
SUM:                             6             51              6            260
-------------------------------------------------------------------------------
```
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

`engine/modules/datastructures`:
  - `engine/modules/datastructures.NewSparseArray`
  - `engine/modules/datastructures.SparseArray`

`engine/modules/ecs`:
  - `engine/modules/ecs.EntityID`
  - `engine/modules/ecs.NewSystemRegister`
  - `engine/modules/ecs.SystemRegister`

`engine/modules/typeregistry/pkg`:
  - `engine/modules/typeregistry/pkg.PkgT`

### Third Party
- `github.com/ogiusek/events`
- `github.com/ogiusek/ioc/v2`
- `github.com/veandco/go-sdl2/mix`