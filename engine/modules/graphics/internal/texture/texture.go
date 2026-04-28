package texture

import (
	"engine/modules/graphics"
	"image"

	"github.com/go-gl/gl/v4.5-core/gl"
)

type texture struct {
	id uint32
}

func (t *texture) ID() uint32 { return t.id }

func (t *texture) Bind() {
	gl.BindTexture(gl.TEXTURE_2D, t.id)
}

func (t *texture) Release() {
	gl.DeleteTextures(1, &t.id)
}

//

type factory struct {
	wrappers []func(graphics.Texture)
}

func NewFactory() graphics.TextureFactory {
	return &factory{
		wrappers: nil,
	}
}

func (f *factory) New(img image.Image) (graphics.Texture, error) {
	id, err := newTexture(img)
	if err != nil {
		return nil, err
	}
	texture := &texture{
		id: id,
	}
	for _, wrapper := range f.wrappers {
		wrapper(texture)
	}
	return texture, err
}

func (f *factory) Wrap(wrapper func(graphics.Texture)) {
	f.wrappers = append(f.wrappers, wrapper)
}
