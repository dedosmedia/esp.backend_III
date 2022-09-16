package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	fmt.Printf("Hola mundo V2")

	http.HandleFunc("/hola", holHandler)

	http.ListenAndServe(":8080", nil)
}

func holHandler(w http.ResponseWriter, re *http.Request) {

	fmt.Fprintf(w, "hola v2 %s\n", os.Getenv("HOLA"))
}
