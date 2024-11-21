package main

import "fmt"

func main() {
	s := []int{1, 2, 3}

	s1 := s[0:2]

	fmt.Println(s1)

	s[0] = 100
	fmt.Println(s)
	fmt.Println(s1)

	fmt.Println("====")

	s2 := make([]int, 3)
	copy(s2, s1)
	fmt.Println(s2)

	s2[0] = 200
	fmt.Println(s)
	fmt.Println(s2)
}
