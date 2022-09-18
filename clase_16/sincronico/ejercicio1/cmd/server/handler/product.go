package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/bootcamp-go/Consignas-Go-Web.git/internal/domain"
	"github.com/bootcamp-go/Consignas-Go-Web.git/internal/product"
	"github.com/gin-gonic/gin"
)

type productHandler struct {
	s product.Service
}

// NewProductHandler crea un nuevo controller de productos
func NewProductHandler(s product.Service) *productHandler {
	return &productHandler{
		s: s,
	}
}

// GetAll obtiene todos los productos
func (h *productHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		products, _ := h.s.GetAll()
		c.JSON(200, products)
	}
}

// GetByID obtiene un producto por su id
func (h *productHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		product, err := h.s.GetByID(id)
		if err != nil {
			c.JSON(404, gin.H{"error": "product not found"})
			return
		}
		c.JSON(200, product)
	}
}

// Search busca un producto por precio mayor a un valor
func (h *productHandler) Search() gin.HandlerFunc {
	return func(c *gin.Context) {
		priceParam := c.Query("priceGt")
		price, err := strconv.ParseFloat(priceParam, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid price"})
			return
		}
		products, err := h.s.SearchPriceGt(price)
		if err != nil {
			c.JSON(404, gin.H{"error": "no products found"})
			return
		}
		c.JSON(200, products)
	}
}

// validateEmptys valida que los campos no esten vacios
func validateEmptys(product *domain.Product) (bool, error) {
	switch {
	case product.Name == "" || product.CodeValue == "" || product.Expiration == "":
		return false, errors.New("fields can't be empty")
	case product.Quantity <= 0 || product.Price <= 0:
		if product.Quantity <= 0 {
			return false, errors.New("quantity must be greater than 0")
		}
		if product.Price <= 0 {
			return false, errors.New("price must be greater than 0")
		}
	}
	return true, nil
}

// validateExpiration valida que la fecha de expiracion sea valida
func validateExpiration(product *domain.Product) (bool, error) {
	dates := strings.Split(product.Expiration, "/")
	list := []int{}
	if len(dates) != 3 {
		return false, errors.New("invalid expiration date, must be in format: dd/mm/yyyy")
	}
	for value := range dates {
		number, err := strconv.Atoi(dates[value])
		if err != nil {
			return false, errors.New("invalid expiration date, must be numbers")
		}
		list = append(list, number)
	}
	condition := (list[0] < 1 || list[0] > 31) && (list[1] < 1 || list[1] > 12) && (list[2] < 1 || list[2] > 9999)
	if condition {
		return false, errors.New("invalid expiration date, date must be between 1 and 31/12/9999")
	}
	return true, nil
}

// Post crear un producto nuevo
func (h *productHandler) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var product domain.Product
		err := ctx.ShouldBindJSON(&product)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid product"})
			return
		}
		valid, err := validateEmptys(&product)
		if !valid {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		valid, err = validateExpiration(&product)
		if !valid {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		p, err := h.s.Create(product)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(201, p)
	}
}

func (h *productHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idS := ctx.Param("id")

		id, err := strconv.Atoi(idS)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid id " + idS,
			})
			return
		}

		var product domain.Product

		err = ctx.ShouldBindJSON(&product)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid json ",
			})
			return
		}

		// Validar empty
		valid, err := validateEmptys(&product)
		if !valid {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Validar expoiration
		valid, err = validateExpiration(&product)
		if !valid {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Mandar al servicio

		h.s.Update(id, product)

	}
}

func (h *productHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idS := ctx.Param("id")

		id, err := strconv.Atoi(idS)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid id " + idS,
			})
		}

		pr, err := h.s.Delete(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		ctx.JSON(http.StatusOK, pr)
	}
}

func (h *productHandler) Patch() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		type Request struct {
			Name        string  `json:"name,omitempty"`
			Quantity    int     `json:"quantity,omitempty"`
			CodeValue   string  `json:"code_value,omitempty"`
			IsPublished bool    `json:"is_published,omitempty"`
			Expiration  string  `json:"expiration,omitempty"`
			Price       float64 `json:"price,omitempty"`
		}
		idS := ctx.Param("id")

		id, err := strconv.Atoi(idS)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid id " + idS,
			})
		}

		var request Request
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		// mapear Request a Product
		var product domain.Product

		product.Name = request.Name
		product.Quantity = request.Quantity
		product.CodeValue = request.CodeValue
		product.Expiration = request.Expiration
		product.Price = request.Price

		pr, err := h.s.Patch(id, product)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		ctx.JSON(http.StatusOK, pr)

	}
}
