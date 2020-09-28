package main

import "fmt"

func main() {
	a := 0
	for i := 0; i < 10; i++ {
		a += i
	}
	fmt.Println(a)

	b := 0
	for ; b < 100; {
		b += 10
	}
	fmt.Println(b)

	// for是go的while
	c := 0
	for c < 10000000 {
		c += 1
	}
	fmt.Println(c)

}
