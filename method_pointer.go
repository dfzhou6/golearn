package main

import "fmt"

type Student struct {
	name string
	score int
	age int
}

func (s *Student)calculateScore(rate int) {
	s.score = s.score * rate
}

func (s Student)printScore() {
	fmt.Printf("my name is %v, age is %v, score(rate) is %v\n", s.name, s.age, s.score)
}

func calculateScoreMethod(s *Student, rate int) {
	s.score = s.score * rate
}

func printScoreMethod(s Student) {
	fmt.Printf("my name is %v, age is %v, score(rate) is %v\n", s.name, s.age, s.score)
}

func main() {
	student := Student{ name: "zdf", age: 24, score: 100 }

	student.calculateScore(10)
	student.printScore()

	calculateScoreMethod(&student, 10)
	printScoreMethod(student)


	student2 := &Student{ name: "zdf2", age: 25, score: 120 }

	student2.calculateScore(100)
	student2.printScore()

	calculateScoreMethod(student2, 10)
	printScoreMethod(*student2)
}
