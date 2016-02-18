package gonpy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKendallTau(t *testing.T) {
	rank1 := map[int]int{
		100: 0,
		200: 1,
		300: 2,
	}
	rank2 := map[int]int{
		100: 0,
		200: 2,
		300: 1,
	}
	assert.Equal(t, int64(1), KendallTau(rank1, rank2))
	assert.Equal(t, int64(1), KendallTau(rank2, rank1))

	rank2 = map[int]int{
		100: 2,
		200: 1,
		300: 0,
	}
	assert.Equal(t, int64(3), KendallTau(rank1, rank2))
	assert.Equal(t, int64(3), KendallTau(rank2, rank1))
}

func TestKendalTau(t *testing.T) {
	columns := make([]*Column, mat2d.Shape.Col)

	for col := 0; col < mat2d.Shape.Col; col++ {
		columns[col] = NewColumn(mat2d, col)
	}

	assert.Equal(t, int64(0), KendallTau(columns[0].Rank(), columns[0].Rank()))
	assert.Equal(t, int64(0), KendallTau(columns[0].Rank(), columns[1].Rank()))
	assert.Equal(t, int64(0), KendallTau(columns[0].Rank(), columns[2].Rank()))
}
