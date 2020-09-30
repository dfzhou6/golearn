package main

import "fmt"

func plusOne(i int) int {
	i += 1
	return i
}

func compute(fn func(i int) int) string {
	v := fn(3)
	return fmt.Sprintf("v:%v", v)
}

func main() {
	rt := compute(plusOne)
	fmt.Println(rt)

	plusTwo := func(i int) int {
		i += 2
		return i
	}

	rt2 := compute(plusTwo)
	fmt.Println(rt2)
}
