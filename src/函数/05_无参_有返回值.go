package main

import "fmt"

func myfunc01() int { // 无参，只有一个返回值
	return 666
}

//
func myfunc02() (result int) { //给返回值一个变量名字 ，go 推荐写法
	result = 666
	return
}

func main() {
	var a int // 无参有返回值函数调用
	a = myfunc01()
	fmt.Println("a =", a)

	b := myfunc01() //自动推导类型
	fmt.Println("b = ", b)

}
