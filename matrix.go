package gonpy

import "github.com/wangkuiyi/gonpy/header"

type Matrix struct {
	*header.Shape
	Data []float64
}
