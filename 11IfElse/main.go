package main

import "fmt"

func main() {
	fmt.Println("If Else in go")

	score := 76
	var result string

	if score >= 90 {
		result = "Topper"
	} else if score >= 50 && score < 70 {
		result = "Average"
	} else {
		result = "Failed"
	}
	fmt.Printf("Result is %v\n", result)

	if num := 3; num < 10 {
		fmt.Println("Num is < 10")
	} else {
		fmt.Println("Num is >= 10")
	}

}
