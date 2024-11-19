package main

import "fmt"

/**
四种变量的声明方式
*/

// 声明全局变量 方法1，方法2，方法3是可以的
var gA int = 100
var gB = 200

// 方法4是不支持全局变量的
// z := 300
func main() {
	//方法1：声明一个变量 默认值是0
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)

	//方法2：声明一个变量，初始化一个值
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)

	//方法3：在初始化的时候，可以省去数据类型，通过值自动匹配当前的变量的数据类型
	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)

	//方法4：（常用的方法）省去var关键字，直接自动匹配
	e := 100
	fmt.Println("e = ", e)
	fmt.Printf("type of e = %T\n", e)

	f := "abcd"
	fmt.Println("f = ", f)
	fmt.Printf("type of f = %T\n", f)

	g := 3.14
	fmt.Println("g = ", g)
	fmt.Printf("type of g = %T\n", g)

	fmt.Println("gA = ", gA, "gB = ", gB)

	// 声明多个变量
	var xx, yy int = 100, 200
	fmt.Println("xx = ", xx, "yy = ", yy)
	var kk, ll = 100, "hello"
	fmt.Println("kk = ", kk, "ll = ", ll)

	// 多行的多变量声明
	var (
		aa = 100
		bb = true
	)
	fmt.Println("aa = ", aa, "bb = ", bb)
}
