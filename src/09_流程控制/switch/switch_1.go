package main

import "fmt"

func main() {
	var a int
	fmt.Printf("请输入一个数字：》》》 ")
	fmt.Scan(&a)

	switch {
	case a > 90:
		fmt.Println("优秀")
	case a > 70 && a < 80:
		fmt.Println("合格")
	case a < 60:
		fmt.Println("垃圾")
	default:
		fmt.Println("输入错误")
	}
}
