package main

import (
	"bufio"
	"io"
	"log"
	"reflect"
	"sort"

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
	log.Printf("Loading matrix %s ..", filename)
	mat := candy.WithOpened(filename, func(r io.Reader) interface{} {
		m, e := gonpy.Load(bufio.NewReader(r))
		candy.Must(e)
		return m
	}).(*gonpy.Matrix)

	log.Printf("Computing baseline rank ...")
	baseline := gonpy.NewColumn(mat, 0).Rank()

	ret := make([]int64, mat.Shape.Col)
	for col := 1; col < mat.Shape.Col; col++ {
		log.Printf("Rank column %d ...", col)
		r := gonpy.NewColumn(mat, col).Rank()

		log.Printf("Kendall'Tau of column %d ...", col)
		ret[col] = KendallTau(baseline, r)
	}
	return ret
}
