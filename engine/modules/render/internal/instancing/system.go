package instancing

import (
	_ "embed"
	"engine"
	"engine/modules/graphics"
	"engine/modules/render"
	"engine/services/datastructures"
	"engine/services/ecs"

	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/ogiusek/events"
	"github.com/ogiusek/ioc/v2"
)

//go:embed s.vert
var vertSource string

//go:embed s.frag
var fragSource string

type locations struct {
	Camera       int32 `uniform:"camera"`
	CameraGroups int32 `uniform:"cameraGroups"`
}

//

type system struct {
	engine.EngineWorld `inject:""`
	VboFactory         graphics.VBOFactory[render.Vertex] `inject:""`

	// batches
	dirtyEntities   ecs.DirtySet
	entitiesBatches datastructures.SparseArray[ecs.EntityID, batchKey]
	batches         map[batchKey]*batch

	meshes   map[ecs.EntityID]graphics.VAO
	textures map[ecs.EntityID]graphics.TextureArray

	program   graphics.Program
	locations locations
}

func NewSystem(c ioc.Dic) render.SystemRenderer {
	world := ioc.GetServices[engine.EngineWorld](c)
	return ecs.NewSystemRegister(func() error {
		vert, err := world.Graphics().NewShader(vertSource, graphics.VertexShader)
		if err != nil {
			return err
		}
		defer vert.Release()

		frag, err := world.Graphics().NewShader(fragSource, graphics.FragmentShader)
		if err != nil {
			return err
		}
		defer frag.Release()

		programID := gl.CreateProgram()
		gl.AttachShader(programID, vert.ID())
		gl.AttachShader(programID, frag.ID())

		p, err := world.Graphics().NewProgram(programID, nil)
		if err != nil {
			return err
		}

		locations, err := graphics.GetProgramLocations[locations](p)
		if err != nil {
			return err
		}

		s := ioc.GetServices[*system](c)

		s.dirtyEntities = ecs.NewDirtySet()
		s.entitiesBatches = datastructures.NewSparseArray[ecs.EntityID, batchKey]()
		s.batches = make(map[batchKey]*batch)

		s.meshes = make(map[ecs.EntityID]graphics.VAO)
		s.textures = make(map[ecs.EntityID]graphics.TextureArray)

		s.program = p
		s.locations = locations

		s.Render().Color().AddDirtySet(s.dirtyEntities)
		s.Render().TextureFrame().AddDirtySet(s.dirtyEntities)
		s.Transform().AddDirtySet(s.dirtyEntities)
		s.Render().Mesh().AddDirtySet(s.dirtyEntities)
		s.Render().Texture().AddDirtySet(s.dirtyEntities)

		events.ListenE(s.EventsBuilder(), s.ListenRender)
		return nil
	})
}

func (s *system) ListenRender(render render.RenderEvent) error {
	// batch
	// for dirtyEntity in entities
	//  if exists than add (create batch if it doesn't exist)
	//  else remove
	var err error
	for _, entity := range s.dirtyEntities.Get() {
		batchKey, batchKeyOk := batchKey{}, true
		if batchKeyOk {
			batchKey.mesh, batchKeyOk = s.Render().Mesh().Get(entity)
		}
		if batchKeyOk {
			batchKey.texture, batchKeyOk = s.Render().Texture().Get(entity)
		}

		oldBatchKey, oldBatchKeyOk := s.entitiesBatches.Get(entity)
		if oldBatchKeyOk && (!batchKeyOk || batchKey != oldBatchKey) {
			oldBatch := s.batches[oldBatchKey]
			oldBatch.Remove(entity)
			s.entitiesBatches.Remove(entity)
		}
		if !batchKeyOk {
			continue
		}
		batch, ok := s.batches[batchKey]
		if !ok {
			batch, err = s.NewBatch(batchKey)
			if err != nil {
				return err
			}
			s.batches[batchKey] = batch
		}
		batch.Upsert(entity)
		s.entitiesBatches.Set(entity, batchKey)
	}

	// render
	// for batch in batches
	//  bind everything
	//  for camera in cameras
	//   bind camera mat4
	//   render
	s.program.Bind()

	for _, batch := range s.batches {
		camMatrix := s.Camera().Mat4(render.Camera)
		gl.UniformMatrix4fv(s.locations.Camera, 1, false, &camMatrix[0])

		camGroups, _ := s.Groups().Component().Get(render.Camera)
		gl.Uniform1ui(s.locations.CameraGroups, camGroups.Mask)

		batch.Render()
	}
	return nil
}
