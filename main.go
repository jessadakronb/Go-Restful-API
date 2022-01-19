package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	// fmt.Println("HELLO WORLD")

	handleRequest()
}
func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the homepage !")
}

type addressBook struct {
	Firstname string
	Lastname  string
	Code      int
	Phone     string
}

func handleRequest() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/getAddress", getAddressBookAll)
	http.ListenAndServe(":8080", nil)
}

func getAddressBookAll(w http.ResponseWriter, r *http.Request) {
	addBook := addressBook{
		Firstname: "Jessadakron",
		Lastname:  "Borisut",
		Code:      1997,
		Phone:     "0812345678",
	}
	json.NewEncoder(w).Encode(addBook)

}
