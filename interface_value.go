package main

import "fmt"

type It interface {
	printHello()
	changeName()
}

type Student3 struct {
	name string
}

func (s *Student3)printHello() {
	if s == nil {
		fmt.Println("s is <nil>")
	} else {
		fmt.Printf("Hello my name is %v\n", s.name)
	}
}

func (s *Student3)changeName() {
	s.name = s.name + "_new"
}

func desc(it It) {
	fmt.Printf("it value:%v type:%T\n", it, it)
}

func main() {
	var it It

	var s3 *Student3
	it = s3
	desc(it)
	it.printHello()

	var s4 = &Student3{ name: "zdf" }
	it = s4
	desc(it)
	it.printHello()
	it.changeName()
	it.printHello()
	s4.printHello()
}
