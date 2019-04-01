package main

import "fmt"

func myfunc01() int  {
	return 666
}

func myfunc02() (result int) {
	result = 777
	return
}

func main() {
	var a int 
	a = myfunc01()
	
	fmt.Println("a = " , a)
	
	b := myfunc02()
	fmt.Printf("b = %d", b)
}
