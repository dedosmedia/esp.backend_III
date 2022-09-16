package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Productos struct {
	Id              int
	Nombre          string
	Precio          int
	Stock           int
	Codigo          int
	Publicado       bool
	FechaDeCreacion string
}

func main() {

	file, err := os.ReadFile("productos.json")
	if err != nil {
		panic(err)
	}

	var lista []Productos

	err = json.Unmarshal([]byte(file), &lista)
	if err != nil {
		panic(err)
	}

	fmt.Println(lista)

	router := gin.Default()

	router.GET("/productos", func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, gin.H{
			"productos": lista,
		})

	})

	router.Run()

}
