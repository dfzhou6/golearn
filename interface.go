package main

import "fmt"

type I interface {
	printHello()
	printHi()
}

type People struct {
	name string
}

func (p *People)printHello() {
	fmt.Printf("hello my name is %v\n", p.name)
}

func (p People)printHi() {
	fmt.Printf("hi my name is %v\n", p.name)
}

func main() {
	var i I = &People{ name: "zdf" }
	i.printHello()
	i.printHi()
}

