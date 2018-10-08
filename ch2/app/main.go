package main

import (
	"fmt"

	"gopl.io/ch2/popcount"
)

func main() {
	var x uint64 = 1537
	a := popcount.PopCount(x)
	fmt.Println(a)

	b := popcount.PopCountEx23(x)
	fmt.Println(b)
}
