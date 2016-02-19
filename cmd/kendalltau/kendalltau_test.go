package main

import (
	"bytes"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/topicai/candy"
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
	assert.Equal(t, 1, KendallTau(rank1, rank2))
	assert.Equal(t, 1, KendallTau(rank2, rank1))

	rank2 = map[int]int{
		100: 2,
		200: 1,
		300: 0,
	}
	assert.Equal(t, 3, KendallTau(rank1, rank2))
	assert.Equal(t, 3, KendallTau(rank2, rank1))
}

func TestKendalTau(t *testing.T) {
	tau, row := KendallTauMatrix(
		path.Join(GoPath(), "src/github.com/wangkuiyi/gonpy/testdata/2d.npy"),
		500)
	assert.Equal(t, []int{0, 0, 0, 0, 0, 0, 0, 0, 0}, tau)
	assert.Equal(t, 2, row)

	var buf bytes.Buffer
	EncodeKendallTauMatrix(&buf, tau, row)
	assert.Equal(t, "0,0,0\n0,0,0\n0,0,0\n", buf.String())
}
