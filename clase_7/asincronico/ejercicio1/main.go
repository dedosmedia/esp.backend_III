package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

type ErrorCustom struct {
	title string
}

func (err *ErrorCustom) Error() string {

	return fmt.Sprintf("Soy error custom: title: %s", err.title)
}

func main2() {

	fmt.Println("Hola")

	archivo, err := os.Open("archivo.txt")

	var myError error = fs.ErrNotExist
	//var myError *ErrorCustom

	if err != nil {
		if errors.As(err, myError) {
			fmt.Println("Error as path... ", myError)
		} else {
			fmt.Println("No Error as path... ", myError)
		}

		if errors.Is(err, myError) {
			println("Error is Path")
		} else {
			println("No Error is Path")
		}

	}

	fmt.Println(archivo.Name())

}

func main3() {

	err1 := fmt.Errorf("Unwrapped Error 1")
	err2 := fmt.Errorf("Error 2 %w", err1)
	err3 := fmt.Errorf("Error 3 %w", err2)
	err4 := errors.New("Unwrapped Error 4")

	if errors.Is(err4, err1) {
		fmt.Println("Soy err1")
	}

	errUn := errors.Unwrap(err4)

	fmt.Println(errUn)
	fmt.Println("ERR", err3)
	fmt.Println("ERR", err4)

	err5 := fmt.Errorf("... %w ...", err4)
	fmt.Println(err5, errors.Unwrap(err5))

}

type Wrap struct  {
	Err error
}

func main(){


fs.PathError
fs.ErrNotExist

}




}