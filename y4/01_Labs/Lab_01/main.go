package main

import (
	"fmt"
)

func main() {
	var a, b, c, d, num, den int
	var CONSTANT int = 4
	fmt.Print("Input 4 positive integers: ")
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)
	fmt.Scanln(&d)
	num = (a%d + c%b) * CONSTANT
	den = b * d * CONSTANT
	fmt.Print("\nModified numerator is ", num)
	fmt.Print("\nModified denominator is ", den)
	fmt.Print("\nStrange calculations!!!")
}
