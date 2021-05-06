// Get user address

package main

import "fmt"

func main() {
	var streetName string
	var streetNumber int

	fmt.Println("This program gets use's street number and name.")
	fmt.Println()

	fmt.Println("Enter your street number: ")
	fmt.Scanln(&streetNumber)

	fmt.Println("Enter your street name: ")
	fmt.Scanln(&streetName)

	fmt.Println("Your street address is:", streetNumber, streetName)
	fmt.Println("\nDone :)")
}
