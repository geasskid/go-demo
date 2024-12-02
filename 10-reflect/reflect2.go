package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func (this *User) Call() {
	fmt.Println("user call")
	fmt.Println(this.Name)
	fmt.Println(this.Age)
}

func main() {
	user := User{"zhangsan", 20}
	user.Call()

	DoFiledAndMethod(user)
}

func DoFiledAndMethod(input interface{}) {
	//获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is", inputType.Name())

	//获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is", inputValue)

	//通过type 获取里面的字段
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()

		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	//通过type 获取里面的方法，调用
	for i := 0; i < inputType.NumMethod(); i++ {
		method := inputType.Method(i)
		fmt.Printf("%s: %v\n", method.Name, method.Type)
	}
}
