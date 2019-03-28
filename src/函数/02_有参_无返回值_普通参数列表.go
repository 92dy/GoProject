package main

import "fmt"

func MyFunc01(a int) {
	fmt.Println("a = ", a)
}

func MyFunc02(A int, B int) {
	fmt.Printf("A=%d, B=%d\n", A, B)
}

func MyFunc03(a1, b1 int, c1 string, d1 byte) {
	fmt.Printf("a1 = %d, b1 =%d, c1 = $s, d1 = %sn", a1, b1, c1, d1)
}
func main() {
	MyFunc01(666)
	MyFunc02(777, 888)
	MyFunc03(111, 222, "hello", 'H')
}
