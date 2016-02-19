package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"reflect"
	"sort"
	"time"

	"github.com/topicai/candy"
	"github.com/wangkuiyi/gonpy"
)

func KendallTau(rank1, rank2 map[int]int) int64 {
	ids1 := make([]int, 0, len(rank1))
	for k := range rank1 {
		ids1 = append(ids1, k)
	}
	sort.Ints(ids1)

	ids2 := make([]int, 0, len(rank2))
	for k := range rank2 {
		ids2 = append(ids2, k)
	}
	sort.Ints(ids2)

	if !reflect.DeepEqual(ids1, ids2) {
		log.Panicf("kendall's Tau is a distance only of two ranks are the same.")
	}

	var tau int64

	ids := ids1
	for i := 0; i < len(ids)-1; i++ {
		for j := i + 1; j < len(ids); j++ {
			if rank1[ids[i]] < rank1[ids[j]] && rank2[ids[i]] > rank2[ids[j]] {
				tau = tau + 1.0
			}
			if rank1[ids[i]] > rank1[ids[j]] && rank2[ids[i]] < rank2[ids[j]] {
				tau = tau + 1.0
			}
		}
	}

	return tau
}

func KendallTauMatrix(filename string) []int64 {
	var mat *gonpy.Matrix
	var baseline map[int]int
	var r map[int]int

	progress(func() {
		mat = candy.WithOpened(filename, func(r io.Reader) interface{} {
			m, e := gonpy.Load(bufio.NewReader(r))
			candy.Must(e)
			return m
		}).(*gonpy.Matrix)
	},
		"Loading matrix %s", filename)

	cap := 500
	progress(func() {
		m := mat.Shape.Row
		if m > cap {
			m = cap
		}
		mat = mat.Slice(0, m)
	},
		"Select only the first %d instances", cap)

	progress(func() {
		baseline = gonpy.NewColumn(mat, 0).Rank()
	},
		"Computing baseline rank")

	ret := make([]int64, mat.Shape.Col)
	for col := 1; col < mat.Shape.Col; col++ {
		progress(func() {
			r = gonpy.NewColumn(mat, col).Rank()
		},
			"Rank column %d", col)

		progress(func() {
			ret[col] = KendallTau(baseline, r)
		},
			"Kendall'Tau of column %d", col)

	}
	return ret
}

func progress(fn func(), format string, args ...interface{}) {
	start := time.Now()
	msg := fmt.Sprintf(format+" ... ", args...)
	log.Print(msg)
	fn()
	log.Printf("%s Done in %v", msg, time.Since(start))
}
