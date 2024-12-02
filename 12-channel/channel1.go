package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine end")

		fmt.Println("goroutine start")
		c <- 666
	}()

	num := <-c
	fmt.Println("main goroutine end, num = ", num)
}
