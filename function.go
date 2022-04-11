package main

import "fmt"

func prompt (name string, age int) {
    
	fmt.Println("your are = ",name)
	fmt.Println("your age is = ",age)
	
	
}

func main (){
	var name string
	var age int
	fmt.Scanln(&name)
	fmt.Scanln(&age)
	prompt(name, age)

}








//****************************************************

/*
func student (name string, age int) (string, int) {
    name = name + "Bhae"
	return name, age

}

func main (){
    
	
	name, age := student("ALI", 5)
	fmt.Println(name)
	fmt.Println(age)
    
} 
*/





//****************************************************
/*
func student (name string) string {

	return name
}

func main () {
	fmt.Println(student("jamal"))
}*/