package main

import "fmt"

// 指针相当于别名，访问或设置某个值时，使用*p，访问地址使用p
func main() {
	// var i int
	i := 7
	// var p *int
	p := &i

	fmt.Println("i:", i)
	fmt.Println("p:", p)
	fmt.Println("*p:", *p)

	*p = 27
	fmt.Println("*p:", *p)
	fmt.Println("i:", i)

	i = 37
	fmt.Println("*p:", *p)
	fmt.Println("i:", i)

}
