package service

import (
	"engine"
	"engine/modules/camera"
	"engine/modules/collider"
	"engine/modules/ecs"
	"engine/modules/window"
	"engine/services/datastructures"
	"fmt"
	"reflect"
	"slices"
	"sort"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
	"github.com/veandco/go-sdl2/sdl"
)

// extra data

type Service interface {
	RegisterProjection(
		reflect.Type,
		ProjectionData,
	) error
	camera.Service
}

type ProjectionData struct {
	Mat4     func(entity ecs.EntityID) mgl32.Mat4
	ShootRay func(entity ecs.EntityID, mousePos window.MousePos) collider.Ray
}

//

// type cameraDataID
type projectionID uint8

type projectionComponent struct {
	projectionID
}

type service struct {
	engine.EngineWorld `inject:""`
	ecs.SystemRegister

	cameraArray      ecs.ComponentArray[camera.Component]
	priorityArray    ecs.ComponentArray[camera.PriorityComponent]
	projectionsArray ecs.ComponentArray[projectionComponent]

	projectionIDs map[reflect.Type]projectionID
	projections   datastructures.SparseArray[projectionID, ProjectionData]

	mobileCamera       ecs.ComponentArray[camera.MobileCameraComponent]
	cameraLimits       ecs.ComponentArray[camera.CameraLimitsComponent]
	viewport           ecs.ComponentArray[camera.ViewportComponent]
	normalizedViewport ecs.ComponentArray[camera.NormalizedViewportComponent]

	ortho              ecs.ComponentArray[camera.OrthoComponent]
	orthoResolution    ecs.ComponentArray[camera.OrthoResolutionComponent]
	perspective        ecs.ComponentArray[camera.PerspectiveComponent]
	dynamicPerspective ecs.ComponentArray[camera.DynamicPerspectiveComponent]
}

func NewService(c ioc.Dic, register ecs.SystemRegister) Service {
	s := ioc.GetServices[*service](c)
	s.SystemRegister = register
	s.cameraArray = ecs.GetComponentArray[camera.Component](s.World())
	s.priorityArray = ecs.GetComponentArray[camera.PriorityComponent](s.World())
	s.projectionsArray = ecs.GetComponentArray[projectionComponent](s.World())

	s.projectionIDs = make(map[reflect.Type]projectionID)
	s.projections = datastructures.NewSparseArray[projectionID, ProjectionData]()

	s.mobileCamera = ecs.GetComponentArray[camera.MobileCameraComponent](s.World())
	s.cameraLimits = ecs.GetComponentArray[camera.CameraLimitsComponent](s.World())
	s.viewport = ecs.GetComponentArray[camera.ViewportComponent](s.World())
	s.normalizedViewport = ecs.GetComponentArray[camera.NormalizedViewportComponent](s.World())

	s.ortho = ecs.GetComponentArray[camera.OrthoComponent](s.World())
	s.orthoResolution = ecs.GetComponentArray[camera.OrthoResolutionComponent](s.World())
	s.perspective = ecs.GetComponentArray[camera.PerspectiveComponent](s.World())
	s.dynamicPerspective = ecs.GetComponentArray[camera.DynamicPerspectiveComponent](s.World())

	s.cameraArray.OnUpsert(s.OnCameraUpsert)

	{
		s.cameraArray.AddDependency(s.ortho)
		s.cameraArray.AddDependency(s.perspective)

		orthoDirtySet := ecs.NewDirtySet()
		s.ortho.AddDirtySet(orthoDirtySet)

		s.cameraArray.BeforeGet(func() {
			entities := orthoDirtySet.Get()
			for _, entity := range entities {
				if !s.World().EntityExists(entity) {
					continue
				}
				s.cameraArray.Set(entity, camera.NewCamera[camera.OrthoComponent]())
			}
		})

		perspectiveDirtySet := ecs.NewDirtySet()
		s.perspective.AddDirtySet(perspectiveDirtySet)

		s.cameraArray.BeforeGet(func() {
			entities := perspectiveDirtySet.Get()
			for _, entity := range entities {
				if !s.World().EntityExists(entity) {
					continue
				}
				s.cameraArray.Set(entity, camera.NewCamera[camera.PerspectiveComponent]())
			}
		})

		events.Listen(s.EventsBuilder(), func(e sdl.WindowEvent) {
			if e.Event == sdl.WINDOWEVENT_RESIZED {
				events.Emit(s.Events(), camera.NewUpdateProjectionsEvent())
			}
		})
	}

	return s
}

