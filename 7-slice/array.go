package main

import "fmt"

func printArray(myArray [4]int) {
	for index, value := range myArray {
		fmt.Println("index = ", index, "value = ", value)
	}
}

func main() {
	//固定长度的数组
	var myArray1 [5]int

	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])
	}

	myArray2 := [10]int{1, 2, 3, 4}
	for index, value := range myArray2 {
		fmt.Println("index = ", index, "value = ", value)
	}

	myArray3 := [4]int{1, 2, 3, 4}

	fmt.Printf("myArray1 types = %T\n", myArray1)
	fmt.Printf("myArray2 types = %T\n", myArray2)
	fmt.Printf("myArray3 types = %T\n", myArray3)

	printArray(myArray3)
}
