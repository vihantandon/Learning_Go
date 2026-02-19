package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Mod in go")
	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func greeter() {
	fmt.Println("yo mod user")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hi go developers</h1>"))
}
