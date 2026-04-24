package renderpkg

import (
	"bytes"
	"engine/modules/assets"
	prototypepkg "engine/modules/prototype/pkg"
	"engine/modules/render"
	"engine/modules/render/internal/instancing"
	"engine/modules/render/internal/service"
	"engine/modules/render/internal/systems"
	transitionpkg "engine/modules/transition/pkg"
	"engine/services/ecs"
	gtexture "engine/services/graphics/texture"
	"engine/services/graphics/vao/vbo"
	"image"
	"image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strings"
	"unsafe"

	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/ogiusek/ioc/v2"
)

var Pkg = ioc.NewPkg(func(b ioc.Builder) {
	pkgs := []ioc.Pkg{
		transitionpkg.PkgT[render.ColorComponent],
		transitionpkg.PkgT[render.TextureFrameComponent],

		prototypepkg.PkgT[render.MeshComponent],
		prototypepkg.PkgT[render.TextureComponent],
		prototypepkg.PkgT[render.TextureFrameComponent],
		prototypepkg.PkgT[render.ColorComponent],
	}
	for _, pkg := range pkgs {
		pkg(b)
	}

	ioc.Register(b, func(c ioc.Dic) vbo.VBOFactory[render.Vertex] {
		return func() vbo.VBOSetter[render.Vertex] {
			vbo := vbo.NewVBO[render.Vertex](func() {
				gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false,
					int32(unsafe.Sizeof(render.Vertex{})), uintptr(unsafe.Offsetof(render.Vertex{}.Pos)))
				gl.EnableVertexAttribArray(0)

				gl.VertexAttribPointerWithOffset(1, 2, gl.FLOAT, false,
					int32(unsafe.Sizeof(render.Vertex{})), uintptr(unsafe.Offsetof(render.Vertex{}.TexturePos)))
				gl.EnableVertexAttribArray(1)
			})
			return vbo
		}
	})

	ioc.Register(b, func(c ioc.Dic) render.Service {
		return service.NewService(c)
	})

	ioc.Register(b, func(c ioc.Dic) render.System {
		return ecs.NewSystemRegister(func() error {
			errs := ecs.RegisterSystems(
				systems.NewErrorLogger(c),
				systems.NewRenderSystem(c),
			)
			if len(errs) != 0 {
				return errs[0]
			}
			return nil
		})
	})

	ioc.Register(b, func(c ioc.Dic) render.SystemRenderer {
		return ecs.NewSystemRegister(func() error {
			errs := ecs.RegisterSystems(
				instancing.NewSystem(c),
			)
			if len(errs) != 0 {
				return errs[0]
			}
			return nil
		})
	})

	ioc.Wrap(b, func(c ioc.Dic, b assets.Service) {
		imageHandler := func(id assets.PathComponent) (assets.Asset, error) {
			source, err := os.ReadFile(id.Path)
			if err != nil {
				return nil, err
			}
			imgFile := bytes.NewBuffer(source)
			img, _, err := image.Decode(imgFile)
			if err != nil {
				return nil, err
			}

			img = gtexture.NewImage(img).FlipV().Image()
			return render.NewTextureAsset(img)
		}
		trimImageHandler := func(id assets.PathComponent) (assets.Asset, error) {
			source, err := os.ReadFile(strings.TrimSuffix(id.Path, "-trim"))
			if err != nil {
				return nil, err
			}
			imgFile := bytes.NewBuffer(source)
			img, _, err := image.Decode(imgFile)
			if err != nil {
				return nil, err
			}

			img = gtexture.NewImage(img).FlipV().TrimTransparentBackground().Image()
			return render.NewTextureAsset(img)
		}
		b.Register("png", imageHandler)
		b.Register("jpg", imageHandler)
		b.Register("jpeg", imageHandler)

		b.Register("png-trim", trimImageHandler)
		b.Register("jpg-trim", trimImageHandler)
		b.Register("jpeg-trim", trimImageHandler)

		gifHandler := func(id assets.PathComponent) (assets.Asset, error) {
			source, err := os.ReadFile(id.Path)
			if err != nil {
				return nil, err
			}
			imgFile := bytes.NewBuffer(source)
			gif, err := gif.DecodeAll(imgFile)
			if err != nil {
				return nil, err
			}

			images := make([]image.Image, 0, len(gif.Image))
			for _, img := range gif.Image {
				finalImg := gtexture.NewImage(img).FlipV().Image()
				images = append(images, finalImg)
			}

			return render.NewTextureAsset(images...)
		}

		gifTrimHandler := func(id assets.PathComponent) (assets.Asset, error) {
			source, err := os.ReadFile(strings.TrimSuffix(id.Path, "-trim"))
			if err != nil {
				return nil, err
			}
			imgFile := bytes.NewBuffer(source)
			gif, err := gif.DecodeAll(imgFile)
			if err != nil {
				return nil, err
			}

			images := make([]image.Image, 0, len(gif.Image))
			for _, img := range gif.Image {
				finalImg := gtexture.NewImage(img).FlipV().TrimTransparentBackground().Image()
				images = append(images, finalImg)
			}

			return render.NewTextureAsset(images...)
		}

		b.Register("gif", gifHandler)
		b.Register("gif-trim", gifTrimHandler)
	})
})
