package gamescene

import (
	"core/game"
	"core/modules/definitions"
	"core/modules/settings"
	"core/modules/tile"
	"core/modules/ui"
	"engine/modules/camera"
	"engine/modules/collider"
	"engine/modules/grid"
	"engine/modules/groups"
	"engine/modules/inputs"
	"engine/modules/render"
	"engine/modules/seed"
	"engine/modules/transform"
	"engine/modules/uuid"
	"engine/services/ecs"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

// max zoom to see tiles in 1000 px
const MAX_ZOOM = 1000 // 1000
const MAP_SIZE = 1024 // 1024

func addScene(world game.GameWorld, sceneParent ecs.EntityID) {
	// biggest maps on mods in rusted warfare 2560x1440
	// - all tiles are rendered at once
	// - strategic map is used at some point
	// biggest zoom out in factorio is 448x256 (in 4k)

	{
		uiCamera := world.World().NewEntity()
		world.Hierarchy().SetParent(uiCamera, sceneParent)
		world.Camera().Priority().Set(uiCamera, camera.NewPriority(1))
		world.Camera().Ortho().Set(uiCamera, camera.NewOrtho(-1000, +1000))
		world.Groups().Component().Set(uiCamera, groups.EmptyGroups().Ptr().Enable(definitions.UiGroup).Val())
		world.Ui().UiCamera().Set(uiCamera, ui.UiCameraComponent{})
		world.Ui().CursorCamera().Set(uiCamera, ui.CursorCameraComponent{})

		settingsEntity := world.World().NewEntity()
		world.Hierarchy().SetParent(settingsEntity, uiCamera)
		world.Transform().Pos().Set(settingsEntity, transform.NewPos(10, -10, 0))
		world.Transform().Size().Set(settingsEntity, transform.NewSize(50, 50, 1))
		world.Transform().PivotPoint().Set(settingsEntity, transform.NewPivotPoint(0, 1, .5))
		world.Transform().Parent().Set(settingsEntity, transform.NewParent(transform.RelativePos))
		world.Transform().ParentPivotPoint().Set(settingsEntity, transform.NewParentPivotPoint(0, 1, .5))
		world.Groups().Component().Set(settingsEntity, groups.EmptyGroups().Ptr().Enable(definitions.UiGroup).Val())

		world.Render().Mesh().Set(settingsEntity, render.NewMesh(world.Definitions().Assets().SquareMesh))
		world.Render().Texture().Set(settingsEntity, render.NewTexture(world.Definitions().Hud().Settings))

		world.Inputs().LeftClick().Set(settingsEntity, inputs.NewLeftClick(settings.EnterSettingsEvent{}))
		world.Inputs().KeepSelected().Set(settingsEntity, inputs.KeepSelectedComponent{})
		world.Collider().Component().Set(settingsEntity, collider.NewCollider(world.Definitions().Assets().SquareCollider))
	}

	{
		bgCamera := world.World().NewEntity()
		world.Hierarchy().SetParent(bgCamera, sceneParent)
		world.Camera().Priority().Set(bgCamera, camera.NewPriority(-1))
		world.Camera().Ortho().Set(bgCamera, camera.NewOrtho(-1000, +1000))
		world.Groups().Component().Set(bgCamera, groups.EmptyGroups().Ptr().Enable(definitions.BgGroup).Val())

		bg := world.World().NewEntity()
		world.Hierarchy().SetParent(bg, bgCamera)
		world.Transform().Parent().Set(bg, transform.NewParent(transform.RelativePos|transform.RelativeSizeXY))
		world.Groups().InheritGroups(bg)
		world.Ui().AnimatedBackground().Set(bg, ui.AnimatedBackgroundComponent{})
	}

	worldEntity := world.World().NewEntity()
	world.Hierarchy().SetParent(worldEntity, sceneParent)
	world.Groups().Component().Set(worldEntity, groups.EmptyGroups().Ptr().Enable(definitions.GameGroup).Val())
	world.Tile().Config().Set(worldEntity, tile.NewConfig(
		// seed.New(world.Clock.Now().Unix()),
		seed.New(21377137),
	))

	gameCamera := world.World().NewEntity()
	world.Hierarchy().SetParent(gameCamera, worldEntity)
	world.UUID().Component().Set(gameCamera, uuid.New([16]byte{48}))
	world.Camera().Ortho().Set(gameCamera, camera.NewOrtho(-1000, +1000))
	world.Groups().InheritGroups(gameCamera)
	world.Camera().Mobile().Set(gameCamera, camera.NewMobileCamera())
	world.Camera().Limits().Set(gameCamera, camera.NewCameraLimits(
		10./float32(MAX_ZOOM), 10,
		mgl32.Vec3{0, 0, -1000}, mgl32.Vec3{
			world.Tile().GetTileSize().Size[0] * float32(MAP_SIZE),
			world.Tile().GetTileSize().Size[1] * float32(MAP_SIZE),
			1000,
		},
	))

	for x := range MAP_SIZE / world.Grid().ChunkSize() {
		for y := range MAP_SIZE / world.Grid().ChunkSize() {
			events.Emit(world.Events(), tile.NewMissingChunkEvent(grid.NewChunkCoords(x, y)))
		}
	}
}

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	ioc.Register(b, func(c ioc.Dic) game.GameBuilder {
		return func(sceneParent ecs.EntityID) {
			world := ioc.Get[game.GameWorld](c)
			addScene(world, sceneParent)
		}
	})
})
