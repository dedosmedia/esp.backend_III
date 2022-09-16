package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Name      string `json:"full_name"`
	Price     int    `json:"price,omitempty"`
	Published bool   `json:"-"`
	author    string
}

func main() {

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {

		p1 := Product{
			Name:      "Pizza",
			Price:     15,
			Published: true,
			author:    "Nadie",
		}

		p2 := Product{
			Name:      "Tinto",
			Published: false,
			author:    "Nadie",
		}

		producto, error := json.Marshal([]Product{p1, p2})

		fmt.Printf("%s", producto)

		if error != nil {
			panic(error)
		}

		c.JSON(http.StatusOK, gin.H{
			"message": string(producto),
		})
	})

	router.Run()

}
