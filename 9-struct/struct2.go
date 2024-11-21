package main

import "fmt"

// Hero 如果类名首字母大写，表示其他包也能访问
type Hero struct {
	// 如果说累的属性首字母大写，表示该属性是公开属性，其他包也能访问，否则是私有属性，其他包不能访问
	name  string
	age   int
	level int
}

func (hero *Hero) GetName() {
	fmt.Println("Name = ", hero.name)
}

func (hero *Hero) SetName(name string) {
	hero.name = name
}

func main() {
	hero := Hero{
		name:  "hero",
		age:   18,
		level: 1,
	}
	hero.GetName()
	hero.SetName("hero2")
	hero.GetName()
}
