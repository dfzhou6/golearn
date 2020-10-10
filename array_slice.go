package main

import "fmt"

func printSlice(array []int) {
	fmt.Printf("len: %v, cap: %v, value: %v\n", len(array), cap(array), array)
}

// 这里要注意，切片从1开始切，会改变其cap和本身，从0则不会
func main() {
	//s := []int{2, 3, 5, 7, 11, 13}
	var s []int = []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[1:4]
	printSlice(s)
	// 3,5,7 len:3 cap:5

	s = s[:2]
	printSlice(s)
	// 3,5 len:2 cap:5

	s = s[1:]
	printSlice(s)
	// 5 len:1 cap:4

	v := s
	printSlice(v)

	// 切片就像数组的引用
	v[0] = 10000
	printSlice(v)
	printSlice(s)
}
