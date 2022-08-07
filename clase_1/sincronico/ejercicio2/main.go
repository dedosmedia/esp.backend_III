package main

import "fmt"

// Realizar una aplicación que contenda una variable con el número del mes
// Según el número, imprimir el mes al que corresponda en texto

func main() {

	month := 4

	fmt.Println("Mes:", Month(month))
}

func Month(month int) (monthName string) {

	switch month {
	case 1:
		monthName = "Enero"
	case 2:
		monthName = "Febrero"
	case 3:
		monthName = "Marzo"
	case 4:
		monthName = "Abril"
	case 5:
		monthName = "Mayo"
	case 6:
		monthName = "Junio"
	case 7:
		monthName = "Julio"
	case 8:
		monthName = "Agosto"
	case 9:
		monthName = "Septiembre"
	case 10:
		monthName = "Octubre"
	case 11:
		monthName = "Noviembre"
	case 12:
		monthName = "Diciembre"
	}

	return
}
