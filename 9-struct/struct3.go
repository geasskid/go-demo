package main

import "fmt"

type Human struct {
	name string
	sex  int
}

func (human *Human) Eat() {
	fmt.Println("Human Eat")
}

func (human *Human) Walk() {
	fmt.Println("Human walk")
}

type SuperMan struct {
	Human //superMan类继承了Human类的方法
	level int
}

func (superMan *SuperMan) Eat() {
	fmt.Println("SuperMan Eat")
}

func (superMan *SuperMan) fly() {
	fmt.Println("SuperMan fly")
}

func main() {
	h := Human{"小王", 0}

	h.Eat()
	h.Walk()

	s := SuperMan{Human{"小明", 1}, 100}
	s.Eat()
	s.Walk()
	s.fly()
}
