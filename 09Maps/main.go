package main

import "fmt"

func main() {
	fmt.Println("Maps")
	m := make(map[string]string)
	m["name"] = "go"
	m["type"] = "programming language"
	m["origin"] = "google"

	fmt.Printf("%v\n", m["origin"])

	delete(m, "origin")

	for key, val := range m {
		fmt.Printf("key: %v , value: %v\n", key, val)
	}
}
