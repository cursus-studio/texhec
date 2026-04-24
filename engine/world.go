package engine

import (
	"engine/modules/assets"
	"engine/modules/audio"
	"engine/modules/batcher"
	"engine/modules/camera"
	"engine/modules/codec"
	"engine/modules/collider"
	"engine/modules/connection"
	"engine/modules/groups"
	"engine/modules/hierarchy"
	"engine/modules/inputs"
	"engine/modules/layout"
	"engine/modules/loop"
	"engine/modules/metadata"
	"engine/modules/netsync"
	"engine/modules/noise"
	"engine/modules/prototype"
	"engine/modules/record"
	"engine/modules/registry"
	"engine/modules/render"
	"engine/modules/scene"
	"engine/modules/smooth"
	"engine/modules/text"
	"engine/modules/transform"
	"engine/modules/transition"
	"engine/modules/uuid"
	"engine/modules/warmup"
	"engine/modules/window"
	"engine/services/clock"
	"engine/services/console"
	"engine/services/ecs"
	"engine/services/graphics/texturearray"
	"engine/services/logger"

	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

type EngineWorld struct {
	World         ioc.Lazy[ecs.World]      `inject:""`
	EventsBuilder ioc.Lazy[events.Builder] `inject:""`
	Events        ioc.Lazy[events.Events]  `inject:""`

	Assets     ioc.Lazy[assets.Service]     `inject:""`
	Audio      ioc.Lazy[audio.Service]      `inject:""`
	Batcher    ioc.Lazy[batcher.Service]    `inject:""`
	Camera     ioc.Lazy[camera.Service]     `inject:""`
	Codec      ioc.Lazy[codec.Service]      `inject:""`
	Collider   ioc.Lazy[collider.Service]   `inject:""`
	Connection ioc.Lazy[connection.Service] `inject:""`
	Groups     ioc.Lazy[groups.Service]     `inject:""`
	Hierarchy  ioc.Lazy[hierarchy.Service]  `inject:""`
	Inputs     ioc.Lazy[inputs.Service]     `inject:""`
	Layout     ioc.Lazy[layout.Service]     `inject:""`
	Loop       ioc.Lazy[loop.Service]       `inject:""`
	Metadata   ioc.Lazy[metadata.Service]   `inject:""`
	NetSync    ioc.Lazy[netsync.Service]    `inject:""`
	Noise      ioc.Lazy[noise.Service]      `inject:""`
	Prototype  ioc.Lazy[prototype.Service]  `inject:""`
	Record     ioc.Lazy[record.Service]     `inject:""`
	Registry   ioc.Lazy[registry.Service]   `inject:""`
	Render     ioc.Lazy[render.Service]     `inject:""`
	Scene      ioc.Lazy[scene.Service]      `inject:""`
	Smooth     ioc.Lazy[smooth.Service]     `inject:""`
	Text       ioc.Lazy[text.Service]       `inject:""`
	Transform  ioc.Lazy[transform.Service]  `inject:""`
	Transition ioc.Lazy[transition.Service] `inject:""`
	UUID       ioc.Lazy[uuid.Service]       `inject:""`
	WarmUp     ioc.Lazy[warmup.Service]     `inject:""`
	Window     ioc.Lazy[window.Service]     `inject:""`

	Clock   ioc.Lazy[clock.Clock]     `inject:""`
	Console ioc.Lazy[console.Console] `inject:""`
	// graphics {
	TextureArrayFactory ioc.Lazy[texturearray.Factory] `inject:""`
	// }
	Logger ioc.Lazy[logger.Logger] `inject:""`
}
