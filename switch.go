package main

import (
	"fmt"
	"time"
)

func main() {
	// go的switch中的每个case都会自动终止，类似break效果
	// case的执行顺序由上至下，遇到匹配的case则停止执行
	switch v := "hello"; v {
	case "hello":
		fmt.Println(v)
	case "hi":
		fmt.Println(v)
	default:
		fmt.Println(v)
	}

	// swith相当于多个if-else
	t := time.Now()
	fmt.Println("Good Morning ! \n", t)
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning ! \n", t.Hour())
	case t.Hour() < 18:
		fmt.Println("Good Afternoon ! \n", t.Hour())
	default:
		fmt.Println("Good Eve ! \n", t.Hour())
		
	}
}
