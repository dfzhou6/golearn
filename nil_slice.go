package main

import "fmt"

func main() {
	var s []int
	fmt.Println("s:", s)
	if s == nil {
		fmt.Println("nil")
	}

	var s2 [3]int
	fmt.Println("s2:", s2)
}
