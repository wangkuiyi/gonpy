package main

import (
	"flag"
	"fmt"
	"runtime"
)

func main() {
	parallelism := flag.Int("p", 1, "GOMACPROCS value")
	capRows := flag.Int("cap", 500, "No more than this max number of rows of the matrix will be handled")
	flag.Parse()

	runtime.GOMAXPROCS(*parallelism)

	fmt.Println(KendallTauMatrix(flag.Arg(0), *capRows))
}
