package main

import "fmt"

// defer调用的函数会等到外层函数执行返回后再执行，但参数会立刻求值，执行顺序后进先出
func main() {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	fmt.Println("right noe")
}
