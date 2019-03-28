package main

import "fmt"

func main() {
	var a int
	fmt.Printf("input a num>>>: ")
	fmt.Scan(&a)

	if a == 10 {
		fmt.Printf("%d == 10\n", a)
	} else if a > 10 {
		fmt.Printf("%d > 10\n", a)
	} else {
		fmt.Printf("%d < 10\n", a)
	}

}
