
package main

import "fmt"

func main() {
	fmt.Println("11111111")
		
	//End是用户定义的标签,无条件跳转，从此处退出
	goto End
		
	fmt.Println("22222222")
End:
	fmt.Println("33333333")
}
