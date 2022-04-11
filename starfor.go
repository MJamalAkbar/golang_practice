package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter a number")
	fmt.Scanln(&n)
	for i := 0; i <= n; i++ {
		for j := 0; j <= n-i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
		
	}

}
