package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ablum struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var ablums = []ablum{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAblum(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, ablums)
}

func postAlbum(c *gin.Context) {
	var newAblum ablum

	if err := c.BindJSON(&newAblum); err != nil {
		return
	}

	ablums = append(ablums, newAblum)

	c.IndentedJSON(http.StatusCreated, newAblum)
}

func getAblumById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range ablums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album is not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAblum)
	router.GET("/albums/:id", getAblumById)
	router.POST("/albums", postAlbum)

	router.Run("localhost:8080")
}
