package main 

import "fmt"

func Myfunc01(a int) {  //形参 name type
	fmt.Println("a = ", a)
}

func Myfunc02(A int, B int) {
	fmt.Printf("A = %d, B = %d\n", A, B)
}

func Myfunc03(a1, b1 int, ch byte, str string) {
	fmt.Printf("a1 = %d, b1 = %d, ch = %c, str =%s", a1, b1, ch, str)
}

func main() {
	Myfunc01(111) // 实参， 传参
	Myfunc02(222, 333)
	Myfunc03(444, 555, 'c', "hello world") 
}

