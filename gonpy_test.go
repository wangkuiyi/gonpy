package gonpy

import (
	"bufio"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/topicai/candy"
)

var (
	mat2d, mat1d *Matrix
)

func init() {
	mat1d = WithOpened(TestData("1d.npy"), func(r io.Reader) interface{} {
		m, e := Load(bufio.NewReader(r))
		Must(e)
		return m
	}).(*Matrix)

	mat2d = WithOpened(TestData("2d.npy"), func(r io.Reader) interface{} {
		m, e := Load(bufio.NewReader(r))
		Must(e)
		return m
	}).(*Matrix)
}

func TestLoad1DInt64File(t *testing.T) {
	assert.Equal(t, 10, mat1d.Row)
	assert.Equal(t, 1, mat1d.Col)
	for i := 0; i < mat1d.Row; i++ {
		assert.Equal(t, float64(i), mat1d.Data[i])
	}
}

func TestLoad2DFloat64File(t *testing.T) {
	assert.Equal(t, 2, mat2d.Row)
	assert.Equal(t, 3, mat2d.Col)
	for i := 0; i < mat2d.Row*mat2d.Col; i++ {
		assert.Equal(t, float64(i+1), mat2d.Data[i])
	}
}
