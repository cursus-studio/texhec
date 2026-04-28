package definitions

import (
	"core/modules/tile"
	"engine/modules/audio"
	"engine/modules/groups"
	"engine/modules/scene"
	"engine/services/ecs"
)

// In DI container
// Definitions have more dependencies
type Service interface {
	Load()

	Assets() Assets
	Hud() Hud
	Tiles() Tiles
	Objects() Objects
	Transitions() Transitions
}

var (
	MenuID     = scene.NewSceneId("menu")
	GameID     = scene.NewSceneId("game")
	SettingsID = scene.NewSceneId("settings")
	CreditsID  = scene.NewSceneId("credits")
)

const (
	EffectChannel audio.Channel = iota
	MusicChannel
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

	Blank          ecs.EntityID
	SquareMesh     ecs.EntityID
	SquareCollider ecs.EntityID
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

type Transitions struct {
	Linear         ecs.EntityID
	MyEasing       ecs.EntityID
	EaseOutElastic ecs.EntityID
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
