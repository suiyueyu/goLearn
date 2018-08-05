package main

import (
	"fmt"

	"github.com/suiyueyu/goLearn/src/main/go/ch2/popcount"
)

func main() {
	var x uint64 = 63
	fmt.Println(popcount.PopCount(x))
	fmt.Println(popcount.PopCount(x) == popcount.PopCount2(x))
}
