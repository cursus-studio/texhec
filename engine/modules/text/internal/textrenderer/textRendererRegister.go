package textrenderer

import (
	_ "embed"
	"engine"
	"engine/modules/graphics"
	"engine/modules/text"
	"engine/services/datastructures"
	"engine/services/ecs"

	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

//go:embed shader.vert
var vertSource string

//go:embed shader.geom
var geomSource string

//go:embed shader.frag
var fragSource string

type textRendererRegister struct {
	engine.EngineWorld `inject:""`
	FontService        FontService                `inject:""`
	VboFactory         graphics.VBOFactory[Glyph] `inject:""`
	LayoutService      LayoutService              `inject:""`
	FontsKeys          FontKeys                   `inject:""`

	removeOncePerNCalls uint16
}

func NewTextRenderer(c ioc.Dic,
	removeOncePerNCalls uint16,
) text.SystemRenderer {
	s := ioc.GetServices[*textRendererRegister](c)
	s.removeOncePerNCalls = removeOncePerNCalls
	return s
}

func (f *textRendererRegister) Register() error {
	vert, err := f.Graphics().NewShader(vertSource, graphics.VertexShader)
	if err != nil {
		return err
	}
	defer vert.Release()

	geom, err := f.Graphics().NewShader(geomSource, graphics.GeomShader)
	if err != nil {
		return err
	}
	defer geom.Release()

	frag, err := f.Graphics().NewShader(fragSource, graphics.FragmentShader)
	if err != nil {
		return err
	}
	defer frag.Release()

	programID := gl.CreateProgram()
	gl.AttachShader(programID, vert.ID())
	gl.AttachShader(programID, geom.ID())
	gl.AttachShader(programID, frag.ID())

	p, err := f.Graphics().NewProgram(programID, nil)
	if err != nil {
		return err
	}

	locations, err := graphics.GetProgramLocations[locations](p)
	if err != nil {
		p.Release()
		return err
	}

	renderer := &textRenderer{
		textRendererRegister: f,

		program:   p,
		locations: locations,

		fontsBatches: datastructures.NewSparseArray[FontKey, fontBatch](),

		dirtyEntities:  ecs.NewDirtySet(),
		layoutsBatches: datastructures.NewSparseArray[ecs.EntityID, layoutBatch](),
	}

	renderer.Transform().AddDirtySet(renderer.dirtyEntities)

	arrays := []ecs.AnyComponentArray{
		ecs.GetComponentsArray[text.TextComponent](f.World()),
		ecs.GetComponentsArray[text.BreakComponent](f.World()),
		ecs.GetComponentsArray[text.FontFamilyComponent](f.World()),
		// ecs.GetComponentsArray[text.Overflow](w),
		ecs.GetComponentsArray[text.FontSizeComponent](f.World()),
		ecs.GetComponentsArray[text.AlignComponent](f.World()),
	}

	for _, array := range arrays {
		array.AddDirtySet(renderer.dirtyEntities)
	}

	events.Listen(f.EventsBuilder(), renderer.ListenRender)

	return nil
}
