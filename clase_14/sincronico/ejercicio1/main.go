package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Producto struct {
	Id         int     `json:"id,omitempty"`
	Name       string  `json:"name,omitempty"`
	Quantity   int     `json:"quantity,omitempty"`
	Code       string  `json:"code_value,omitempty"`
	Published  bool    `json:"is_published,omitempty"`
	Expiration string  `json:"expiration,omitempty"`
	Price      float64 `json:"price,omitempty"`
}

var productos []Producto

func main() {

	router := gin.Default()

	router.GET("/productparams", func(ctx *gin.Context) {
		idS := ctx.Query("id")
		id, err := strconv.ParseInt(idS, 10, 0)
		if err != nil {
			panic(err)
		}

		name := ctx.Query("name")
		quantityS := ctx.Query("quantity")
		quantity, err := strconv.ParseInt(quantityS, 10, 64)
		if err != nil {
			panic(err)
		}
		code := ctx.Query("code_value")
		publishedS := ctx.Query("is_published")
		published, err := strconv.ParseBool(publishedS)
		if err != nil {
			panic(err)
		}
		expiration := ctx.Query("expiration")
		priceS := ctx.Query("price")
		price, err := strconv.ParseFloat(priceS, 64)
		if err != nil {
			panic(err)
		}
		p := Producto{
			Id:         int(id),
			Name:       name,
			Quantity:   int(quantity),
			Code:       code,
			Published:  published,
			Expiration: expiration,
			Price:      price,
		}

		productos = append(productos, p)

		ctx.JSON(http.StatusOK, p)

	})

	router.GET("/products/:id", func(ctx *gin.Context) {

		idS := ctx.Param("id")

		id, err := strconv.ParseInt(idS, 10, 0)
		if err != nil {
			panic(err)
		}

		ps := filter(productos, func(p Producto) bool {
			return p.Id == int(id)
		})

		if len(ps) > 1 {
			ctx.JSON(http.StatusOK, ps[0])
		} else {
			ctx.String(http.StatusNotFound, "No hay un producto con id %s", idS)
		}

	})

	router.GET("/searchbyquantity", func(ctx *gin.Context) {

		minS := ctx.Query("min")
		min, err := strconv.ParseInt(minS, 10, 0)
		if err != nil {
			panic(err)
		}
		maxS := ctx.Query("max")
		max, err := strconv.ParseInt(maxS, 10, 0)
		if err != nil {
			panic(err)
		}

		ps := filter(productos, func(p Producto) bool {

			return p.Quantity >= int(min) && p.Quantity <= int(max)
		})

		ctx.JSON(http.StatusOK, ps)

	})

	router.GET("/buy/:code_value/amount/:amount", func(ctx *gin.Context) {

		code := ctx.Param("code_value")
		amountS := ctx.Param("amount")
		amount, err := strconv.ParseInt(amountS, 10, 0)
		if err != nil {
			panic(err)
		}

		ps := filter(productos, func(p Producto) bool {
			return p.Code == code
		})
		total := 0.0
		if len(ps) > 0 {
			total = ps[0].Price * float64(amount)
		} else {
			ctx.String(http.StatusNotFound, "No se encuentra el producto con code: %s", code)
			return
		}

		ctx.JSON(http.StatusOK, Producto{
			Name:     ps[0].Name,
			Quantity: int(amount),
			Price:    total,
		})

	})

	router.Run(":8080")

}

func filter[T any](slice []T, f func(T) bool) []T {
	var n []T
	for _, e := range slice {
		if f(e) {
			n = append(n, e)
		}
	}
	return n
}
