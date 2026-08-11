package internal

import (
	"container/heap"
	"core/modules/obstruction"
	"core/modules/pathfind"
	"core/modules/tile"
	"engine/modules/grid"
)

// algorithm is AI generated

type Item struct {
	coords grid.Coords
	gScore int
	fScore int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].fScore == pq[j].fScore {
		// Tie-breaker: favor nodes closer to the goal (lower hScore / higher gScore)
		return pq[i].gScore > pq[j].gScore
	}
	return pq[i].fScore < pq[j].fScore
}
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x any)   { *pq = append(*pq, x.(*Item)) }
func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func heuristic(a, b grid.Coords) int {
	return abs(int(a.X-b.X)) + abs(int(a.Y-b.Y))
}

func (s *service) findPath(
	from, to grid.Coords,
	size tile.SizeComponent,
	obstruction obstruction.Component,
) (path []grid.Coords, ok bool) {
	r1, _ := s.CoordsRegion(from, obstruction.Obstruction)
	r2, _ := s.CoordsRegion(to, obstruction.Obstruction)
	if r1 != r2 {
		return nil, false
	}

	toData, ok := s.Obstruction().Grid().CoordsData(to)
	if !ok || obstruction.Obstruction&toData.Component.GetTile(toData.Index) != 0 {
		return nil, false
	}

	gScore := map[grid.Coords]int{from: 0}
	cameFrom := map[grid.Coords]grid.Coords{}

	openSet := &PriorityQueue{}
	heap.Init(openSet)

	hStart := heuristic(from, to)
	heap.Push(openSet, &Item{coords: from, gScore: 0, fScore: hStart})

	dirs := []struct{ x, y grid.Coord }{
		{0, 1}, {1, 0}, {0, grid.NewCoord(-1)}, {grid.NewCoord(-1), 0},
	}

	for openSet.Len() > 0 {
		current := heap.Pop(openSet).(*Item)

		// Skip stale entries that were re-queued with better gScores
		if current.gScore > gScore[current.coords] {
			continue
		}

		if current.coords == to {
			return reconstructPath(cameFrom, from, to), true
		}

		for _, d := range dirs {
			neighborCoords := grid.Coords{
				X: current.coords.X + d.x,
				Y: current.coords.Y + d.y,
			}

			step := pathfind.NewStep(neighborCoords.X, neighborCoords.Y)
			if !s.Pathfind().CanStep(current.coords, size, obstruction, step) {
				continue
			}

			tentativeGScore := current.gScore + 1

			bestG, exists := gScore[neighborCoords]
			if !exists || tentativeGScore < bestG {
				cameFrom[neighborCoords] = current.coords
				gScore[neighborCoords] = tentativeGScore

				fScore := tentativeGScore + heuristic(neighborCoords, to)
				heap.Push(openSet, &Item{
					coords: neighborCoords,
					gScore: tentativeGScore,
					fScore: fScore,
				})
			}
		}
	}

	return nil, false
}

func reconstructPath(
	cameFrom map[grid.Coords]grid.Coords,
	from, to grid.Coords,
) []grid.Coords {
	var path []grid.Coords
	curr := to

	for curr != from {
		path = append(path, curr)
		curr = cameFrom[curr]
	}

	// Reverse slice in place to get path from start to end (excluding start)
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}
