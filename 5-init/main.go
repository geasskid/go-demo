package main

import (
	_ "modtest/lib1"      //无法使用当前包的方法，但是会执行当前的包内部的init()方法
	mylib2 "modtest/lib2" //给包名其个别名，通过别名来直接调用
)

func main() {
	//lib1.Lib1Test()
	mylib2.Lib2Test()
}
