package main

import "fmt"

func main() {
	
	const num int = 10
	
	for {
		var number int
		fmt.Println("please inter a number>>>： ")
		fmt.Scan(&number)
		
		if number > num {
			fmt.Println(" Too Big !!!")
		} else if number < num {
			fmt.Println(" too small !!!")
		} else {
			fmt.Println("good !!!")
			break
		}
		
		
	}
	
}
