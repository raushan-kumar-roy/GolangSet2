package main

import (
	"fmt"
	"strings"
)

func main() {

	var input string
	fmt.Print("Enter a string: ")
	fmt.Scanln(&input)

	aPresent := strings.Contains(input, "a")
	ePresent := strings.Contains(input, "e")
	pPresent := strings.Contains(input, "p")

	if aPresent && ePresent && pPresent {
		fmt.Println("All Present")
	} else if aPresent || ePresent || pPresent {
		fmt.Println("One or more - Present")
	} else {
		fmt.Println("None Present")
	}
}
