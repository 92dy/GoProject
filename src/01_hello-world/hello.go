//1、go 语言以包围管理单位
//2、每个文件必须先声明包
//3、程序必须有一个 main 包（重要）
package main

import "fmt"

//入口函数
func main() { //左括号必须与函数名同行
	//打印
	// Println() 会自动换行
	// 调用函数，大部分需要导入包
	/*
		这也是注释
	*/

	fmt.Println("hello world!") //go语言语句结尾没有分号

}
