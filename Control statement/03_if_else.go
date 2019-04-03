package main 

import "fmt"

func main() {
	var a int 
	fmt.Printf("plase input a num >>>： ")
	fmt.Scan(&a)
	
	if a > 10 {
		fmt.Printf(" too big\n")
	} else if a < 10 {
		fmt.Printf("too smart\n")
	} else {
		fmt.Printf("right")
	}
}
