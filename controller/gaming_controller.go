package controller

import (
	"fmt"
	"gin_api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitGame(c *gin.Context) {
	fmt.Print("## Game Start!!")

	hero := service.CreateNewHero()

	c.IndentedJSON(http.StatusOK, hero)
}

func GetHero(c *gin.Context) {
	fmt.Print("## GET HERO!!")

	hero := service.GetHero()

	c.IndentedJSON(http.StatusOK, hero)
}

func LearnSkill(c *gin.Context) {
	fmt.Print("## Hero learn skill!!")

	skill_id := c.Param("skill_id")

	err := service.LearnSkill(skill_id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
	}

	c.IndentedJSON(http.StatusOK, "Success")
}
