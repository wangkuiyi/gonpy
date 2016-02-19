package gonpy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlice1DMatrix(t *testing.T) {
	for row := 0; row < mat1d.Shape.Row; row++ {
		m := mat1d.Slice(row, row+1)
		assert.Equal(t, 1, m.Shape.Row)
		assert.Equal(t, 1, m.Shape.Col)
		assert.Equal(t, float64(row), m.Data[0])
	}
}

func TestSlice2DMatrix(t *testing.T) {
	for row := 0; row < mat2d.Shape.Row; row++ {
		m := mat2d.Slice(0, row+1)
		assert.Equal(t, row+1, m.Shape.Row)
		assert.Equal(t, mat2d.Shape.Col, m.Shape.Col)
		assert.Equal(t, 1.0, m.Data[0])
	}
}
