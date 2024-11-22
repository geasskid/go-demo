package main

import "fmt"

// AnimalService 本质是一个指针
type AnimalService interface {
	Slesp()
	GetColor() string
	GetType() string
}

type Dog struct {
	color string
}

func (d Dog) Slesp() {
	println("wang")
}

func (d Dog) GetColor() string {
	return d.color
}

func (d Dog) GetType() string {
	return "dog"
}

func showAnimal(a AnimalService) {
	a.Slesp()
	fmt.Println(a.GetColor())
	fmt.Println(a.GetType())
}

func main() {
	var animal AnimalService
	animal = &Dog{"black"}

	animal.Slesp()
	fmt.Println(animal.GetColor())
	fmt.Println(animal.GetType())

	dog := &Dog{"Yellow"}
	showAnimal(dog)
}
