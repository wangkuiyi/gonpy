package main

import (
	"flag"
	"fmt"
	"runtime"
)

func main() {
	parallelism := flag.Int("p", 1, "GOMACPROCS value")
	flag.Parse()

	runtime.GOMAXPROCS(*parallelism)
	fmt.Println(KendallTauMatrix(flag.Arg(0), *parallelism))
}
