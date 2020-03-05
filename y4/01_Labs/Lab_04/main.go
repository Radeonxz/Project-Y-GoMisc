package main

import (
	"fmt"
	"strings"
)

func main() {
	var inputWord, result string
	var first = true
	fmt.Print("Enter a word:")
	fmt.Scanln(&inputWord)
	for i := len(inputWord); i > 0; i-- {
		if first {
			result = strings.ToUpper(inputWord[0:i])
			first = false
		} else {
			result = strings.ToLower(inputWord[0:i])
			first = true
		}
		fmt.Println(result)
	}
	fmt.Println("--> All done!")
}
