package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(KendallTauMatrix(os.Args[1]))
}
