package main

import "fmt"

func main() {
	// v1
	a := 11
	if a == 10 {
		fmt.Println("a == 10")
	} else {
		fmt.Println("a != 10")
	}

	// v2
	if a := 10; a == 10 { // 包含初始化
		fmt.Println("a == 10")
	} else {
		fmt.Println("a != 10")
	}

	// v3
	b := 10
	if b == 10 {
		fmt.Println("b == 10")
	} else if b > 10 {
		fmt.Println("b > 10")
	} else if b < 10 {
		fmt.Println("b < 10")
	} else {
		fmt.Println("这是不可能的！！")
	}
}
