package main

import "fmt"

func main() {
	var arr []string = []string{"a", "b", "c"}
	fmt.Printf("len: %v, cap: %v, value: %v\n", len(arr), cap(arr), arr)
	arr = append(arr, "d", "e")
	fmt.Printf("len: %v, cap: %v, value: %v\n", len(arr), cap(arr), arr)
}

