package texturearray

import (
	"engine/modules/datastructures"
	"engine/modules/graphics"
	"fmt"
	"image"
	"math"
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

	imagesSize := images.Size()
	if imagesSize < 0 || imagesSize > math.MaxInt16 {
		return nil, fmt.Errorf("invalid images count. Cannot have more than %v images", math.MaxInt16)
	}

	texture, err := createTexs(w, h, images)
	if err != nil {
		return nil, err
	}

	array.texture = texture
	array.imagesCount = int16(imagesSize)

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
