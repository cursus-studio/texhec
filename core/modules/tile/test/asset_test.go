package test

import (
	"core/game"
	corepkg "core/pkg"
	"engine/modules/grid"
	"image"
	"image/color"
	"testing"

	"github.com/ogiusek/ioc/v2"
)

func newImage(v uint8) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, 2, 2))
	if v&0b1000 != 0 {
		img.Set(0, 0, color.White)
	}
	if v&0b0100 != 0 {
		img.Set(1, 0, color.White)
	}
	if v&0b0010 != 0 {
		img.Set(0, 1, color.White)
	}
	if v&0b0001 != 0 {
		img.Set(1, 1, color.White)
	}
	return img
}

func TestNewAsset(t *testing.T) {
	world := ioc.Get[game.GameWorld](ioc.NewContainer(corepkg.Pkg))
	srcImages := [6][]image.Image{
		{newImage(0b0011)},
		{newImage(0b1111)},
		{newImage(0b1110)},
		{newImage(0b1010)},
		{newImage(0b1001)},
		{newImage(0b0001)},
	}

	biome, err := world.Tile().NewBiomeAsset(srcImages)
	if err != nil {
		t.Error(err)
	}
	images := biome.Images()

	for i, imgs := range images {
		img := imgs[0]
		bounds := img.Bounds()
		// images are flipped so t and b are swapped
		_, _, _, lt := img.At(bounds.Min.X, bounds.Max.Y-1).RGBA()
		_, _, _, rt := img.At(bounds.Max.X-1, bounds.Max.Y-1).RGBA()
		_, _, _, lb := img.At(bounds.Min.X, bounds.Min.Y).RGBA()
		_, _, _, rb := img.At(bounds.Max.X-1, bounds.Min.Y).RGBA()
		v := 0
		if rb > 0 {
			v |= 1 << 3
		}
		if lb > 0 {
			v |= 1 << 2
		}
		if rt > 0 {
			v |= 1 << 1
		}
		if lt > 0 {
			v |= 1 << 0
		}
		expected := i + 1
		if v != expected {
			t.Errorf("expected %4b but got %4b", expected, v)
		}
	}
}

func TestNthCoordinates(t *testing.T) {
	world := ioc.Get[game.GameWorld](ioc.NewContainer(corepkg.Pkg))
	tests := []struct {
		index    grid.Coord
		expected grid.Coords
	}{
		// Center
		{0, grid.NewCoords(0, 0)},

		// Ring 1 (Indices 1-8)
		{1, grid.NewCoords(1, 0)},
		{2, grid.NewCoords(1, 1)},
		{3, grid.NewCoords(0, 1)},
		{4, grid.NewCoords(-1, 1)},
		{5, grid.NewCoords(-1, 0)},
		{6, grid.NewCoords(-1, -1)},
		{7, grid.NewCoords(0, -1)},
		{8, grid.NewCoords(1, -1)},

		// Ring 2 (Indices 9-24)
		{9, grid.NewCoords(2, -1)},
		{10, grid.NewCoords(2, 0)},
		{11, grid.NewCoords(2, 1)},
		{12, grid.NewCoords(2, 2)},
		{13, grid.NewCoords(1, 2)},
		{14, grid.NewCoords(0, 2)},
		{15, grid.NewCoords(-1, 2)},
		{16, grid.NewCoords(-2, 2)},
		{17, grid.NewCoords(-2, 1)},
		{18, grid.NewCoords(-2, 0)},
		{19, grid.NewCoords(-2, -1)},
		{20, grid.NewCoords(-2, -2)},
		{21, grid.NewCoords(-1, -2)},
		{22, grid.NewCoords(0, -2)},
		{23, grid.NewCoords(1, -2)},
		{24, grid.NewCoords(2, -2)},
	}

	for _, tt := range tests {
		got := world.Tile().GetNthNeighbor(tt.index)
		if got != tt.expected {
			t.Errorf("GetCoordinates(%d) = %v, expected %v", tt.index, got, tt.expected)
		}
	}
}
