package gonpy

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wangkuiyi/gonpy/header"
)

func TestNewColumn(t *testing.T) {
	c := NewColumn(mat2d, 0)
	assert.Equal(t, 1.0, c.At(0))
	assert.Equal(t, 4.0, c.At(1))

	c = NewColumn(mat2d, 1)
	assert.Equal(t, 2.0, c.At(0))
	assert.Equal(t, 5.0, c.At(1))

	c = NewColumn(mat2d, 2)
	assert.Equal(t, 3.0, c.At(0))
	assert.Equal(t, 6.0, c.At(1))

	c = NewColumn(mat1d, 0)
	assert.Equal(t, 10, c.Len())
	for i := 0; i < 10; i++ {
		assert.Equal(t, float64(i), c.At(i))
	}
}

func TestColumnRank(t *testing.T) {
	m := &Matrix{
		Shape: &header.Shape{Row: 10, Col: 1},
		Data:  make([]float64, 10),
	}

	for i := 0; i < 10; i++ {
		m.Data[i] = math.Sin(float64(i) / 10.0 * math.Pi * 2.0)
	}

	r := NewColumn(m, 0).Rank()

	expected := map[int]int{
		4: 7,
		2: 8,
		3: 9,
		9: 2,
		6: 3,
		0: 4,
		1: 6,
		8: 0,
		7: 1,
		5: 5,
	}
	assert.Equal(t, r, expected)
}

func TestColumnOrder(t *testing.T) {
	m := &Matrix{
		Shape: &header.Shape{Row: 10, Col: 1},
		Data:  make([]float64, 10),
	}

	for i := 0; i < 10; i++ {
		m.Data[i] = math.Sin(float64(i) / 10.0 * math.Pi * 2.0)
	}

	r := NewColumn(m, 0).Order()
	assert.Equal(t, []int{8, 7, 9, 6, 0, 5, 1, 4, 2, 3}, r)
}
