package main

import "fmt"
			// y is identifier with varaible 
var y = 10   // declare var with y .... used in whole program which Short declaration operator can be use in a internal function
var z int 

func main() {
	// Short declaration operator
	x := 23      // Auto pick the type
	fmt.Println(x)

	x = 44    // no need of again declare the type already declare.
	fmt.Println(x)

	
	foo()

}

func foo() {
	fmt.Println(y)
	fmt.Println(z)
}