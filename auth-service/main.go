package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/hello", helloHandle)

	err := http.ListenAndServe(":8081", r) // server <-
	if err != nil {
		log.Fatalf("Error starting server")
	}
	fmt.Println("Connected")

}
func helloHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}
