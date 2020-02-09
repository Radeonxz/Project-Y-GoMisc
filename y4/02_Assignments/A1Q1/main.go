package main

import (
	"fmt"
	"strconv"
)

func main() {
	var numerator1, denominator1, numerator2, denominator2 int
	var fraction1, fraction2, product, sum float32
	var productStr, sumStr string
	// ask user to input numerator and denominator for first fraction
	fmt.Print("Enter numerators and non-zero denominators")
	fmt.Print("  * for the 1st fraction:")
	fmt.Scanln(&numerator1)
	fmt.Scanln(&denominator1)
	// ask user to input numerator and denominator for second fraction
	fmt.Print("  * for the 2nd fraction:")
	fmt.Scanln(&numerator2)
	fmt.Scanln(&denominator2)

	// calculate the first fraction and save as float
	fraction1 = float32(numerator1 / denominator1)
	// concatenate the first fraction and save as string
	fractionStr1 := strconv.Itoa(numerator1) + "/" + strconv.Itoa(denominator1)

	// calculate the second fraction and save as float
	fraction2 = float32(numerator2 / denominator2)
	// concatenate the second fraction and save as string
	fractionStr2 := strconv.Itoa(numerator2) + "/" + strconv.Itoa(denominator2)
	fmt.Print("\nresult is ", fraction1, fraction2, product, sum, fractionStr1, fractionStr2, productStr, sumStr)
	fmt.Print("\nStrange calculations!!!")
}
