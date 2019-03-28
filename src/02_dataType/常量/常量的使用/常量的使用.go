package main

import "fmt"

func main() {
	// 常量： 程序运行期间，不可以改变的量
	const a int = 10
	// a = 20 // err, 常量不允许修改
	fmt.Println("a = ", a)

	const b = 10 // 自动推导类型不需要使用:=
	fmt.Printf("b type is %T\n", b)
}
