package main

import (
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

func TestKendallTauPerf(t *testing.T) {
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
	assert.Equal(t, int64(1), KendallTauPerf(rank1, rank2, 2))
	assert.Equal(t, int64(1), KendallTauPerf(rank2, rank1, 2))

	rank2 = map[int]int{
		100: 2,
		200: 1,
		300: 0,
	}
	assert.Equal(t, int64(3), KendallTauPerf(rank1, rank2, 4))
	assert.Equal(t, int64(3), KendallTauPerf(rank2, rank1, 4))
}

func TestKendalTau(t *testing.T) {
	assert.Equal(t, []int64{0, 0, 0},
		KendallTauMatrix(
			path.Join(GoPath(), "src/github.com/wangkuiyi/gonpy/testdata/2d.npy"),
			4))
}
