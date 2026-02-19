package main

import "fmt"

func main() {
	fmt.Println("Structs in go")
	//no inheritance in golang; No parent

	vihan := User{"vihan", "tandon.vihantandon@gmail.com", 21, true}
	fmt.Println(vihan)
	fmt.Printf("Name: %v\n", vihan.Name)
	fmt.Printf("All details: %+v\n", vihan)
}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}
