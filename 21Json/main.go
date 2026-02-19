package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` //to change the name of the key in json like alias
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              //to ignore this field in json
	Tags     []string `json:"tags,omitempty"` //to ignore this field if it is empty in json
}

func main() {
	fmt.Println("Json in go")
	// EncodedJson()
	DecodeJson()
}

func EncodedJson() {

	data := []course{
		{"ReactJS Bootcamp", 299, "Udemy", "abc123", []string{"web-dev", "js"}},
		{"Angular Bootcamp", 199, "Udemy", "def456", nil},
		{"Go Bootcamp", 399, "Udemy", "ghi789", []string{"web-dev", "go"}},
	}

	//package this data as JSON data

	// finalJson, _ := json.Marshal(data)
	finalJson, _ := json.MarshalIndent(data, "", "\t") //indent on \t
	fmt.Println(string(finalJson))
}

func DecodeJson() {

	//any data from web is in byte format
	JsonDataFromWeb := []byte(` 
	{
        "coursename": "Go Bootcamp",
        "Price": 399,
        "website": "Udemy",
         "tags": ["web-dev","go"]
    }
	`)

	var courseData course

	isValid := json.Valid(JsonDataFromWeb)

	if isValid {
		fmt.Println("Json was valid")
		json.Unmarshal(JsonDataFromWeb, &courseData)
		fmt.Printf("%#v\n", courseData)
	} else {
		fmt.Println("Json was not valid")
	}

	//some cases where u want to add data to key values
	var myOnlineData map[string]interface{} //interface{} is a type that can hold any value

	json.Unmarshal(JsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key : %v and Value : %v \n", k, v)
	}
}
