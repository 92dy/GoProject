package main

import "fmt"

func Sum() int {
	// i := 1
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	// fmt.Println("sum = ", sum)
	return sum
}

func main() {
	SUM := Sum()
	fmt.Println("sum = ", SUM)
}
