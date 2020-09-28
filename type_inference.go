package main

import "fmt"

// 类型必须是相同的才能运算
func main() {
	i := 12.5
	v := 42 + 0.5 + 0.5i
	w := 0.555 * i
	fmt.Printf("v is of the type %T\n", i)
	fmt.Printf("v is of the type %T value %v\n", v, v)
	fmt.Printf("v is of the type %T value %v\n", w, w)
}
