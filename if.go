package main

import (
	"fmt"
	"math"
)

func sqrtTest(x float64) string {
	if x < 0 {
		return "less than 0\n"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func ifStatement(v float64) float64 {
	// vv的变量作用域仅在if-else内
	if vv := math.Pow(2, v); vv > 100 {
		fmt.Printf("vv: %v > 100\n", vv)
	} else {
		fmt.Printf("vv: %v <= 100\n", vv)
	}
	return math.Pow(2, v)
}

func main() {
	fmt.Println(sqrtTest(-2), sqrtTest(3))
	fmt.Println(ifStatement(8))
}

