package camerapkg

import (
	"engine"
	"engine/modules/camera"
	"engine/modules/camera/internal/cameralimitsys"
	"engine/modules/camera/internal/mobilecamerasys"
	"engine/modules/camera/internal/projectionsys"
	"engine/modules/camera/internal/service"
	"engine/modules/collider"
	typeregistrypkg "engine/modules/typeregistry/pkg"
	"engine/modules/window"
	"engine/services/ecs"
	"errors"
	"reflect"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ogiusek/ioc/v2"
	"github.com/veandco/go-sdl2/sdl"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		typeregistrypkg.PkgT[camera.Component],
		typeregistrypkg.PkgT[camera.MobileCameraComponent],
		typeregistrypkg.PkgT[camera.CameraLimitsComponent],
		typeregistrypkg.PkgT[camera.ViewportComponent],
		typeregistrypkg.PkgT[camera.NormalizedViewportComponent],

		typeregistrypkg.PkgT[camera.OrthoComponent],
		typeregistrypkg.PkgT[camera.OrthoResolutionComponent],
		typeregistrypkg.PkgT[camera.PerspectiveComponent],
		typeregistrypkg.PkgT[camera.DynamicPerspectiveComponent],

		typeregistrypkg.PkgT[camera.ChangedResolutionEvent],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}
	ioc.Register(b, func(c ioc.Dic) service.Service {
		return service.NewService(c, ecs.NewSystemRegister(func() error {
			errs := ecs.RegisterSystems(
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
		}))
	})
	ioc.Register(b, func(c ioc.Dic) camera.Service {
		return ioc.Get[service.Service](c)
	})

	ioc.Register(b, func(c ioc.Dic) camera.CameraUp { return camera.CameraUp(mgl32.Vec3{0, 1, 0}) })
	ioc.Register(b, func(c ioc.Dic) camera.CameraForward { return camera.CameraForward(mgl32.Vec3{0, 0, -1}) })

	ioc.Wrap(b, func(c ioc.Dic, s service.Service) {
		world := ioc.GetServices[engine.EngineWorld](c)
		cameraService := s
		// transform := ioc.Get[transform.Service](c)
		if err := s.RegisterProjection(reflect.TypeFor[camera.OrthoComponent](), func() service.ProjectionData {
			getCameraTransformMatrix := func(entity ecs.EntityID) mgl32.Mat4 {
				pos, _ := world.Transform().AbsolutePos().Get(entity)
				rot, _ := world.Transform().AbsoluteRotation().Get(entity)

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
		}()); err != nil {
			world.Logger().Log(err)
		}

		//

		if err := s.RegisterProjection(reflect.TypeFor[camera.PerspectiveComponent](), func() service.ProjectionData {
			getCameraTransformMatrix := func(entity ecs.EntityID) mgl32.Mat4 {
				pos, _ := world.Transform().AbsolutePos().Get(entity)
				rot, _ := world.Transform().AbsoluteRotation().Get(entity)

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
					pos, _ := world.Transform().AbsolutePos().Get(entity)
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
		}()); err != nil {
			world.Logger().Log(err)
		}
	})
})
