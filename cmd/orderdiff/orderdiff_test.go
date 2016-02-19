package main

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/topicai/candy"
)

func TestOrderDiff(t *testing.T) {
	assert.Equal(t,
		[]int{0, 0, 0},
		orderDiff(path.Join(
			candy.GoPath(),
			"src/github.com/wangkuiyi/gonpy/testdata/2d.npy")))
}
