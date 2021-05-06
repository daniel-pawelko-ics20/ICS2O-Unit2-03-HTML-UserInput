// Get user name and age

package main

import "fmt"

func main() {
	var userName string
	var userAge int

	fmt.Println("This program gets user's name and age")
	fmt.Println()

	fmt.Println("Enter your name: ")
	fmt.Scanln(&userName)

	fmt.Println("Enter your age: ")
	fmt.Scanln(&userAge)

	fmt.Println("Your info is:", userName, "age", userAge)
	fmt.Println("\nDone :)")
}
