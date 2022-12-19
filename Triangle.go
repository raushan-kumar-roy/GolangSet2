package main

import "fmt"

func main() {
	var side1, side2, side3 int
	fmt.Print("enter value of side 1 : ")
	fmt.Scan(&side1)
	fmt.Print("enter value of side 2 : ")
	fmt.Scan(&side2)
	fmt.Print("enter value of side 3 : ")
	fmt.Scan(&side3)

	if side1 == side2 || side2 == side3 || side3 == side1 {
		fmt.Println("this is equilateral triangle")
	} else if side1 != side2 || side2 != side3 || side3 != side1 {
		fmt.Println("this is scalene triangle")
	} else if side1 == side2 || side2 != side3 || side3 != side1 {
		fmt.Println("this is isoscales triangle")
	}
}
