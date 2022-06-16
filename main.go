package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/save", saveData)
	log.Fatal(http.ListenAndServe(":5011", nil))
}
func saveData(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "hello this is my first web in %v", "golang")
	}
}
