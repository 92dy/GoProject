package main

import "fmt"

func main() {
	
	var num int
	fmt.Printf("请输入你想进入的楼层>>>: ")
	fmt.Scan(&num)
	
	switch num {
	case 1:
		fmt.Printf("电梯即将到达%d楼", num)
	case 2:
		fmt.Printf("电梯即将到达%d楼", num)
	case 3:
		fmt.Printf("电梯即将到达%d楼", num)
	case 4:
		fmt.Printf("电梯即将到达%d楼", num)
	case 5:
		fmt.Printf("电梯即将到达%d楼", num)
	default:
		fmt.Println(" 您输入的数字有误")
	}
}
