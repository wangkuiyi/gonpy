package gonpy

import (
	"log"
	"sort"
)

type Column struct {
	*Matrix
	Col int
}

func NewColumn(mat *Matrix, col int) *Column {
	if col < 0 || col >= mat.Shape.Col {
		log.Panicf("col (%d) out of range [0, %d)", col, mat.Shape.Col)
	}

	return &Column{
		Matrix: mat,
		Col:    col,
	}
}

func (c *Column) At(row int) float64 {
	return c.Data[c.Shape.Col*row+c.Col]
}

func (c *Column) Len() int {
	return c.Shape.Row
}

// Returns a map from Id to rank.
func (c *Column) Rank() map[int]int {
	rkr := newColumnRanker(c)
	sort.Stable(rkr)

	m := make(map[int]int)
	for rk, row := range rkr.order {
		m[row] = rk
	}
	return m
}

func (c *Column) Order() []int {
	rkr := newColumnRanker(c)
	sort.Stable(rkr)
	return rkr.order
}

type columnRanker struct {
	*Column
	order []int
}

func newColumnRanker(c *Column) *columnRanker {
	return &columnRanker{
		Column: c,
		order:  interval(0, c.Len()),
	}
}

func interval(left, right int) []int {
	r := make([]int, right-left)
	for i := left; i < right; i++ {
		r[i-left] = i
	}
	return r
}

func (rkr *columnRanker) Less(i, j int) bool {
	return rkr.At(rkr.order[i]) < rkr.At(rkr.order[j])
}

func (rkr *columnRanker) Swap(i, j int) {
	rkr.order[i], rkr.order[j] = rkr.order[j], rkr.order[i]
}
