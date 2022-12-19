package main

import "fmt"

func main() {
	var principal, year, ROI, payment int
	var tanure = 12 * year
	fmt.Print("enter the principal : ")
	fmt.Scan(&principal)
	fmt.Print("enter year : ")
	fmt.Scan(&year)
	fmt.Print("ROI : ")
	fmt.Scan(&ROI)

	payment = (principal*ROI)/1 - (1 + ROI) ^ (-tanure)
	fmt.Println("payment is : ", payment)

}
