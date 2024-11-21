package main

import "fmt"

func printSlice(myArray []int) {
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}
	myArray[0] = 100
}
func main() {
	myArray := []int{1, 2, 3, 4}

	fmt.Printf("myArray types = %T\n", myArray)

	printSlice(myArray)

	fmt.Println("====")
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}
}
