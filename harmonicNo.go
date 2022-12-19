package main

import (
	"fmt"
)

func main() {
	var n = 0
	fmt.Print("enter the term : ")
	fmt.Scan(&n)

	fmt.Println("1")
	for i := 1; i <= n; i++ {
		fmt.Println(" + 1"+"/", i+1)

	}

}
