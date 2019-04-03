package main

import "fmt"

func main() {
	var a bool
	a = true
	fmt.Println("a = ", a)
	
	var b bool = false
	fmt.Println("b = ", b)
	
	c := false
	fmt.Printf("c type is %T\n", c)
	
	var d bool //没有初始值，默认为false
	fmt.Println("d = ", d)
}
