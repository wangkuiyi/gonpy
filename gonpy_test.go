package gonpy

import (
	"bufio"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/topicai/candy"
)

func TestLoad1DInt64File(t *testing.T) {
	m := WithOpened(TestData("1d.npy"), func(r io.Reader) interface{} {
		m, e := Load(bufio.NewReader(r))
		Must(e)
		return m
	}).(*Matrix)

	assert.Equal(t, 10, m.Row)
	assert.Equal(t, 1, m.Col)
	for i := 0; i < m.Row; i++ {
		assert.Equal(t, float64(i), m.Data[i])
	}
}

func TestLoad2DFloat64File(t *testing.T) {
	m := WithOpened(TestData("2d.npy"), func(r io.Reader) interface{} {
		m, e := Load(bufio.NewReader(r))
		Must(e)
		return m
	}).(*Matrix)

	assert.Equal(t, 2, m.Row)
	assert.Equal(t, 3, m.Col)
	for i := 0; i < m.Row*m.Col; i++ {
		assert.Equal(t, float64(i+1), m.Data[i])
	}
}
