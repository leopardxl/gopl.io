package main

import "fmt"

func main() {
	x, y := 1, 1

	p := &x
	fmt.Printf("var x has a value of %v and lives at %v in memory\n", *p, p)
	*p = 24
	fmt.Printf("var x now has a value of %v and lives at %v in memory\n", *p, p)

	fmt.Println(&x == &x, &x == &y, &x == nil) // "True false false"

	fmt.Println(f() == f()) // False

	v := 1
	incr(&v)              // side effect: v is now 2
	fmt.Println(incr(&v)) // "3" and v is 3

	u := 1
	localIncr(u)              // side effect: v is now 2
	fmt.Println(localIncr(u)) // "3" and v is 3

	a := new(int)
	fmt.Println(*a)
	*a = 2
	fmt.Println(*a)

	medal := []string{"gold", "silver", "bronze"}
	fmt.Println(medal)
}

func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++ // increments what p points to; does not change p
	return *p
}

func localIncr(p int) int {
	p++
	return p
}
