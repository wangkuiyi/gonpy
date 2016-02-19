package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"reflect"
	"sort"
	"sync/atomic"
	"time"

	"github.com/topicai/candy"
	"github.com/wangkuiyi/gonpy"
	"github.com/wangkuiyi/parallel"
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

// High performance version of KendallTau which does less valid
// parameter checking and using mutliple goroutines.
func KendallTauPerf(rank1, rank2 map[int]int, parallelism int) int64 {
	if parallelism < 1 {
		log.Panicf("parallelism (%d) must be >= 1", parallelism)
	}

	if len(rank1) != len(rank2) {
		log.Panicf("kendall's Tau is a distance only of two ranks are the same.")
	}

	ids := make([]int, 0, len(rank1))
	for k := range rank1 {
		ids = append(ids, k)
	}

	var tau int64

	type Idxs struct {
		idx1, idx2 int
	}

	ch := make(chan Idxs)

	go parallel.For(0, parallelism, 1, func(i int) {
		for idxs := range ch {
			i := idxs.idx1
			j := idxs.idx2

			if rank1[ids[i]] < rank1[ids[j]] && rank2[ids[i]] > rank2[ids[j]] {
				atomic.AddInt64(&tau, 1)
			}
			if rank1[ids[i]] > rank1[ids[j]] && rank2[ids[i]] < rank2[ids[j]] {
				atomic.AddInt64(&tau, 1)
			}
		}
	})

	for i := 0; i < len(ids)-1; i++ {
		for j := i + 1; j < len(ids); j++ {
			ch <- Idxs{i, j}
		}
	}
	close(ch)

	return tau
}

func KendallTauMatrix(filename string, parallelism int) []int64 {
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
			ret[col] = KendallTauPerf(baseline, r, parallelism)
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
