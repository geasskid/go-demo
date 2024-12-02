package main

import "fmt"

func main() {

	var a string
	//pair<type:string, value:hello>
	a = "hello"

	var allType interface{}
	allType = a

	str, _ := allType.(string)
	fmt.Println(str)

}
