package main

import "fmt"

func main() {

	const a int = 10
	// a= 20 //err 常量不允许修改
	fmt.Println("a = ", a)
	
	//自动推导类型 ,不需要使用 :=
	const b = "abc"
	fmt.Printf("b type is : %T", b)
	
}
