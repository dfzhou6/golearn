package main

import "fmt"

func closure(sum int) func(i int) int {
	return func (i int) int {
		sum += i
		return sum
	}
}

func main() {
	var v int

	fn1 := closure(10)
	for i := 0; i < 10; i++ {
		v = fn1(i)
		fmt.Println(v)
	}
}