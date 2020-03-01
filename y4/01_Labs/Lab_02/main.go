package main

import (
	"fmt"
)

func main() {
	var n1, n2, n3, sum int
	fmt.Print("Waiting for 3 integer numbers: ")
	fmt.Scanln(&n1)
	fmt.Scanln(&n2)
	fmt.Scanln(&n3)
	sum = n1 + n2 + n3
	if n3 > n1 && n3 > n2 {
		fmt.Print("\nn3 > n1 and n3 > n2")
	}
	if n3 < n1 && n3 < n2 {
		fmt.Print("\nn3 < n1 and n3 < n2")
	}
	if n1 == n2 && n1 == n3 {
		fmt.Print("\nn1, n2 and n3 are all the same")
	}
	if sum%5 == 0 {
		fmt.Print("\nn1 + n2 + n3 = ", sum, " a multiple of 5")
	}
	fmt.Print("\nFinished!")
}
