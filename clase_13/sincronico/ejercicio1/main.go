package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Persona struct {
	Nombre    string
	Apellido  string
	Edad      int
	Dirección string
	Teléfono  string
	Activo    bool
}

func main() {

	router := gin.Default()

	jsonData := `{"Nombre":"Pedro","Apellido":"Pérez","Edad":45,"Dirección":"Street","Teléfono":"15452","Activo":false}`

	var p1 Persona

	error := json.Unmarshal([]byte(jsonData), &p1)

	if error != nil {
		panic(error)
	}

	fmt.Println(p1)

	p2 := Persona{
		Nombre:    "Diego",
		Apellido:  "Diaz",
		Edad:      40,
		Dirección: "Call 13",
		Teléfono:  "74515",
		Activo:    false,
	}

	fmt.Println(p2)

	router.GET("/persona", func(ctx *gin.Context) {

		response, err := json.Marshal(p2)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(response))

		ctx.JSON(200, gin.H{
			"persona": p2,
		})
	})

	router.Run()

}
