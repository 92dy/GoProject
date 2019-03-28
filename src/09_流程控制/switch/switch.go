package main

import "fmt"

func main() {

	var num int
	fmt.Printf("请输入你要进入的楼层号：>>>")

	fmt.Scan(&num)

	switch num {
	case 1:
		fmt.Printf("按下的是%d楼\n", num)
		// fallthrough //不跳出switch语句，后面的无条件执行
	case 2:
		fmt.Printf("按下的是%d楼\n", num)
	case 3:
		fmt.Printf("按下的是%d楼\n", num)
	case 4:
		fmt.Printf("按下的是%d楼\n", num)
	case 5:
		fmt.Printf("按下的是%d楼\n", num)
	default:
		fmt.Println(" 你输入的楼层不存在")
	}
}
