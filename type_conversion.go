package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 12.5
	f := uint(i)
	// 数字转换为字符串
	s := strconv.Itoa(int(f))
	fmt.Printf("Type %T Value %v\n", i, i)
	fmt.Printf("Type %T Value %v\n", f, f)
	fmt.Printf("Type %T Value %v\n", s, s)
}
