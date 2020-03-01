package main

import (
	"fmt"
)

func main() {
	var coffees, muffins, cookies int
	fmt.Print("How many coffees? ")
	fmt.Scanln(&coffees)
	fmt.Print("How many muffins and cookies? ")
	fmt.Scanln(&muffins)
	fmt.Scanln(&cookies)
	if (coffees + cookies) > 100 {
		fmt.Print("\nAre you feeding an army?")
	} else {
		if coffees < ((muffins + cookies) / 2) {
			fmt.Print("\nThere are some very hungry people there.")
		} else {
			if coffees > (muffins + cookies) {
				fmt.Print("\nThis is a health conscious group.")
			} else {
				if coffees == muffins && coffees == cookies {
					fmt.Print("\nSame number of cookies, muffins and coffees ordered.")
				} else {
					if coffees < 100 && coffees > 35 {
						fmt.Print("\nHope you have a big conference room.")
					}
				}
			}
		}
	}
	fmt.Print("\nBon Appetit!!!")
}
