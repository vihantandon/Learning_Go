package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in go")
	content := "This needs to go in the file"
	file, err := os.Create("./file1.txt")

	checknilerror(err)

	length, err := io.WriteString(file, content)
	checknilerror(err)
	fmt.Println("Length of content written:", length)

	defer file.Close()

	readFile("./file1.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checknilerror(err)
	fmt.Println("Text data inside file: \n", databyte)
	fmt.Println("Text data inside file: \n", string(databyte))
}

func checknilerror(err error) {
	if err != nil {
		panic(err)
	}
}
