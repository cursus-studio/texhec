package texturearray

import (
	"engine/modules/graphics"
	"engine/services/datastructures"
	"image"
)

type factory struct {
	wrappers []func(graphics.TextureArray)
}

func NewFactory() graphics.TextureArrayFactory {
	return &factory{}
}

func (f *factory) New(asset datastructures.SparseArray[uint32, image.Image]) (graphics.TextureArray, error) {
	array := &textureArray{}
	images := datastructures.NewSparseArray[uint32, image.Image]()

	w, h := 0, 0
	if len(asset.GetValues()) != 0 {
		bounds := asset.GetValues()[0].Bounds()
		w, h = bounds.Dx(), bounds.Dy()
	}

	for _, i := range asset.GetIndices() {
		image, _ := asset.Get(i)

		if w != image.Bounds().Dx() || h != image.Bounds().Dy() {
			return nil, graphics.ErrTexturesHaveToShareSize
		}

		images.Set(i, image)
	}

	array.texture = createTexs(w, h, images)
	array.imagesCount = images.Size()

	for _, wrapper := range f.wrappers {
		wrapper(array)
	}

	return array, nil
}

func (f *factory) NewFromSlice(images []image.Image) (graphics.TextureArray, error) {
	arr := datastructures.NewSparseArray[uint32, image.Image]()
	for i, image := range images {
		arr.Set(uint32(i), image)
	}
	return f.New(arr)
}

func (f *factory) Wrap(wrapper func(graphics.TextureArray)) {
	f.wrappers = append(f.wrappers, wrapper)
}
