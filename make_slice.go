package main

import "fmt"

func Printf(arr []int) {
	fmt.Printf("len: %v, cap: %v, value: %v\n", len(arr), cap(arr), arr)
}

func main() {
	var arr []int
	arr = make([]int, 2, 5)
	Printf(arr)
	
}
