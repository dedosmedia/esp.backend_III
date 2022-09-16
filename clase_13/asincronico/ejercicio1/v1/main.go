package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	fmt.Printf("Hola mundo V1")

	http.HandleFunc("/hola", holaHandler)

	http.ListenAndServe(":8080", nil)

}

func holaHandler(w http.ResponseWriter, re *http.Request) {

	fmt.Fprintf(w, "hola V1%s\n", os.Getenv("HOLA"))
}
