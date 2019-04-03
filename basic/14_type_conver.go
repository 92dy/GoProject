package main

import "fmt"

func main() {
	var flag bool
	flag = true

	fmt.Printf("flag = %t\n", flag)

	//bool类型不允许转换成Int类型, string类型,反指，int, string类型也不能转换为bool类型
	// fmt.Printf("flag = %d", int(flag))
	// fmt.Printf("flag = %s", string(flag))

	// 0就是假，非0就是真
	// flag = bool(1)

	var ch byte
	ch = 'a'

	var t int
	t = int(ch) // 此为类型转换
	fmt.Printf("%v\n", t)

}
