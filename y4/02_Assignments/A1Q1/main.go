package main

import (
	"fmt"
	"strconv"
)

func main() {
	var numerator1, denominator1, numerator2, denominator2 int
	var fraction1, fraction2, product, sum float32
	// ask user to input numerator and denominator for first fraction
	fmt.Println("Enter numerators and non-zero denominators")
	fmt.Println("  * for the 1st fraction:")
	fmt.Scanln(&numerator1)
	fmt.Scanln(&denominator1)
	// ask user to input numerator and denominator for second fraction
	fmt.Println("  * for the 2nd fraction:")
	fmt.Scanln(&numerator2)
	fmt.Scanln(&denominator2)

	// calculate the first fraction and save as float
	fraction1 = float32(numerator1) / float32(denominator1)
	// concatenate the first fraction and save as string
	fractionStr1 := strconv.Itoa(numerator1) + "/" + strconv.Itoa(denominator1)

	// calculate the second fraction and save as float
	fraction2 = float32(numerator2) / float32(denominator2)
	fmt.Println("fraction2 is", fraction2)
	// concatenate the second fraction and save as string
	fractionStr2 := strconv.Itoa(numerator2) + "/" + strconv.Itoa(denominator2)

	// calculate the product of 2 fractions and save as double
	product = fraction1 * fraction2
	// calculate the sum of 2 fractions and save as double
	sum = fraction1 + fraction2

	// concatenate the product
	productStr := strconv.Itoa(numerator1*numerator2) + "/" + strconv.Itoa(denominator1*denominator2)
	// concatenate the sum
	sumStr := strconv.Itoa((numerator1*denominator2)+(numerator2*denominator1)) + "/" + strconv.Itoa(denominator1*denominator2)

	fmt.Println("The product of", fractionStr1, "and", fractionStr2, "is", productStr, "or", product)
	fmt.Println("The sum of", fractionStr1, "and", fractionStr2, "is", sumStr, "or", sum)
	fmt.Print("\nStrange calculations!!!")
}
