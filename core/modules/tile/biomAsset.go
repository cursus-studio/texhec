package tile

import (
	"image"
)

//

type BiomAsset interface {
	Images() [15][]image.Image
	Res() image.Rectangle
	AspectRatio() image.Rectangle
	Release()
}
