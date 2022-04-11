// even and odd program
// 1. declare first then from user

package main

import "fmt"

/*
func main(){
	var a int 

	fmt.Scanln(&a)
	if a%2==0 {
		fmt.Print("Even")
	} else {
		fmt.Println("Odd")
	}
}
*/


func main(){
	var a int
	fmt.Scanln(&a)

	switch  {
		case a==1:
		fmt.Println("Monday")
		case a==2:
		fmt.Println("Tue")
		case a==3:
		fmt.Println("Wed")
		case a==4:
		fmt.Println("Thur")
		case a==5:
		fmt.Println("Fri")
		case a==6:
		fmt.Println("sat")
		default:
		fmt.Println("Kch nhi")
	}

}

/*

func main(){ 
	switch day:=4; day{
	     case 1:
	     fmt.Println("Monday")
	     case 2:
	     fmt.Println("Tuesday")
	     case 3:
	     fmt.Println("Wednesday")
	     case 4:
	     fmt.Println("Thursday")
	     case 5:
	     fmt.Println("Friday")
	     case 6:
	     fmt.Println("Saturday")
	     case 7:
	     fmt.Println("Sunday")
	     default: 
	     fmt.Println("Invalid")
	}
}
*/