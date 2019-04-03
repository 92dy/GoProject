package main

import "fmt"

func main() {
 
	// for 初始化条件;判断条件; 条件变化 {}  1+2+3+...+100
	sum := 0
	for i := 0; i<= 100; i++ {
		sum = sum +i
	}
	fmt.Println("sum = ", sum)
	
	// for 打印字符串
	str := "hello"
	for i := 0; i < len(str); i++ {
		fmt.Printf("str[%d] = %c\n", i, str[i])
	}
	
	// range 打印,默认返回两个值， i, data
	str1 := "world"
	for i, data := range str1 {
		fmt.Printf("str[%d] = %c\n", i, data)
	}
	
	// range 默认返回 i , str[i]
	str2 := "dengyou"
	for i := range str2{
		fmt.Printf("str2[%d] = %c\n", i, str2[i])
	}
	
}
