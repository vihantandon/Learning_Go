package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	var ptr *int
	fmt.Println(ptr)

	num := 23
	var ptr2 *int = &num
	fmt.Println(ptr2)
	fmt.Println(*ptr2)

	*ptr2 = *ptr2 * 2
	fmt.Println(num, *ptr2)

}
