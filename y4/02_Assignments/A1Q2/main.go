package main

import (
	"fmt"
	"strings"
)

func main() {
	// declare city, upperCity, secondUpperCity, secondToLastUpperCity, middleUpperCity as string
	var city, upperCity, secondUpperCity, secondToLastUpperCity, middleUpperCity string
	// declare length, middleIndex as integer
	var length, middleIndex int

	// print welcome message
	fmt.Println("/             Welcome To              /")
	fmt.Println("/  Concordia's City Name Manipulator  /")
	fmt.Println("   ---------------------------------   ")
	// ask user to input city
	fmt.Println("Please enter the name of your favorite city in lower case:")
	// assign input string to city
	fmt.Scanln(&city)

	// get the length from city string
	length = len(city)
	// calculate the middle index of city string
	middleIndex = length / 2
	// format all the letters to upper case
	upperCity = strings.ToUpper(city)

	// index 1 which is second letter to be upper case, concatenate the rest letters
	secondUpperCity = city[0:1] + strings.ToUpper(city[1:2]) + city[2:length]
	// index length - 1 which is second to last letter to be upper case, concatenate the rest letters
	secondToLastUpperCity = city[0:length-2] + strings.ToUpper(city[length-2:length-1]) + city[length-1:length]
	// middleIndex to be upper case, concatenate the rest letters
	middleUpperCity = city[0:middleIndex] + strings.ToUpper(city[middleIndex:middleIndex+1]) + city[middleIndex+1:length]

	// print output
	fmt.Println("You entered", city+"which has", length, " characters.")
	fmt.Println("Here is the city name")
	fmt.Println(" * With all letters in upper case:", upperCity)
	fmt.Println(" * With the second letter in upper case:", secondUpperCity)
	fmt.Println(" * With the second to last letter in upper case:", secondToLastUpperCity)
	fmt.Println(" * With the middle letter in upper case:", middleUpperCity)
	fmt.Println("Hope you are confortable manipulating String variables now...")
	fmt.Println("Program exited...")
}
