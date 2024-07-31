package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	fmt.Println("Gin REST API, starting Server...")

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbum)
	router.POST("/albums", postAlbum)
	router.DELETE("/albums/:id", deleteAlbum)

	err := router.Run("localhost:8080")
	if err != nil {
		fmt.Println("Error starting server")
	}
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
func getAlbum(c *gin.Context) {
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.JSON(http.StatusOK, a)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func postAlbum(c *gin.Context) {
	var new_album album
	if err := c.BindJSON(&new_album); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	albums = append(albums, new_album)
	c.JSON(http.StatusCreated, new_album)
}
func deleteAlbum(c *gin.Context) {
	id := c.Param("id")
	for i, a := range albums {
		if a.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "album deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
