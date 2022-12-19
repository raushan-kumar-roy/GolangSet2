package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Print("enter no : ")
	fmt.Scan(&num)

	if num%3 == 0 {
		fmt.Println("pling")
	} else if num%5 == 0 {
		fmt.Println("plang")
	} else if num%3 == 0 || num%5 == 0 {
		fmt.Println("plingPlang")
	} else if num%7 == 0 {
		fmt.Println("plong")
	} else {
		fmt.Println(num)
	}
}
