package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Empleados struct {
	Id     int
	Nombre string
	Activo bool
}

var empleados []Empleados

func init() {
	empleados = GetEmployees()
}

func main() {

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Bienvenido a la empresa Gophers!")
	})

	router.GET("/employees", func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, gin.H{
			"empleados": empleados,
		})
	})

	router.GET("/employees/:id", func(ctx *gin.Context) {

		idString := ctx.Param("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		for _, v := range empleados {
			if v.Id == id {
				ctx.JSON(http.StatusOK, gin.H{
					"empleado": v,
				})
				return
			}
		}
		ctx.String(http.StatusNotFound, "No se encuentra un usuario con id: %s ", idString)
	})

	router.GET("/employeesparams/:id/:nombre/:activo", func(ctx *gin.Context) {

		idString := ctx.Param("id")
		nombreString := ctx.Param("nombre")
		activoString := ctx.Param("activo")

		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		activo, err := strconv.ParseBool(activoString)
		if err != nil {
			panic(err)
		}

		e := Empleados{
			Id:     id,
			Nombre: nombreString,
			Activo: activo,
		}

		empleados = append(empleados, e)

		ctx.JSON(http.StatusOK, gin.H{
			"empleados": empleados,
		})

	})

	router.GET("/employeesactive", func(ctx *gin.Context) {

		activos := filter(empleados, func(e Empleados) bool {
			return e.Activo == true
		})

		ctx.JSON(http.StatusOK, gin.H{
			"activos": activos,
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

func GetEmployees() []Empleados {

	return []Empleados{
		{
			Id:     1,
			Nombre: "Diego",
			Activo: true,
		},
		{
			Id:     2,
			Nombre: "Diaz",
			Activo: false,
		},
	}

}
