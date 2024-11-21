package main

import "fmt"

func main() {
	//声明myMap1是一种map类型，key是string，value是string类型
	myMap1 := make(map[string]string)

	myMap1["one"] = "1"
	myMap1["two"] = "2"
	myMap1["three"] = "3"

	fmt.Println(myMap1)

	myMap2 := map[string]string{
		"one": "1",
		"two": "2",
	}
	fmt.Println(myMap2)
}
