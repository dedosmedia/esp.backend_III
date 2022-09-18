package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Producto struct {
	Id         int     `json:"id,omitempty"`
	Name       string  `json:"name,omitempty" binding:"required"`
	Quantity   int     `json:"quantity,omitempty"  binding:"required"`
	Code       string  `json:"code_value,omitempty" binding:"required"`
	Published  bool    `json:"is_published,omitempty" default:"false"`
	Expiration string  `json:"expiration,omitempty" binding:"required"`
	Price      float64 `json:"price,omitempty" binding:"required"`
}

var productos []Producto

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("unexpected error ", err)
		}
	}()

	router := gin.Default()

	router.GET("/productparams", func(ctx *gin.Context) {
		idS := ctx.Query("id")

		id, err := strconv.ParseInt(idS, 10, 0)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("invalid id %s", idS),
			})
			return
		}

		name := ctx.Query("name")
		quantityS := ctx.Query("quantity")
		quantity, err := strconv.ParseInt(quantityS, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("invalid quantity %s", quantityS),
			})
			return
		}
		code := ctx.Query("code_value")
		publishedS := ctx.Query("is_published")
		published, err := strconv.ParseBool(publishedS)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("invalid is_published %s", publishedS),
			})
			return
		}
		expiration := ctx.Query("expiration")
		priceS := ctx.Query("price")
		price, err := strconv.ParseFloat(priceS, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("invalid price %s", priceS),
			})
			return
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

	router.POST("/products", AddProduct)
	router.GET("/products", GetProduct)
	//router.GET("/products/:id", GetProductById)

	router.GET("/products/:id", func(ctx *gin.Context) {

		idS := ctx.Param("id")

		id, err := strconv.ParseInt(idS, 10, 0)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("invalid id %s", idS),
			})
			return
		}

		ps := filter(productos, func(p Producto) bool {
			return p.Id == int(id)
		})

		if len(ps) > 1 {
			ctx.JSON(http.StatusOK, ps[0])
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf("No hay productos con id %s", idS),
			})
		}

	})

	router.GET("/searchbyquantity", func(ctx *gin.Context) {

		minS := ctx.Query("min")
		min, err := strconv.ParseInt(minS, 10, 0)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("invalid min %s", minS),
			})
			return
		}
		maxS := ctx.Query("max")
		max, err := strconv.ParseInt(maxS, 10, 0)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("invalid max %s", maxS),
			})
			return
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
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("invalid amount %s", amountS),
			})
			return
		}

		ps := filter(productos, func(p Producto) bool {
			return p.Code == code
		})
		total := 0.0
		if len(ps) > 0 {
			total = ps[0].Price * float64(amount)
		} else {

			ctx.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf("product not found with code %s", code),
			})
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

func AddProduct(ctx *gin.Context) {

	var product Producto

	err := ctx.ShouldBindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	product.Id = len(productos) + 1

	code := product.Code

	any := filter(productos, func(p Producto) bool {
		return p.Code == code
	})
	if len(any) > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Ya existe producto con code " + code,
		})
		return
	}

	productos = append(productos, product)

	ctx.JSON(http.StatusOK, product)

}

func GetProduct(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, productos)

}

func GetProductById(ctx *gin.Context) {

	idS := ctx.Param("id")
	id, err := strconv.ParseInt(idS, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	any := filter(productos, func(p Producto) bool {
		return p.Id == int(id)
	})

	if len(any) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "No hay ningún producto con el id " + idS,
		})
		return
	}

	ctx.JSON(http.StatusOK, any[0])

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
