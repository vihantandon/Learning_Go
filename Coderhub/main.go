package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Contest struct {
	Code  string `json:"contest_code"`
	Name  string `json:"contest_name"`
	Start string `json:"contest_start_date"`
	End   string `json:"contest_end_date"`
}

type Response struct {
	FutureContests []Contest `json:"future_contests"`
}

func main() {
	url := "https://www.codechef.com/api/list/contests/all"

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var data Response
	json.NewDecoder(res.Body).Decode(&data)

	for _, c := range data.FutureContests {
		fmt.Println("Contest:", c.Name)
		fmt.Println("Start:", c.Start)
		fmt.Println("End:", c.End)
		fmt.Println("-----")
	}
}
