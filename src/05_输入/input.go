package main

import "fmt"

func main() {
	var a int
	fmt.Printf("请输入变量a>>: ")

	// 堵塞等待用户的输入
	// fmt.Scanf("%d", &a)
	fmt.Scan(&a)
	fmt.Printf("%d\n", a)

}
