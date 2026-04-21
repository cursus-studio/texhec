package definitions

import (
	"core/modules/tile"
	"engine/modules/groups"
	"engine/services/ecs"

	"github.com/ogiusek/ioc/v2"
)

const (
	UiGroup groups.Group = iota + 1
	GameGroup
	BgGroup
)

const (
	TileLayer tile.Coord = iota + 1
	ConstructLayer
	PathLayer
	UnitLayer
	TilePlaceholderLayer
	ObjectPlaceholderLayer
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

	Can    ecs.EntityID `path:"hud/can.png-trim"`
	Cannot ecs.EntityID `path:"hud/cannot.png-trim"`
}

// In DI container
// Definitions have more dependencies
type Definitions struct {
	Assets
	Hud     ioc.Lazy[Hud]     `inject:""`
	Tiles   ioc.Lazy[Tiles]   `inject:""`
	Objects ioc.Lazy[Objects] `inject:""`

	Transitions ioc.Lazy[Transitions] `inject:""`
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
	Texhec   ecs.EntityID `path:"tiles/texhec.biom" tile:"" generate:"5" obstruction:"water"`
	Grass    ecs.EntityID `path:"tiles/grass.biom" tile:"" generate:"40" obstruction:"water"`
	Mountain ecs.EntityID `path:"tiles/mountain.biom" tile:"" generate:"5" obstruction:"water lowland"`
}

type Objects struct {
	Farm    ecs.EntityID `path:"constructs/farm.png" name:"farm" object:"construct" obstruction:"lowland" size:"2x2"`
	HouseT1 ecs.EntityID `path:"constructs/houseT1.png" name:"house t1" object:"construct" obstruction:"lowland" size:"1x1"`
	HouseT2 ecs.EntityID `path:"constructs/houseT2.png" name:"house t2" object:"construct" obstruction:"lowland" size:"2x2"`
	HouseT3 ecs.EntityID `path:"constructs/houseT3.png" name:"house t3" object:"construct" obstruction:"lowland" size:"3x3"`
	HouseT4 ecs.EntityID `path:"constructs/houseT4.png" name:"house t4" object:"construct" obstruction:"lowland" size:"4x4"`

	Tank ecs.EntityID `path:"units/tank.png" name:"tank" object:"unit" obstruction:"lowland" speed:"2"`
}
