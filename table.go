package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter any Integer between 2 to 25 : ")
	fmt.Scan(&n)

	for i := 1; i <= 10; i++ {
		fmt.Println(n, "X", i, "=", n*i)
	}

}
