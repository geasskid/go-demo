package main

import "fmt"

func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println(arg)

	//给interface{} 提供“类型断言”机制
	value, ok := arg.(string)
	if ok {
		fmt.Println("arg is string type, value = ", value)
	} else {
		fmt.Println("arg is not string type")
	}
}

type Book1 struct {
	auth string
}

func main() {
	book := Book1{"Go"}

	myFunc(book)
	myFunc(100)
	myFunc("hello")
}
