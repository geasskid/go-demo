package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个带缓冲的channel
	c := make(chan int, 3)

	fmt.Println("len(c) = ", len(c), ", cap(c) = ", cap(c))

	go func() {
		defer fmt.Println("goroutine end")

		for i := 0; i < 5; i++ {
			c <- i
			fmt.Println("子线程正在运行，发送的元素 = ", i, " len(c) = ", len(c), ", cap(c) =", cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 5; i++ {
		num := <-c
		fmt.Println("接收的元素 = ", num)
	}

	fmt.Println("main goroutine end")
}
