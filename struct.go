package main

import "fmt"

func main() {
	type location_st struct {
		X int
		Y int
		name string
	}

	var st location_st
	st = location_st{X: 1, Y: 2, name: "Beijing"}
	// st := location_st{X: 1, Y: 2, name: "Beijing"}
	fmt.Println("st:", st)
	fmt.Println("st.X", st.X)
	
	var st1 location_st
	st1 = location_st{3,4,"Shanghai"}
	// st1 := location_st{3,4,"Shanghai"}
	fmt.Println("st1:", st1)

	var p *location_st
	p = &st1
	// p := &st1
	fmt.Println("&st1:", &st1)
	fmt.Println("*p:", *p)
	fmt.Println("p:", p)
	fmt.Println("p.name:", p.name)
	p.name = "Shenzhen"
	fmt.Println("p.name:", p.name)
	fmt.Println("st1.name:", st1.name)
	fmt.Println("(*p).name:", (*p).name)
	fmt.Println("&p:", &p)
	
}
