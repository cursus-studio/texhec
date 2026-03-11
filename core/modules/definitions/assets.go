package definitions

import (
	"engine/services/ecs"
)

// asset ID should be a number.
// asset path and its dispatcher should be pointed by id.
// and this approach should be used by every registry
type Assets struct {
	Hud        HudAssets
	Tiles      TileAssets
	Constructs ConstructAssets
	Units      UnitAssets

	ExampleAudio ecs.EntityID `path:"audio.wav"`

	Blank          ecs.EntityID `path:"blank texture"`
	SquareMesh     ecs.EntityID `path:"square mesh"`
	SquareCollider ecs.EntityID `path:"square collider"`
	FontAsset      ecs.EntityID `path:"font1.ttf"`
}

type HudAssets struct {
	Btn         ecs.EntityID `path:"hud/btn.png-trim"`
	Cursor      ecs.EntityID `path:"hud/cursor.png-trim"`
	Settings    ecs.EntityID `path:"hud/settings.png-trim"`
	Background1 ecs.EntityID `path:"hud/bg1.gif-trim"`
	Background2 ecs.EntityID `path:"hud/bg2.gif-trim"`
}

type TileAssets struct {
	Grass    ecs.EntityID `path:"tiles/grass.biom"`
	Sand     ecs.EntityID `path:"tiles/sand.biom"`
	Mountain ecs.EntityID `path:"tiles/mountain.biom"`
	Water    ecs.EntityID `path:"tiles/water.biom"`
}

type ConstructAssets struct {
	Farm ecs.EntityID `path:"constructs/farm.png"`
}

type UnitAssets struct {
	Tank ecs.EntityID `path:"units/tank.png"`
}
