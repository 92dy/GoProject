package main

import "fmt"

func MinAndMax(num1 int, num2 int) (min int, max int) {
	if num1 > num2 {
		min = num2
		max = num1
	} else {
		min = num1
		max = num2
	}
	
	return
}

func main () {
	min, max := MinAndMax(22, 44)
	fmt.Printf("min = %d\nmax = %d\n", min, max)
}
