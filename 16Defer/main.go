package main

import "fmt"

// what differ does is it will execute the function after the main function is done executing
func main() {
	defer fmt.Println("hello")
	defer fmt.Println("one")
	defer fmt.Println("two") //It'll follow LIFO
	fmt.Println("Differ in go")
	myDefer()

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
