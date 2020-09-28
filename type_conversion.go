package main

import "fmt"

func main() {
	i := 12.5
	f := uint(i)
	fmt.Printf("Type %T Value %v\n", i, i)
	fmt.Printf("Type %T Value %v\n", f, f)
}
