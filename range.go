package main

import "fmt"

func main() {
	var arr []int = make([]int, 10, 10)

	for i, v := range arr {
		v = (v+1)*i
		fmt.Printf("i:%v, v:%v\n", i, v)
	}

	for i := range arr {
		fmt.Printf("i:%v\n", i)
	}

	for _, v := range arr {
		fmt.Printf("v:%v\n", v)
	}

	arr = append(arr, 1000, 20000)
	fmt.Printf("arr:%v\n", arr)

	var t_arr [2]int = [2]int{1,2}
	fmt.Printf("t_arr:%v\n", t_arr)
	// 数组的数量固定则不可再扩展元素
	// t_arr = append(t_arr, 3)
	// fmt.Printf("t_arr:%v\n", t_arr)

}
