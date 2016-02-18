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
