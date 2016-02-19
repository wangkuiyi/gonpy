package gonpy

import (
	"log"

	"github.com/wangkuiyi/gonpy/header"
)

func (m *Matrix) Slice(row1, row2 int) *Matrix {
	if row1 >= row2 {
		log.Panicf("row1 (%d) must be less than row2 (%d)", row1, row2)
	}
	if row1 < 0 || row1 >= m.Shape.Row {
		log.Panicf("row1 (%d) out of range [0, %d)", row1, m.Shape.Row)
	}
	if row2 <= 0 || row2 > m.Shape.Row {
		log.Panicf("row2 (%d) out of range [0, %d)", row2, m.Shape.Row)
	}

	return &Matrix{
		Shape: &header.Shape{Row: row2 - row1, Col: m.Shape.Col},
		Data:  m.Data[row1*m.Shape.Col : row2*m.Shape.Col],
	}
}
