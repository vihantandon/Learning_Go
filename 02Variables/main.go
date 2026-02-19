package main

import (
	"fmt"
)

const PI = 3.14

func main() {
	var name string = "Vihan"
	fmt.Println(name)
	fmt.Printf("Variable is of type: %T \n", name)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallval uint8 = 255
	fmt.Println(smallval)
	fmt.Printf("Variable is of type: %T \n", smallval)

	var floatval float64 = 255.9867896756779
	fmt.Println(floatval)
	fmt.Printf("Variable is of type: %T \n", floatval)

	// Default values and aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type
	var website = "example.com"
	fmt.Println(website)

	//no var style
	numberOfUsers := 3000 //This is not allowed outside functions
	fmt.Println(numberOfUsers)

	fmt.Println(PI)
	fmt.Printf("Variable is of type: %T \n", PI)
}
