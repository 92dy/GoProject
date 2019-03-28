package main

import "fmt"

func main() {
	var ch byte    //字符单引号
	var str string //字符串双引号

	ch = 'a'
	fmt.Printf("a = %c\n", ch)

	str = "a" //字符串默认都是隐藏了一个结束符 '\0'
	fmt.Println("str = ", str)

	str1 := "hello go"
	fmt.Println("str1 = ", str1)

	fmt.Printf("str1[0] = %c, str1[1] = %c, str[5]=%c\n", str1[0], str1[1], str1[5])

}
