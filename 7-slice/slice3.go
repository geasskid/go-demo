package main

import "fmt"

func main() {
	numbers := make([]int, 3, 5)

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 1)
	numbers = append(numbers, 2)

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 5)

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	fmt.Println("========")
	numbers1 := make([]int, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers1), cap(numbers1), numbers1)
	numbers1 = append(numbers1, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers1), cap(numbers1), numbers1)

}
