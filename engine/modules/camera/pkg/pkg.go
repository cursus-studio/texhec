package camerapkg

import (
	"engine/modules/camera"
	"engine/modules/camera/internal/cameralimitsys"
	"engine/modules/camera/internal/mobilecamerasys"
	"engine/modules/camera/internal/projectionsys"
	"engine/modules/camera/internal/service"
	codecpkg "engine/modules/codec/pkg"
	"engine/modules/collider"
	prototypepkg "engine/modules/prototype/pkg"
	"engine/modules/transform"
	"engine/modules/window"
	"engine/services/ecs"
	"errors"
	"reflect"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
	"github.com/veandco/go-sdl2/sdl"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		codecpkg.PkgT[camera.Component],
		codecpkg.PkgT[camera.MobileCameraComponent],
		codecpkg.PkgT[camera.CameraLimitsComponent],
		codecpkg.PkgT[camera.ViewportComponent],
		codecpkg.PkgT[camera.NormalizedViewportComponent],

		codecpkg.PkgT[camera.OrthoComponent],
		codecpkg.PkgT[camera.OrthoResolutionComponent],
		codecpkg.PkgT[camera.PerspectiveComponent],
		codecpkg.PkgT[camera.DynamicPerspectiveComponent],

		codecpkg.PkgT[camera.ChangedResolutionEvent],

		prototypepkg.PkgT[camera.Component],
		prototypepkg.PkgT[camera.MobileCameraComponent],
		prototypepkg.PkgT[camera.CameraLimitsComponent],
		prototypepkg.PkgT[camera.ViewportComponent],
		prototypepkg.PkgT[camera.NormalizedViewportComponent],

		prototypepkg.PkgT[camera.OrthoComponent],
		prototypepkg.PkgT[camera.OrthoResolutionComponent],
		prototypepkg.PkgT[camera.PerspectiveComponent],
		prototypepkg.PkgT[camera.DynamicPerspectiveComponent],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
	ioc.Register(b, func(c ioc.Dic) service.Service {
		return service.NewService(c)
	})
	ioc.Register(b, func(c ioc.Dic) camera.Service {
		return ioc.Get[service.Service](c)
	})

	ioc.Register(b, func(c ioc.Dic) camera.CameraUp { return camera.CameraUp(mgl32.Vec3{0, 1, 0}) })
	ioc.Register(b, func(c ioc.Dic) camera.CameraForward { return camera.CameraForward(mgl32.Vec3{0, 0, -1}) })

	ioc.Wrap(b, func(c ioc.Dic, s service.Service) {
		transform := ioc.Get[transform.Service](c)
		cameraService := s
		// transform := ioc.Get[transform.Service](c)
		s.Register(reflect.TypeFor[camera.OrthoComponent](), func() service.ProjectionData {
			getCameraTransformMatrix := func(entity ecs.EntityID) mgl32.Mat4 {
				pos, _ := transform.AbsolutePos().Get(entity)
				rot, _ := transform.AbsoluteRotation().Get(entity)

				cameraRot := rot.Rotation.Inverse()
				cameraPos := rot.Rotation.Rotate(pos.Pos.Mul(-1))
				return cameraRot.Mat4().Mul4(mgl32.Translate3D(cameraPos.X(), cameraPos.Y(), cameraPos.Z()))
			}
			getProjectionMatrix := func(entity ecs.EntityID) mgl32.Mat4 {
				p, _ := cameraService.Ortho().Get(entity)
				orthoResolution, ok := cameraService.OrthoResolution().Get(entity)
				if !ok {
					orthoResolution = camera.GetViewportOrthoResolution(cameraService.GetViewport(entity))
				}
				return p.GetMatrix(orthoResolution.Elem())
			}
			return service.ProjectionData{
				Mat4: func(entity ecs.EntityID) mgl32.Mat4 {
					projMatrix := getProjectionMatrix(entity)
					cameraTransformMatrix := getCameraTransformMatrix(entity)
					return projMatrix.Mul4(cameraTransformMatrix)
				},
				ShootRay: func(entity ecs.EntityID, mousePos window.MousePos) collider.Ray {
					return mobilecamerasys.ShootRay(
						getProjectionMatrix(entity),
						getCameraTransformMatrix(entity),
						mousePos,
						func() (x int32, y int32, w int32, h int32) {
							return cameraService.GetViewport(entity)
						},
						nil,
					)
				},
			}
		}())

		//

		s.Register(reflect.TypeFor[camera.PerspectiveComponent](), func() service.ProjectionData {
			getCameraTransformMatrix := func(entity ecs.EntityID) mgl32.Mat4 {
				pos, _ := transform.AbsolutePos().Get(entity)
				rot, _ := transform.AbsoluteRotation().Get(entity)

				up, forward := ioc.Get[camera.CameraUp](c), ioc.Get[camera.CameraForward](c)
				return mgl32.LookAtV(
					pos.Pos,
					pos.Pos.Add(rot.Rotation.Rotate(mgl32.Vec3(forward))),
					mgl32.Vec3(up),
				)
			}
			getProjectionMatrix := func(entity ecs.EntityID) mgl32.Mat4 {
				p, _ := cameraService.Perspective().Get(entity)
				return mgl32.Perspective(p.FovY, p.AspectRatio, p.Near, p.Far)
			}

			return service.ProjectionData{
				Mat4: func(entity ecs.EntityID) mgl32.Mat4 {
					projMatrix := getProjectionMatrix(entity)
					cameraTransformMatrix := getCameraTransformMatrix(entity)
					return projMatrix.Mul4(cameraTransformMatrix)
				},
				ShootRay: func(entity ecs.EntityID, mousePos window.MousePos) collider.Ray {
					pos, _ := transform.AbsolutePos().Get(entity)
					return mobilecamerasys.ShootRay(
						getProjectionMatrix(entity),
						getCameraTransformMatrix(entity),
						mousePos,
						func() (x int32, y int32, w int32, h int32) {
							return cameraService.GetViewport(entity)
						},
						&pos.Pos,
					)
				},
			}
		}())
	})

	ioc.Register(b, func(c ioc.Dic) camera.System {
		return ecs.NewSystemRegister(func() error {
			w := ioc.Get[ecs.World](c)
			eventsBuilder := ioc.Get[events.Builder](c)
			errs := ecs.RegisterSystems(
				ecs.NewSystemRegister(func() error {
					cameraArray := ecs.GetComponentsArray[camera.Component](w)
					orthoArray := ecs.GetComponentsArray[camera.OrthoComponent](w)
					perspectiveArray := ecs.GetComponentsArray[camera.PerspectiveComponent](w)

					cameraArray.AddDependency(orthoArray)
					cameraArray.AddDependency(perspectiveArray)

					orthoDirtySet := ecs.NewDirtySet()
					orthoArray.AddDirtySet(orthoDirtySet)

					cameraArray.BeforeGet(func() {
						entities := orthoDirtySet.Get()
						for _, entity := range entities {
							if !w.EntityExists(entity) {
								continue
							}
							cameraArray.Set(entity, camera.NewCamera[camera.OrthoComponent]())
						}
					})

					perspectiveDirtySet := ecs.NewDirtySet()
					perspectiveArray.AddDirtySet(perspectiveDirtySet)

					cameraArray.BeforeGet(func() {
						entities := perspectiveDirtySet.Get()
						for _, entity := range entities {
							if !w.EntityExists(entity) {
								continue
							}
							cameraArray.Set(entity, camera.NewCamera[camera.PerspectiveComponent]())
						}
					})

					events.Listen(eventsBuilder, func(e sdl.WindowEvent) {
						if e.Event == sdl.WINDOWEVENT_RESIZED {
							events.Emit(eventsBuilder.Events(), camera.NewUpdateProjectionsEvent())
						}
					})
					return nil
				}),
				// todo change this to change ortho and size according to viewport
				projectionsys.NewUpdateProjectionsSystem(c),
				mobilecamerasys.NewScrollSystem(c),
				mobilecamerasys.NewDragSystem(c,
					sdl.BUTTON_LEFT,
				),
				mobilecamerasys.NewWasdSystem(c,
					1.0, // speed
				),
				cameralimitsys.NewOrthoSys(c),
			)
			if len(errs) != 0 {
				return errors.Join(errs...)
			}
			return nil
		})
	})
})
