package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Rate between 1 and 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n') //dont end ubtil we hit enter key

	fmt.Println("You entered", input)

	// numrating, err := strconv.ParseFloat(input, 64)
	//  In this we get an error because of \n at the end of the input , so we need to trim it first
	numrating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		// fmt.Println(err)
		panic(err)
	} else {
		fmt.Println("Added 1 to rating", numrating+1)
	}
}
