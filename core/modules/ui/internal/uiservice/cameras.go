package uiservice

import (
	"core/modules/definitions"
	"core/modules/settings"
	"core/modules/ui"
	"engine/modules/camera"
	"engine/modules/collider"
	"engine/modules/ecs"
	"engine/modules/groups"
	"engine/modules/inputs"
	"engine/modules/render"
	"engine/modules/transform"
	"engine/modules/uuid"

	"github.com/go-gl/mathgl/mgl32"
)

type gameCameraComponent struct{}

func (s *service) GameCamera() ecs.ComponentArray[ui.GameCameraComponent] { return s.gameCamera }

func (s *service) GetCamera(gameCameraEntity ecs.EntityID) ecs.EntityID {
	if entities := s.camera.GetEntities(); len(entities) != 0 {
		return entities[0]
	}
	sceneParent := s.Scene().Scene()
	s.Hierarchy().SetParent(gameCameraEntity, sceneParent)
	s.Groups().Component().Set(gameCameraEntity, groups.EmptyGroups().Enable(definitions.GameGroup))

	gameCamera := s.World().NewEntity()
	s.Hierarchy().SetParent(gameCamera, sceneParent)
	s.Camera().Priority().Set(gameCamera, camera.NewPriority(0))
	s.camera.Set(gameCamera, gameCameraComponent{})
	s.UUID().Component().Set(gameCamera, uuid.New([16]byte{48}))
	s.Camera().Ortho().Set(gameCamera, camera.NewOrtho(-1000, +1000))
	s.Groups().Component().Set(gameCamera, groups.EmptyGroups().Enable(definitions.GameGroup))
	s.Camera().Mobile().Set(gameCamera, camera.NewMobileCamera())
	return gameCamera
}

func (s *service) OnGameCameraUpsert(entity ecs.EntityID) {
	gameCamera := s.GetCamera(entity)
	gameCameraComp, _ := s.GameCamera().Get(entity)

	tile := s.Grid().GetTileSize().Size[0]
	tilesPerSide := float32(gameCameraComp.MaxChunksZoomOut * uint32(s.Grid().ChunkSize()))
	mapSide := tilesPerSide * tile

	x, y, w, h := s.Camera().GetViewport(gameCamera)
	viewW := float32(w - x)
	viewH := float32(h - y)

	minZoom := max(viewW/mapSide, viewH/mapSide)

	s.Camera().Limits().Set(gameCamera, camera.NewCameraLimits(
		minZoom, 10,
		mgl32.Vec3{0, 0, -1000},
		mgl32.Vec3{mapSide, mapSide, 1000},
	))
}

func (s *service) AddUICamera(sceneParent ecs.EntityID) ecs.EntityID {
	uiCamera := s.World().NewEntity()
	s.Hierarchy().SetParent(uiCamera, sceneParent)
	s.Camera().Priority().Set(uiCamera, camera.NewPriority(1))
	s.Camera().Ortho().Set(uiCamera, camera.NewOrtho(-1000, +1000))
	s.Groups().Component().Set(uiCamera, groups.EmptyGroups().Enable(definitions.UiGroup))
	s.Ui().UiCamera().Set(uiCamera, ui.UiCameraComponent{})
	s.Ui().CursorCamera().Set(uiCamera, ui.CursorCameraComponent{})

	settingsEntity := s.World().NewEntity()
	s.Hierarchy().SetParent(settingsEntity, uiCamera)
	s.Transform().Pos().Set(settingsEntity, transform.NewPos(10, -10, 0))
	s.Transform().Size().Set(settingsEntity, transform.NewSize(50, 50, 1))
	s.Transform().PivotPoint().Set(settingsEntity, transform.NewPivotPoint(0, 1, .5))
	s.Transform().Inherit().Set(settingsEntity, transform.NewInherit(transform.RelativePos))
	s.Transform().ParentPivotPoint().Set(settingsEntity, transform.NewParentPivotPoint(0, 1, .5))
	s.Groups().Component().Set(settingsEntity, groups.EmptyGroups().Enable(definitions.UiGroup))

	s.Render().Mesh().Set(settingsEntity, render.NewMesh(s.Definitions().Assets().SquareMesh))
	s.Render().Texture().Set(settingsEntity, render.NewTexture(s.Definitions().Hud().Settings))

	s.Inputs().LeftClick().Set(settingsEntity, inputs.NewLeftClick(settings.EnterSettingsEvent{}))
	s.Inputs().KeepSelected().Set(settingsEntity, inputs.KeepSelectedComponent{})
	s.Collider().Component().Set(settingsEntity, collider.NewCollider(s.Definitions().Assets().SquareCollider))
	return uiCamera
}
func (s *service) AddBgCamera(sceneParent ecs.EntityID) ecs.EntityID {
	bgCamera := s.World().NewEntity()
	s.Hierarchy().SetParent(bgCamera, sceneParent)
	s.Camera().Priority().Set(bgCamera, camera.NewPriority(-1))
	s.Camera().Ortho().Set(bgCamera, camera.NewOrtho(-1000, +1000))
	s.Groups().Component().Set(bgCamera, groups.EmptyGroups().Enable(definitions.BgGroup))

	bg := s.World().NewEntity()
	s.Hierarchy().SetParent(bg, bgCamera)
	s.Transform().Inherit().Set(bg, transform.NewInherit(transform.RelativePos|transform.RelativeSizeXY))
	s.Groups().InheritGroups(bg)
	s.Ui().AnimatedBackground().Set(bg, ui.AnimatedBackgroundComponent{})
	return bgCamera
}
