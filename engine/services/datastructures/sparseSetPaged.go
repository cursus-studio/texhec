package datastructures

import (
	"slices"

	"golang.org/x/exp/constraints"
)

type sparseSetWithPaging[Index constraints.Integer] struct {
	PageSize   Index
	EmptyIndex Index
	Pages      [][]Index // here some indices have special meaning (read constants above)

	// both arrays below Dense correspond
	Dense []Index // here value means index in sparse array
}

func NewSparseSetWithPaging[Index constraints.Integer]() SparseSet[Index] {
	var zero Index
	pageSize := 4096
	return &sparseSetWithPaging[Index]{
		PageSize:   max(Index(pageSize), 64),
		EmptyIndex: ^zero,
	}
}

func (a *sparseSetWithPaging[Index]) Get(index Index) bool {
	pageIndex := index / a.PageSize
	if int(pageIndex) >= len(a.Pages) {
		return false
	}

	page := a.Pages[pageIndex]
	if page == nil {
		return false
	}
	valueIndex := page[index%a.PageSize]
	return valueIndex != a.EmptyIndex
}

func (a *sparseSetWithPaging[Index]) GetIndices() []Index { return a.Dense }

func (a *sparseSetWithPaging[Index]) Add(index Index) bool {
	pageIndex := index / a.PageSize
	if diff := int(pageIndex) + 1 - len(a.Pages); diff > 0 {
		a.Pages = slices.Grow(a.Pages, diff)[:int(pageIndex)+1]
	}

	page := a.Pages[pageIndex]
	if page == nil {
		page = make([]Index, a.PageSize)
		for i := range page {
			page[i] = a.EmptyIndex
		}
		a.Pages[pageIndex] = page
	}
	valueIndex := index % a.PageSize
	denseIndex := page[valueIndex]

	if denseIndex == a.EmptyIndex {
		page[valueIndex] = Index(len(a.Dense))
		a.Dense = append(a.Dense, index)
		return true
	}

	a.Dense[denseIndex] = index

	return false
}

func (a *sparseSetWithPaging[Index]) Remove(index Index) bool {
	pageIndex := index / a.PageSize
	if int(pageIndex) >= len(a.Pages) {
		return false
	}

	page := a.Pages[pageIndex]
	if page == nil {
		return false
	}
	valueIndex := index % a.PageSize
	value := page[valueIndex]
	if page[valueIndex] == a.EmptyIndex {
		return false
	}

	page[valueIndex] = a.EmptyIndex

	if len(a.Dense)-1 != int(value) {
		movedIndex := a.Dense[len(a.Dense)-1]
		a.Dense[value] = movedIndex

		a.Pages[movedIndex/a.PageSize][movedIndex%a.PageSize] = value
	}

	a.Dense = a.Dense[:len(a.Dense)-1]
	return true
}
