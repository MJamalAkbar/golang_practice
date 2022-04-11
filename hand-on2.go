package main

import "fmt"

const (
	a = 2016 + iota
	b = iota + 2016
	c = iota + 2016
	d = iota + 2016
)

func main(){
	fmt.Println(a, b, c, d)
}
