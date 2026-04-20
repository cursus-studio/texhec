package textrenderer

import (
	_ "embed"
	"engine"
	"engine/modules/text"
	"engine/services/datastructures"
	"engine/services/ecs"
	"engine/services/graphics/program"
	"engine/services/graphics/shader"
	"engine/services/graphics/vao/vbo"

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
	FontService        FontService           `inject:""`
	VboFactory         vbo.VBOFactory[Glyph] `inject:""`
	LayoutService      LayoutService         `inject:""`
	FontsKeys          FontKeys              `inject:""`

	defaultTextAsset    ecs.EntityID
	defaultColor        text.TextColorComponent
	removeOncePerNCalls uint16
}

func NewTextRenderer(c ioc.Dic,
	defaultTextAsset ecs.EntityID,
	defaultColor text.TextColorComponent,
	removeOncePerNCalls uint16,
) text.SystemRenderer {
	s := ioc.GetServices[*textRendererRegister](c)
	s.defaultTextAsset = defaultTextAsset
	s.defaultColor = defaultColor
	s.removeOncePerNCalls = removeOncePerNCalls
	return s
}

func (f *textRendererRegister) Register() error {
	vert, err := shader.NewShader(vertSource, shader.VertexShader)
	if err != nil {
		return err
	}
	defer vert.Release()

	geom, err := shader.NewShader(geomSource, shader.GeomShader)
	if err != nil {
		return err
	}
	defer geom.Release()

	frag, err := shader.NewShader(fragSource, shader.FragmentShader)
	if err != nil {
		return err
	}
	defer frag.Release()

	programID := gl.CreateProgram()
	gl.AttachShader(programID, vert.ID())
	gl.AttachShader(programID, geom.ID())
	gl.AttachShader(programID, frag.ID())

	p, err := program.NewProgram(programID, nil)
	if err != nil {
		return err
	}

	locations, err := program.GetProgramLocations[locations](p)
	if err != nil {
		p.Release()
		return err
	}

	renderer := &textRenderer{
		textRendererRegister: f,

		program:   p,
		locations: locations,

		defaultColor: f.defaultColor,

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
		ecs.GetComponentsArray[text.TextAlignComponent](f.World()),
	}

	for _, array := range arrays {
		array.AddDirtySet(renderer.dirtyEntities)
	}

	events.Listen(f.EventsBuilder(), renderer.ListenRender)

	return nil
}
