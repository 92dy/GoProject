package main

import "fmt"

func myfunc(tmp ...int) {
	for _, data := range tmp {
		fmt.Println("data = ", data)
	}
}

func myfunc2(tmp2 ...int) {
	for _, data2 := range tmp2 {
		fmt.Println("data2 = ", data2)
	}
}

func Test01(args ...int) {
	//全部参数传递给myfunc
	myfunc(args...)

	//传递部分参数
	fmt.Println("================")

	myfunc2(args[3:]...)  // 从args[3]到最后的元素传递过去,包含3自身
	myfunc2(args[0:3]...) //从args[0] 到args[2],不包含3本身
}

func main() {
	Test01(1, 2, 3, 4, 5)
}
