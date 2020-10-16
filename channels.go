package main

import (
	"fmt"
	"time"
)

func test(ch chan int, b bool) {
	n := 10
	if b == true {
		fmt.Println("start sleep")
		time.Sleep(1000 * time.Millisecond)
		n = 20
	}
	ch <- n
}

func main() {
	ch := make(chan int)
	go test(ch, true)
	go test(ch, true)
	go test(ch, false)
	a := <- ch
	fmt.Printf("a:%v\n", a)
	b := <- ch
	fmt.Printf("b:%v \n", b)
	c := <- ch
	fmt.Printf("c:%v \n", c)
}