func (t *service) Component() ecs.ComponentArray[camera.Component] {
	return t.cameraArray
}

func (t *service) Priority() ecs.ComponentArray[camera.PriorityComponent] {
	return t.priorityArray
}

func (t *service) Mobile() ecs.ComponentArray[camera.MobileCameraComponent] {
	return t.mobileCamera
}
func (t *service) Limits() ecs.ComponentArray[camera.CameraLimitsComponent] {
	return t.cameraLimits
}
func (t *service) Viewport() ecs.ComponentArray[camera.ViewportComponent] {
	return t.viewport
}
func (t *service) NormalizedViewport() ecs.ComponentArray[camera.NormalizedViewportComponent] {
	return t.normalizedViewport
}

func (t *service) Ortho() ecs.ComponentArray[camera.OrthoComponent] {
	return t.ortho
}
func (t *service) OrthoResolution() ecs.ComponentArray[camera.OrthoResolutionComponent] {
	return t.orthoResolution
}
func (t *service) Perspective() ecs.ComponentArray[camera.PerspectiveComponent] {
	return t.perspective
}
func (t *service) DynamicPerspective() ecs.ComponentArray[camera.DynamicPerspectiveComponent] {
	return t.dynamicPerspective
}

//

// returns cameras from smallest to biggest
func (t *service) OrderedCameras() []ecs.EntityID {
	cameras := t.Component().GetEntities()
	cameras = slices.Clone(cameras)
	sort.Slice(cameras, func(i, j int) bool {
		o1, _ := t.priorityArray.Get(cameras[i])
		o2, _ := t.priorityArray.Get(cameras[j])
		return o1.Priority < o2.Priority
	})
	return cameras
}

func (t *service) GetViewport(entity ecs.EntityID) (x, y, w, h int32) {
	viewportComponent, ok := t.viewport.Get(entity)
	if ok {
		return viewportComponent.Viewport()
	}
	normalizedViewportComponent, ok := t.normalizedViewport.Get(entity)
	if ok {
		return normalizedViewportComponent.Viewport(t.Window().Window().GetSize())
	}

	w, h = t.Window().Window().GetSize()
	return 0, 0, w, h
}
func (t *service) Mat4(entity ecs.EntityID) mgl32.Mat4 {
	comp, ok := t.projectionsArray.Get(entity)
	if !ok {
		return mgl32.Mat4{}
	}
	data, ok := t.projections.Get(comp.projectionID)
	if !ok {
		return mgl32.Mat4{}
	}
	return data.Mat4(entity)
}
func (t *service) ShootRay(camera ecs.EntityID, mousePos window.MousePos) collider.Ray {
	comp, ok := t.projectionsArray.Get(camera)
	if !ok {
		return collider.Ray{}
	}
	data, ok := t.projections.Get(comp.projectionID)
	if !ok {
		return collider.Ray{}
	}

	ray := data.ShootRay(camera, mousePos)
	groups, _ := t.Groups().Component().Get(camera)
	ray.Groups = groups
	return ray
}

//

func (t *service) RegisterProjection(
	componentType reflect.Type,
	data ProjectionData,
) error {
	if _, ok := t.projectionIDs[componentType]; ok {
		return nil
	}
	i := len(t.projections.GetIndices())
	if i > int(^projectionID(0)) {
		return fmt.Errorf("exceeded maximal projections count")
	}
	projectionI := projectionID(i)
	t.projectionIDs[componentType] = projectionI
	t.projections.Set(projectionI, data)
	return nil
}

//

func (t *service) OnCameraUpsert(entity ecs.EntityID) {
	cam, ok := t.cameraArray.Get(entity)
	if !ok {
		t.projectionsArray.Remove(entity)
		return
	}
	projID, ok := t.projectionIDs[cam.Projection]
	if !ok {
		t.projectionsArray.Remove(entity)
		return
	}
	projComp := projectionComponent{projID}
	t.projectionsArray.Set(entity, projComp)
}
