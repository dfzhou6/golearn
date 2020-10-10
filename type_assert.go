package main

import "fmt"

func main() {
	var i interface{}
	fmt.Printf("i value:%v, type:%T\n", i, i)

	i = 27
	fmt.Printf("i value:%v, type:%T\n", i, i)

	i = "hello"
	fmt.Printf("i value:%v, type:%T\n", i, i)

	var ii interface{} = "hi"
	s, ok := ii.(string)
	fmt.Printf("s:%v, ok:%v\n", s, ok)

	ss, ok := ii.(float32)
	fmt.Printf("ss:%v, ok:%v\n", ss, ok)

	switch v := ii.(type) {
	case int32:
		fmt.Println("int32")
	case string:
		fmt.Println("string")
	default:
		fmt.Printf("other type %T\n", v)
	}

}
