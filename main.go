package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/", indexHandler)

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
