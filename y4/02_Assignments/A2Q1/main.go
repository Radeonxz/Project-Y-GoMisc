package main

import (
	"fmt"
)

func main() {
	// print welcome message
	fmt.Println("-------****-------****-------****-------****-------****-------")
	fmt.Println("           Welcome to Montreal Taxi Fare Estimator!           ")
	fmt.Println("-------****-------****-------****-------****-------****-------")

	// declare carModel, serviceDay, serviceTime and distance as float64
	var carModel, serviceDay, serviceTime, distance float64
	// declare fare as float64
	var fare float64
	// assign 0.81 to baseFare
	baseFare := 0.81
	// assign 1.55 to limoFare
	limoFare := 1.55
	modelFare := 0.0
	rushHourRatio := 0.0

	// print car model
	fmt.Println("Please select your taxi car model:")
	fmt.Println("     1 - base car")
	fmt.Println("     2 - limo car")
	// ask user to select model
	fmt.Println("Please enter the digit corresponding to your case:(1 or 2)")
	// assign first input number to car model
	fmt.Scanln(&carModel)

	/**
	 * multiple conditions check to determine modelFare
	 * */
	// condition check based on if car model is base car or limo car
	if carModel == 1.0 {
		// assign baseFare to modelFare if carModel is 1
		modelFare = baseFare
	} else if carModel == 2.0 {
		// assign limoFare to modelFare if carModel is 2
		modelFare = limoFare
	}

	// ask user to select service day
	fmt.Println("Please select which day you will be using the taxi service:")
	fmt.Println("     1 - Weekday")
	fmt.Println("     2 - Weekend")
	// assign second input number to serviceDay
	fmt.Scanln(&serviceDay)

	/**
	 * multiple conditions check to determine rushHourRatio
	 * */
	// condition check based on if service day is weekday or weekend
	if serviceDay == 1 {
		// ask user to select service time if service day is weekday
		fmt.Println("Please select the time you will be using the taxi service:")
		fmt.Println("     1 - During 8am - 10am and 4pm - 5pm")
		fmt.Println("     2 - From midnight 12am to 6am")
		fmt.Println("     3 - None of the above")
		// assign third input number to serviceTime
		fmt.Scanln(&serviceTime)
		// condition check based on service time
		if serviceTime == 1.0 {
			// assign 1.5 to rushHourRatio if serviceTime is 1
			rushHourRatio = 1.5
		} else if serviceTime == 2.0 {
			// assign 2 to rushHourRatio if serviceTime is 2
			rushHourRatio = 2.0
		} else {
			// assign 1 to rushHourRatio if serviceTime is any number
			rushHourRatio = 1.0
		}
	} else if serviceDay == 2.0 {
		// ask user to select service time if service day is weekend
		fmt.Println("Please select the time you will be using the taxi service:")
		fmt.Println("     1 - Between midnight 12am and 6am")
		fmt.Println("     2 - None of the above")
		// assign third input number to serviceTime
		fmt.Scanln(&serviceTime)
		// condition check based on service time
		if serviceTime == 1.0 {
			// assign 2 to rushHourRatio if serviceTime is 1
			rushHourRatio = 2.0
		} else {
			// assign 1.3 to rushHourRatio if serviceTime is any number
			rushHourRatio = 1.3
		}
	}

	// ask user to input distance
	fmt.Println("Please enter the estimated distance (km) of your ride:")
	// assign fourth input number to distance
	fmt.Scanln(&distance)

	// calculate the fare
	fare = distance * modelFare * rushHourRatio

	// print the estimated fare
	fmt.Println("Your estimated fare is $", fare)
	fmt.Println("Thank you for using our service!")
}
