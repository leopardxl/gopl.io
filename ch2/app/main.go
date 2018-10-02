package main

import (
	"fmt"

	"gopl.io/ch2/popcount"
)

func main() {
	a := popcount.PopCount(2345678)
	fmt.Println(a)
}
