package main

import "fmt"

func main() {
	fmt.Println("Functions  in go")

	result := adder(3, 78)
	fmt.Println("Result is:", result)

	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum, message := sliceadder(values...)
	fmt.Println("Sum of values is:", sum, "Message:", message)
}

func adder(x int, y int) int {
	return x + y
}

func sliceadder(values ...int) (int, string) {
	t := 0
	for _, x := range values {
		t += x
	}
	return t, "yo nigga"
}
