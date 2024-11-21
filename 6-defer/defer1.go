package main

import "fmt"

func main() {
	//1. defer的执行顺序
	defer fmt.Println("main end1")
	defer fmt.Println("main end2")

	fmt.Println("main start")
}
