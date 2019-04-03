package main

import "fmt"

func main() {
	i := 12
	f := 3.14
	ch := 'a'
	str := "abc"

	fmt.Printf("%T\t,%T\t, %T\t,%T\n", i, f, ch, str)

	fmt.Printf("i = %d, f = %f, ch = %c, str = %s\n", i, f, ch, str)
	//%v自动匹配格式输出
	fmt.Printf("i = %v, f = %v, ch = %v, str = %v\n", i, f, ch, str) 
}
