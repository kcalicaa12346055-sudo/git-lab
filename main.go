package main

import (
	"fmt"
)

func main() {
	var amount float64
	var choice int

	fmt.Println("--- PHP Currency Converter ---")
	fmt.Print("Enter amount in PHP: ")
	fmt.Scan(&amount)

	fmt.Println("Choose target currency:")
	fmt.Println("1. USD (0.018)")
	fmt.Println("2. EUR (0.016)")
	fmt.Print("Selection: ")
	fmt.Scan(&choice)

	convert(amount, choice)
}

func convert(amt float64, res int) {
	if res == 1 {
		fmt.Printf("%.2f PHP is %.2f USD\n", amt, amt*0.018)
	} else if res == 2 {
		fmt.Printf("%.2f PHP is %.2f EUR\n", amt, amt*0.016)
	} else {
		fmt.Println("Invalid choice.")
	}
}