package main

import "fmt"

func main() {
	//声明slice1是一个切片，并且初始化，默认值是1，2，3，长度len是3
	slice1 := []int{1, 2, 3}

	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)

	var slice2 []int
	slice2 = make([]int, 3)
	fmt.Printf("len = %d, slice = %v\n", len(slice2), slice2)

	slice3 := make([]int, 3)
	fmt.Printf("len = %d, slice = %v\n", len(slice3), slice3)

	var slice4 []int
	if slice4 == nil {
		fmt.Println("slice4 is nil")
	} else {
		fmt.Println("slice4 is not nil")
	}
}
