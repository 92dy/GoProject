package main

import "fmt"

func main() {
	var a bool
	a = true
	fmt.Println("a = ", a)
	// 自动推导类型
	var b = false
	fmt.Println("b = ", b)

	c := false
	fmt.Println("c =", c)

	var d bool //没有初始值默认为false
	fmt.Println("d = ", d)
}
