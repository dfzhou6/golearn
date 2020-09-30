package main

import "fmt"

func Printf(arr []string) {
	fmt.Printf("len: %v, cap: %v, value: %v\n", len(arr), cap(arr), arr)
}

func main() {
	var matrix [][]string = [][]string {
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	matrix[0][1] = "+"
	matrix[1][0] = "+"
	matrix[1][2] = "+"
	matrix[2][1] = "+"

	for i := 0; i < len(matrix); i++ {
		Printf(matrix[i])
	}

	matrix2 := [][]string {
		{"+", "+", "+"},
		{"+", "+", "+"},
		{"+", "+", "+"},
	}
	matrix2[0][1] = "-"
	matrix2[1][0] = "-"
	matrix2[1][2] = "-"
	matrix2[2][1] = "-"
	for i := 0; i < len(matrix2); i++ {
		Printf(matrix2[i])
	}

}