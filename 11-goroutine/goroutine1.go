package main

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new task %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go newTask()
	for i := 0; i < 5; i++ {
		fmt.Printf("main task %d\n", i)
		time.Sleep(5 * time.Second)
	}
}
