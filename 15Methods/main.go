package main

import "fmt"

func main() {
	fmt.Println("Methods in go")
	//no inheritance in golang; No parent

	vihan := User{"vihan", "tandon.vihantandon@gmail.com", 21, true}
	fmt.Println(vihan)
	fmt.Printf("Name: %v\n", vihan.Name)
	fmt.Printf("All details: %+v\n", vihan)
	vihan.GetStatus()
	vihan.NewMail() // this will not change the email of the user because we are passing the value of the user and not the pointer to the user. To change the email of the user we need to pass the pointer
}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@mail.com"
	fmt.Println("Email of the user is: ", u.Email)
}
