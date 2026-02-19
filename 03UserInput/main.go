package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin) //just a reference
	fmt.Println("Enter: ")

	input, _ := reader.ReadString('\n') //we dont care about the error so we put _
	// if we only care about the error we will put input as _
	fmt.Println("You entered", input)
	fmt.Printf("Type of the input %T", input)
}
