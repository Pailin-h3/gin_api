package main

import (
	"fmt"

	"net/http"

	"gin_api/controller"
	"gin_api/model"
	"gin_api/service"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/album/:id", getAlbumById)

	router.POST("/albums", postAlbums)

	router.GET("/game", controller.InitGame)
	router.GET("/game/learn/:skill_id", controller.LearnSkill)
	router.GET("/game/hero", controller.GetHero)

	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	fmt.Println("## GET albums")
	var albums = service.GetAlbums()
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumById(c *gin.Context) {
	fmt.Println("## GET album by id: " + c.Param("id"))
	id := c.Param("id")

	album, err := service.GetAlbumById(id)

	if err != nil {
		fmt.Println("Album not found")
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	}

	c.IndentedJSON(http.StatusOK, album)
}

func postAlbums(c *gin.Context) {
	fmt.Println("## POST albums")
	var requestBody []model.Album
	if err := c.BindJSON(&requestBody); err != nil {
		return
	}

	is_done := service.PostAlbums(requestBody)

	if !is_done {
		c.IndentedJSON(http.StatusBadRequest, nil)
	}
	c.IndentedJSON(http.StatusCreated, nil)
}
