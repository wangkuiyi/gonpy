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
	"github.com/wangkuiyi/parallel"
)

func KendallTau(rank1, rank2 map[int]int) int {
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

	var tau int

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

// The KendallTau algorithm is O(N^2) with respect to the length of
// its two rank operands.  If they are too long, it would take too
// much time to compute.  So we handle the first cap rows of the
// matrix in filename.
func KendallTauMatrix(filename string, cap int) []int {
	var mat *gonpy.Matrix
	progress(func() {
		mat = candy.WithOpened(filename, func(r io.Reader) interface{} {
			m, e := gonpy.Load(bufio.NewReader(r))
			candy.Must(e)
			return m
		}).(*gonpy.Matrix)
	},
		"Loading matrix %s", filename)

	progress(func() {
		m := mat.Shape.Row
		if m > cap {
			m = cap
		}
		mat = mat.Slice(0, m)
	},
		"Select only the first %d rows", cap)

	ranks := make([]map[int]int, mat.Shape.Col*mat.Shape.Col)
	progress(func() {
		parallel.For(0, mat.Shape.Col, 1, func(col int) {
			ranks[col] = gonpy.NewColumn(mat, col).Rank()
		})
	},
		"Rank columns")

	ret := make([]int, mat.Shape.Col*mat.Shape.Col)
	progress(func() {
		parallel.For(0, mat.Shape.Col-1, 1, func(col1 int) {
			parallel.For(col1+1, mat.Shape.Col, 1, func(col2 int) {
				tau := KendallTau(ranks[col1], ranks[col2])
				ret[col1*mat.Shape.Col+col2] = tau
				ret[col2*mat.Shape.Col+col1] = tau
			})
		})
	},
		"Kendall'Tau of all column pairs")

	return ret
}

func progress(fn func(), format string, args ...interface{}) {
	start := time.Now()
	msg := fmt.Sprintf(format+" ... ", args...)
	log.Print(msg)
	fn()
	log.Printf("%s Done in %v", msg, time.Since(start))
}
