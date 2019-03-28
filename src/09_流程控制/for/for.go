package main

import "fmt"

func main() {
	// for 初始化条件; 判断条件; 条件变化 {}
	// 1+2+3+...+10 累加
	// sum := 0
	// for i := 1; i <= 100; i++ {
	// 	sum = sum + i
	// }
	// fmt.Println("sum = ", sum)

	// for 打印每个字符串
	str := "abc"
	for i := 0; i < len(str); i++ {
		fmt.Printf("str[%d]=%c\n", i, str[i])
	}

	// 迭代打印每个元素，默认返回2个值
	for i, data := range str {
		fmt.Printf("str[%d]=%c\n", i, data)
	}

	//
	for i := range str {
		fmt.Printf("str[%d]=%c\n", i, str[i])
	}
}
