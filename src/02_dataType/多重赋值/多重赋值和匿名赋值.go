package main

import "fmt"

func main() {
	// a := 10
	// b := 20
	// c := 30

	// a, b := 10, 30
	// var tmp int
	// tmp = a
	// a = b
	// b = tmp
	// fmt.Printf("a = %d, b = %d\n", a, b)

	// i := 10
	// j := 20
	// i, j = j, i
	// fmt.Printf("i = %d, j = %d\n", i, j)

	//_匿名处理，丢弃数据不处理
	i := 10
	j := 20
	var tmp int
	tmp, _ = i, j
	fmt.Println("tmp = ", tmp)

}
