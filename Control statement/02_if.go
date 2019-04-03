package main

import "fmt"

func main() {
	var str string
	fmt.Printf("Pleas input String>>>: ")
	fmt.Scan(&str)
	
	if str == "yes" {
		fmt.Printf("hello world\n")
	}
	
	if a := 10; a == 10 {
		fmt.Println("a = ", a)
	}
}
