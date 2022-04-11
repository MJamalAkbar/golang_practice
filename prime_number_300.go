package main

import "fmt"

func fact1(num1 int)int {
	if num1%2 == 0 {
		fmt.Println("Prime number")
	} else {
		fmt.Println("Non prime number")
	}
		

}

func main(){
	var num int
	fmt.Println("Enter a Number", num)
	fmt.Scanln(&num)
}