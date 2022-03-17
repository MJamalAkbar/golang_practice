// simple calculator where values are passing to paremeter and perform task.
package main


import "fmt"


func add(a, b int) int{        //addition function 
	sum := a+b
	return sum
	
}
func subt(a, b int) int{
	minus := a-b
	return minus

}

func multiple(a, b int) int{
	multi := a*b
	return multi

}
func divider(a, b float64) float64{
	div := a/b
	return div
}


func main() {                 
	x:= add(1, 2)						//call a funtion add and print it
	fmt.Println("Addition = ",x)

	sub := subt(4,5)					//call a funtion subt and print it
	fmt.Println("Substration = ",sub)

    mult := multiple(1,5)				//call a funtion multiple and print it	
	fmt.Println("Multiple = ",mult)
	
	divi := divider(2,4)				//call a funtion dividi and print it	
	fmt.Println("divide = ",divi)
}