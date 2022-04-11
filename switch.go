package main

import  "fmt"

func main(){
   var num int
   fmt.Println("Enter your Marks")
   fmt.Scanln(&num)

   switch  {
   case (num >= 90 && num <=100):
	fmt.Println("A Grade")
   case (num >= 80 && num <= 90):
	fmt.Println("B Grade")
   case (num >= 70 && num <= 80):
	fmt.Println("c Grade")
   case (num >= 60 && num <= 70):
	fmt.Println("D Grade")
   
   }
}