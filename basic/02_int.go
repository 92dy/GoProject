package main

import "fmt"
func main() {
	//变量，程序运行期间，可以改变的量

	//声明： var 变量名 类型 值,不给值默认为0
	var a int
	fmt.Println("a = ", a)

	//可以同时声明多个变量
	//var b, c int

	//2、变量的初始化，声明变量是同时赋值
	var b int = 10
	fmt.Println("b = ", b)
	//3、自动推导类型，必须初始化，通过初始值确定类型
	// 自动推导，同一个变量名只能使用一次，用于初始化那次
	age := 26 //先声明age的类型，再给age赋值
	fmt.Printf("age type is %T", age)
}
