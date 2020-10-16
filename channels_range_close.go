package main

import (
	"fmt"
)

func f(n int, c chan int, num int) {
	for i := 0; i < n; i++ {
		c <- num
	}
	//close(c)
	fmt.Printf("ok %v\n", num)
}

func main() {
	c := make(chan int, 2)
	go f(cap(c), c, 1)
	fmt.Println("c")

	fmt.Println(<- c)
	//for i := range c {
	//	fmt.Println(i)
	//}

	d := make(chan int, 2)
	go f(cap(d), d, 2)
	fmt.Println("d")

	fmt.Println(<- d)
	//for i := range d {
	//	fmt.Println(i)
	//}
}
