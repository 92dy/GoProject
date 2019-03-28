package main

import "fmt"

func main() {

	fmt.Println("11111111111111")

	goto End //End是用户起的标签，无条件跳转

	fmt.Println("222222222222222")
End:
	fmt.Println("333333333333333")
}
