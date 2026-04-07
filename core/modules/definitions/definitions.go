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
	LowlandsObstruction                              // obstructed by buildings and tanks
)

// In DI container
// BuiltIn have fewer dependencies
type BuiltIn struct {
	ExampleAudio ecs.EntityID `path:"audio.wav"`

	Blank          ecs.EntityID `path:"blank texture"`
	SquareMesh     ecs.EntityID `path:"square mesh"`
	SquareCollider ecs.EntityID `path:"square collider"`
	FontAsset      ecs.EntityID `path:"font1.ttf"`
}

// In DI container
// Definitions have more dependencies
type Definitions struct {
	BuiltIn    `ignore:""`
	Hud        Hud
	Tiles      Tiles
	Constructs Constructs
	Units      Units

	Transitions Transitions
}

type Hud struct {
	Btn         ecs.EntityID `path:"hud/btn.png-trim"`
	Cursor      ecs.EntityID `path:"hud/cursor.png-trim"`
	Settings    ecs.EntityID `path:"hud/settings.png-trim"`
	Background1 ecs.EntityID `path:"hud/bg1.gif-trim"`
	Background2 ecs.EntityID `path:"hud/bg2.gif-trim"`
}

type Transitions struct {
	Linear         ecs.EntityID `transition:"linear"`
	MyEasing       ecs.EntityID `transition:"my easing"`
	EaseOutElastic ecs.EntityID `transition:"ease out elastic"`
}

// domain objects

// generation configs should be in registry or in destined path and dispatched instantly on initialization
type Tiles struct {
	Water    ecs.EntityID `path:"tiles/water.biom" tile:"" generate:"35"`
	Sand     ecs.EntityID `path:"tiles/sand.biom" tile:"" generate:"15"`
	Grass    ecs.EntityID `path:"tiles/grass.biom" tile:"" generate:"45"`
	Mountain ecs.EntityID `path:"tiles/mountain.biom" tile:"" generate:"5"`
}

type Constructs struct {
	Farm ecs.EntityID `path:"constructs/farm.png" name:"farm" construct:""`
}

type Units struct {
	Tank ecs.EntityID `path:"units/tank.png" name:"tank" unit:""`
}
