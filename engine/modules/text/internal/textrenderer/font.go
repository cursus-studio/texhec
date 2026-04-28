package textrenderer

import (
	"engine"
	"engine/modules/assets"
	"engine/modules/text"
	"engine/services/datastructures"
	"engine/services/ecs"

	"github.com/ogiusek/ioc/v2"
	"golang.org/x/image/font/opentype"
)

type FontKey uint32

type FontService interface {
	AssetFont(ecs.EntityID) (text.Glyphs, error)
}

type fontService struct {
	engine.EngineWorld `inject:""`
	usedGlyphs         datastructures.SparseSet[rune]
	faceOptions        opentype.FaceOptions

	cellSize, yBaseline int
}

func NewFontService(
	c ioc.Dic,
	usedGlyphs datastructures.SparseSet[rune],
	face opentype.FaceOptions,
	cellSize, yBaseline int,
) FontService {
	s := ioc.GetServices[*fontService](c)
	s.usedGlyphs = usedGlyphs
	s.faceOptions = face
	s.cellSize = cellSize
	s.yBaseline = yBaseline
	return s
}

// temporary fix for performance
// there should be in public font asset type and added in factory

//

func (s *fontService) AssetFont(assetID ecs.EntityID) (text.Glyphs, error) {
	asset, err := assets.GetAsset[text.FontFaceAsset](s.Assets(), assetID)
	if err != nil {
		return text.Glyphs{}, err
	}
	return asset.Glyphs(), nil
}
