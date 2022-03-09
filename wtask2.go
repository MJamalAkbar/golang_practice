package main

import "fmt"

func main () {
        fmt.Println("Enter your age = ")
		var age int 
		fmt.Scanln(&age)
		

		fmt.Println("Enter your height = ")
		var height float32 
		fmt.Scanln(&height)
		

		fmt.Println("Enter your weight = ")
		var weight int 
		fmt.Scanln(&weight)
		fmt.Println("Your Age is= ", age)
		fmt.Println("Your height is= ", height)
		fmt.Println("Your weight is= ", weight)
}