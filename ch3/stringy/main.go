package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "안녕하세요"
	length := len(s)

	for x, str := range s {
		fmt.Printf("%d %v\n", x, str)
	}
	fmt.Println("\uC548 ", length)
	fmt.Println("Runes:", utf8.RuneCountInString(s))

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
	//or

	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

}
