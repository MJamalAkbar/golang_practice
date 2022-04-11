package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f uint = math.Sqrt(uint(x*x + y*y))
	var z float64 = float64(f)
	fmt.Println(x, y, z)
}
