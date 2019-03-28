package main

import "fmt"

func main() {
	//不同类型变量的定义
	// var a int
	// var b float64
	var (
		a int     = 1
		b float64 = 10.4
	)

	a, b = 10, 3.14
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	// const i int = 10
	// const j float64 = 3.14
	const ( // 常量会自动推导类型
		i = 10
		j = 3.14
		c = "dengyou"
	)
	fmt.Println("i = ", i)
	fmt.Println("j = ", j)
	fmt.Printf("c type is %T\n", c)
}
