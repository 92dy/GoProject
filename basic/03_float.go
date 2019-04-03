package main

import "fmt"

func main() {
	var f1 float64 
	f1 = 2.13
	fmt.Println("f1 = ", f1)
	
	//自动推导
	f2 := 3.14
	fmt.Printf("f2 type is %T, f2 = %f", f2,f2)
}
