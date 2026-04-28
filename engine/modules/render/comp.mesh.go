package render

import (
	"engine/modules/graphics"
	"engine/services/ecs"
)

type MeshComponent struct {
	ID ecs.EntityID
}

func NewMesh(id ecs.EntityID) MeshComponent {
	return MeshComponent{ID: id}
}

//

type Vertex struct {
	Pos [3]float32
	// normal [3]float32
	TexturePos [2]float32
	// color [4]float32
	// vertexGroups (for animation) []VertexGroupWeight {Name string; weight float32} (weights should add up to one)
}

//

type MeshAsset interface {
	Vertices() []Vertex
	Indices() []graphics.Index
	Release()
}

type meshAsset struct {
	vertices []Vertex
	indices  []graphics.Index
}

func NewMeshAsset(
	vertices []Vertex,
	indices []graphics.Index,
) MeshAsset {
	asset := &meshAsset{
		vertices: vertices,
		indices:  indices,
	}
	return asset
}

func (asset *meshAsset) Vertices() []Vertex        { return asset.vertices }
func (asset *meshAsset) Indices() []graphics.Index { return asset.indices }
func (a *meshAsset) Release()                      {}
