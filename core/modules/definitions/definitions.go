package definitions

import (
	"core/modules/tile"
	"engine/modules/groups"
	"engine/services/ecs"
)

const (
	UiGroup groups.Group = iota + 1
	GameGroup
	BgGroup
)

const (
	TileLayer tile.Coord = iota + 1
	ConstructLayer
	UnitLayer
	PlaceholderLayer
)

const (
	AirspaceObstruction tile.Obstruction = 1 << iota // obstructed by mountains and planes
	WaterObstruction                                 // obstructed by non-water tiles and ships
	LowlandObstruction                               // obstructed by buildings and tanks
)

// In DI container
// Assets have fewer dependencies
type Assets struct {
	ExampleAudio ecs.EntityID `path:"audio.wav"`

	Blank          ecs.EntityID `path:"blank texture"`
	SquareMesh     ecs.EntityID `path:"square mesh"`
	SquareCollider ecs.EntityID `path:"square collider"`
	FontAsset      ecs.EntityID `path:"font1.ttf"`
}

type Hud struct {
	Btn         ecs.EntityID `path:"hud/btn.png-trim"`
	Text        ecs.EntityID
	Cursor      ecs.EntityID `path:"hud/cursor.png-trim"`
	Settings    ecs.EntityID `path:"hud/settings.png-trim"`
	Background1 ecs.EntityID `path:"hud/bg1.gif-trim"`
	Background2 ecs.EntityID `path:"hud/bg2.gif-trim"`
}

// In DI container
// Definitions have more dependencies
type Definitions struct {
	Assets     `ignore:""`
	Hud        Hud
	Tiles      Tiles
	Constructs Constructs
	Units      Units

	Transitions Transitions
}

type Transitions struct {
	Linear         ecs.EntityID `transition:"linear"`
	MyEasing       ecs.EntityID `transition:"my easing"`
	EaseOutElastic ecs.EntityID `transition:"ease out elastic"`
}

// domain objects

// generation configs should be in registry or in destined path and dispatched instantly on initialization
type Tiles struct {
	Water    ecs.EntityID `path:"tiles/water.biom" tile:"" generate:"35" obstruction:"lowland"`
	Sand     ecs.EntityID `path:"tiles/sand.biom" tile:"" generate:"15" obstruction:"water"`
	Grass    ecs.EntityID `path:"tiles/grass.biom" tile:"" generate:"45" obstruction:"water"`
	Mountain ecs.EntityID `path:"tiles/mountain.biom" tile:"" generate:"5" obstruction:"water lowland"`
}

type Constructs struct {
	Farm ecs.EntityID `path:"constructs/farm.png" name:"farm" construct:"" obstruction:"lowland"`
}

type Units struct {
	Tank ecs.EntityID `path:"units/tank.png" name:"tank" unit:"" obstruction:"lowland"`
}
