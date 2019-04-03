package main

import "fmt"

func main() {
	//不同类型的定义
	// var a int
	// var f float64
	var (
		a int = 18
		f float64 = 3.14
	)
	fmt.Printf("a = %d, f=%f\n", a,f)
	//自动推导类型, 不需要添加：
	a, f = 18, 3.4
	fmt.Printf("a type is %T,\nf type is %T\n", a, f)
}
