package graphics

import (
	"engine/modules/graphics/internal/buffers"
	"engine/modules/graphics/internal/vbo"
	"engine/services/datastructures"
	"errors"
	"image"
	"reflect"

	"github.com/go-gl/gl/v4.5-core/gl"
)

type Service interface {
	NewImage(image image.Image) Image
	TextureArray() TextureArrayFactory
	Texture() TextureFactory

	NewVAO(VBO VBO, EBO EBO) VAO
	// NewVBO[Vertex any](configure func()) VBOSetter[Vertex]
	NewEBO() EBO

	//  NewBuffer[Stored comparable](
	// 	target uint32, // gl.SHADER_STORAGE_BUFFER / gl.DRAW_INDIRECT_BUFFER
	// 	usage uint32, // gl.STATIC_DRAW / gl.DYNAMIC_DRAW
	// 	index uint32,
	// ) Buffer[Stored]

	NewProgram(p uint32, parameters []Parameter) (Program, error)
	NewShader(shaderSource string, shaderType uint32) (Shader, error)
}

// methods outside lean Service abstraction because of golang underdeveloped generics {
func NewVBO[Vertex any](configure func()) VBOSetter[Vertex] {
	return vbo.NewVBO[Vertex](configure)
}
func NewBuffer[Stored comparable](
	target uint32, // gl.SHADER_STORAGE_BUFFER / gl.DRAW_INDIRECT_BUFFER
	usage uint32, // gl.STATIC_DRAW / gl.DYNAMIC_DRAW
	index uint32,
) Buffer[Stored] {
	return buffers.NewBuffer[Stored](target, usage, index)
}

// }

type TextureArrayFactory interface {
	New(datastructures.SparseArray[uint32, image.Image]) (TextureArray, error)
	NewFromSlice([]image.Image) (TextureArray, error)
	Wrap(func(TextureArray))
}
type TextureFactory interface {
	New(image.Image) (Texture, error)
	Wrap(func(Texture))
}

//

var (
	ErrProgramHasOtherLocations error = errors.New("invalid program locations type")

	ErrNotALocation    error = errors.New("expected 'int32' for location")
	ErrInvalidLocation error = errors.New("invalid location")
)

type Parameter struct {
	Name  uint32
	Value int32
}

type Program interface {
	ID() uint32
	Locations(reflect.Type) (any, error)
	Bind()
	Release()
}

func GetProgramLocations[Locations any](p Program) (Locations, error) {
	locations, err := p.Locations(reflect.TypeFor[Locations]())
	if err != nil {
		var l Locations
		return l, err
	}
	return locations.(Locations), nil
}

//

type Buffer[Stored comparable] interface {
	ID() uint32
	Bind()
	Get() []Stored
	Add(elements ...Stored)
	Set(index int, e Stored)
	Remove(indices ...int)
	Release()

	Flush()
}

//

var (
	VertexShader   uint32 = gl.VERTEX_SHADER
	GeomShader     uint32 = gl.GEOMETRY_SHADER
	FragmentShader uint32 = gl.FRAGMENT_SHADER
	ComputeShader  uint32 = gl.COMPUTE_SHADER
)

type Shader interface {
	ID() uint32
	Release()
}

//

var (
	ErrTexturesHaveToShareSize error = errors.New("all textures have to match size")
)

type Image interface {
	Image() image.Image

	FlipH() Image
	FlipV() Image
	// horizontally and vertically
	FlipHV() Image

	// rotates 90 deg clockwise
	RotateClockwise(times int) Image

	TrimTransparentBackground() Image
	Scale(w, h int) Image
	Opaque() Image
}

type Texture interface {
	ID() uint32
	Bind()
	Release()
}
type TextureArray interface {
	Texture() uint32
	ImagesCount() int
	Bind()
	Release()
}

//

type VAO interface {
	ID() uint32
	VBO() VBO
	EBO() EBO
	Release()
	Bind()
}

type VBOFactory[Vertex any] func() VBOSetter[Vertex]
type VBO interface {
	ID() uint32
	Len() int
	Configure()
	Release()
}
type VBOSetter[Vertex any] interface {
	VBO
	SetVertices(vertices []Vertex)
}

type Index uint32
type EBO interface {
	ID() uint32
	Len() int
	Configure()
	Release()
	SetIndices(indices []Index)
}
