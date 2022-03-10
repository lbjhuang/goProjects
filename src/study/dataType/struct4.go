package main

import "fmt"

//结构体继承
type OnePerson struct {
	id   int
	name string
	age  int
}

//student 以属性的方式继承OnePerson
type Student struct {
	OnePerson
	id        int
	className string
	score     float32
}

func main() {
	var s = Student{
		OnePerson{1, "james", 22}, 2, "class1", 98.565}

	fmt.Printf("one student named %s, his score is %.1f, he is from %s and he is just %d", s.name, s.score, s.className, s.age)
}
