package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Post struct {
	UserId int    `json:"userId,omitempty"`
	Id     int    `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Body   string `json:"body,omitempty"`
}
type Usuario struct {
	Id       int     `json:"id,omitempty"`
	Name     string  `json:"name,omitempty"`
	UserName string  `json:"username,omitempty"`
	Email    string  `json:"email,omitempty"`
	Address  Address `json:"address,omitempty"`
	Phone    string  `json:"phone,omitempty"`
	Website  string  `json:"website,omitempty"`
	Company  Company `json:"company,omitempty"`
}
type Address struct {
	Street  string `json:"street,omitempty"`
	Suite   string `json:"suite,omitempty"`
	City    string `json:"city,omitempty"`
	ZipCode string `json:"zipcode,omitempty"`
	Geo     Geo    `json:"geo,omitempty"`
}

type Geo struct {
	Lat  string `json:"lat,omitempty"`
	Long string `json:"lng,omitempty"`
}

type Company struct {
	Name        string `json:"name,omitempty"`
	CatchPhrase string `json:"catchPhrase,omitempty"`
	Bs          string `json:"bs,omitempty"`
}

var usuarios = []Usuario{}
var posts = []Post{}

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("unexpected error ", err)
		}
	}()

	LoadUsers("data/users.json", &usuarios)
	LoadPosts("data/posts.json", &posts)

	fmt.Println(len(usuarios))
	fmt.Println(len(posts))

	router := gin.Default()

	router.GET("/users", UsersHandler)
	router.GET("/users/:id", UserHandler)

	router.GET("/posts/:id", PostHandler)

	router.GET("/search", SearchHandler)

	router.Run(":8080")

}

func UsersHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, usuarios)
}

func UserHandler(ctx *gin.Context) {

	idS := ctx.Param("id")
	id, err := strconv.ParseInt(idS, 10, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id" + idS,
		})
	}

	var empty Usuario
	user := getBy(usuarios, func(u Usuario) bool {
		return u.Id == int(id)
	})

	if user == empty {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "No hay usuario con el id " + idS,
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func PostHandler(ctx *gin.Context) {

	idS := ctx.Param("id")
	id, err := strconv.ParseInt(idS, 10, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id" + idS,
		})
	}

	postsByUser := filter(posts, func(p Post) bool {
		return p.UserId == int(id)
	})

	ctx.JSON(http.StatusOK, postsByUser)
}

func SearchHandler(ctx *gin.Context) {

	name := ctx.Query("name")

	searchList := filter(usuarios, func(u Usuario) bool {
		return strings.Contains(u.Name, name)
	})

	ctx.JSON(http.StatusOK, searchList)
}

func LoadUsers(path string, usuarios *[]Usuario) {

	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(file, usuarios)
	if err != nil {
		panic(err)
	}
}

func LoadPosts(path string, posts *[]Post) {

	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(file, posts)
	if err != nil {
		panic(err)
	}
}

func getBy[T any](slice []T, f func(T) bool) T {
	var result T

	for _, e := range slice {
		if f(e) {
			return e
		}
	}
	return result
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
