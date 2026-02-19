package main

import "fmt"

func main() {
	fmt.Println("Array")

	var A [5]int
	A[0] = 9
	fmt.Println(A)
	fmt.Println(len(A))

	var B [4]string
	B[0] = "hello"
	B[2] = "world"
	fmt.Println(B)

	var C = [2]string{"go", "lang"}
	fmt.Println(C)
}
