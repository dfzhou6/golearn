package main

import "fmt"

type ErrTest int

// 优先打印Error
func (e ErrTest) Error() string {
	return fmt.Sprintf("cannot be negative number: %v", float64(e))
}

// 没有Error才打印String
func (e ErrTest) String() string {
	return fmt.Sprintf("hello print: %v", float64(e))
}

func main() {
	fmt.Println(ErrTest(120))
}
