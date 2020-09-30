package main

import "fmt"

type Student struct {
	name string
	score int
	age int
}

var stMap map[string]Student

func main() {
	stMap = make(map[string]Student)
	stMap["xiaoming"] = Student{name: "xiaoming", score: 100, age: 24}
	stMap["mike"] = Student{name: "mike", score: 80, age: 25}
	fmt.Println(stMap)
	fmt.Println(stMap["xiaoming"])
	fmt.Println(stMap["mike"])

	var stMap2 = map[string]Student {
		"amy": Student{name: "amy"},
		"sara": Student{name: "sara", score: 20},
	}
	fmt.Println(stMap2)

	stMap3 := map[string]Student {
		"lary": Student{name: "lary"},
		"jose": Student{name: "jose", score: 30},
	}
	fmt.Println(stMap3)

	// 映射和数组必须使用make去初始化声名的变量
	var stMap4 map[string]Student
	if stMap4 == nil {
		fmt.Println("stMap4 is nil")
	}

	stMap4 = make(map[string]Student)
	stMap4["zhaogang"] = Student{name: "zhaogang"}
	stMap4["ligang"] = Student{name: "ligang"}
	// 若键不存在，则返回该类型的零值
	fmt.Println(stMap4["ligang2"])

	var elem Student
	var ok bool

	elem, ok = stMap4["hulang"]
	fmt.Println(elem)
	fmt.Println(ok)

	elem, ok = stMap4["ligang"]
	fmt.Println(elem)
	fmt.Println(ok)
	fmt.Println(stMap4["ligang"].name)

	if _, exist := stMap4["hulang"]; !exist {
		fmt.Println("hulang is not exist")
	}

	delete(stMap4, "ligang")
	fmt.Println(stMap4)
		
	// stMap4["harry"] = Student{name: "harry", score: 52, age: 24}
	// fmt.Println(stMap4)

	var arr []int
	if arr == nil {
		fmt.Println("arr is nil")
	}
	arr = make([]int, 0, 1)
	arr = append(arr, 10, 9, 8)
	fmt.Println(arr)
}
