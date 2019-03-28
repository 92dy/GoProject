package main

import "fmt"

func main() {
	var f1 float32
	f1 = 3.14
	fmt.Println("f1 = ", f1)

	//自动推导
	f2 := 3.24
	fmt.Printf("f1 type is %T\n", f2)
}
