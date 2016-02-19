package main

import (
	"flag"
	"os"
	"runtime"
)

func main() {
	parallelism := flag.Int("p", 1, "GOMACPROCS value")
	capRows := flag.Int("cap", 500, "No more than this max number of rows of the matrix will be handled")
	flag.Parse()

	runtime.GOMAXPROCS(*parallelism)

	tau, row := KendallTauMatrix(flag.Arg(0), *capRows)
	EncodeKendallTauMatrix(os.Stdout, tau, row)
}
