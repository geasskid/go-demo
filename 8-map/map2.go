package main

import "fmt"

func printMap(cityMap map[string]string) {
	for key, value := range cityMap {
		fmt.Println(key, "=", value)
	}
}

func ChangeValue(cityMap map[string]string) {
	cityMap["four"] = "大连"
}

func main() {
	cityMap := make(map[string]string)

	cityMap["one"] = "北京"
	cityMap["two"] = "上海"
	cityMap["three"] = "广州"

	for key, value := range cityMap {
		fmt.Println(key, "=", value)
	}

	//删除
	delete(cityMap, "two")
	//修改
	cityMap["three"] = "深圳"

	fmt.Println("======")

	for key, value := range cityMap {
		fmt.Println(key, "=", value)
	}

	fmt.Println("======")

	printMap(cityMap)

	fmt.Println("======")

	ChangeValue(cityMap)
	printMap(cityMap)

}
