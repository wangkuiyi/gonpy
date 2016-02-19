package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/topicai/candy"
	"github.com/wangkuiyi/gonpy"
)

func main() {
	fmt.Println(orderDiff(os.Args[1]))
}

func orderDiff(filename string) []int {
	var mat *gonpy.Matrix
	var baseline []int

	progress(func() {
		mat = candy.WithOpened(filename, func(r io.Reader) interface{} {
			m, e := gonpy.Load(bufio.NewReader(r))
			candy.Must(e)
			return m
		}).(*gonpy.Matrix)

		baseline = gonpy.NewColumn(mat, 0).Order()
	},
		"Loading matrix %s and rank the first column as baseline", filename)

	ret := make([]int, mat.Shape.Col)
	for col := 1; col < mat.Shape.Col; col++ {
		progress(func() {
			r := gonpy.NewColumn(mat, col).Order()
			for i, b := range baseline {
				if b != r[i] {
					ret[col]++
				}
			}
		},
			"Ordering column %d and compare with baseline", col)
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
