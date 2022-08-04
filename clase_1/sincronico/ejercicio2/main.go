package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "hola mundo"
	letters, count := Count2(str)

	fmt.Printf("Longitud: %v conteo: %v \n", len(letters), count)

	Print(letters)

}

// Opción 1
func Count(word string) (letters []string, count int) {
	letters = strings.Split(word, "")
	count = len(letters)
	fmt.Printf("LEN: %v ", count)
	return
}

// Opción 2
func Count2(word string) (letters []string, count int) {

	i := 0
	for k, v := range word {
		fmt.Printf("%v ", string(v))
		letters = append(letters, fmt.Sprintf("%s", string(v)))
		i = k
	}

	count = i + 1
	return
}

func Print(letters []string) {

	count := 0
	for i := 0; i < len(letters); i++ {
		count++
		fmt.Println(letters[i])
	}
}
