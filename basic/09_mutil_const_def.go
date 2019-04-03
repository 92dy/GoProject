package main

import "fmt"

func main(){
	const a int = 8
	const str string = "abc"
	fmt.Printf("a = %d, str = %s\n", a , str)
	
	//
	const (
		a1 = 8
		str1 = "abc"
	)
	fmt.Printf("a1 type is %T\nstr1 type is %T\n",a1, str1)
}

