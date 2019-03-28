package main

import "fmt"

func MyFunc_1(args ...int) { //不定参数
	fmt.Println("len(args) = ", len(args))

	// for i := 0; i < len(args); i++ {
	// fmt.Printf("args[%d] = %d\n", i, args[i])
	// }
	for i, data := range args {
		fmt.Printf("args[%d]=%d\n", i, data)
	}
}

func main() {
	MyFunc_1(123, 456, 789, 999)
}
