package main

import (
	"fmt"
	"reflect"
)

func reflectNum(arg interface{}) {
	fmt.Println("type : ", reflect.TypeOf(arg))
	fmt.Println("value : ", reflect.ValueOf(arg))
}

func main() {
	var a = 1
	var b = 3.14
	var c = "hello"

	reflectNum(a)
	reflectNum(b)
	reflectNum(c)
}
