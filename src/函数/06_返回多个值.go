package main

import "fmt"

func myfunc_1() (int, int, int) { //多个返回值
	return 1, 2, 3
}

// go语言推荐写法
func myfunc_2() (A int, B int, C int) {
	A, B, C = 111, 222, 333
	return
}

func main() {
	a, b, c := myfunc_1()
	fmt.Printf("a = %d, b= %d, c = %d\n", a, b, c)

	//
	A, B, C := myfunc_2()
	fmt.Printf("A =%d, B=%d, C=%d\n", A, B, C)
}
