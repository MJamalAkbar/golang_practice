package main

import  "fmt"

func main(){
   var num int
   fmt.Println("Enter your Marks")
   fmt.Scanln(&num)
   
   if num >= 90 || num == 100{
      fmt.Println("A Grade")
   } else if num >= 80 && num < 90{
      fmt.Println("B  Grade")
   } else if num >= 70 && num < 80{
      fmt.Println("C Grade")
   } else if num >= 60 && num < 70{
      fmt.Println("D Grade")
   } else {
      fmt.Println("Fail")
   }
}