package routes

import (
	"fmt"
	"go-cassandra-api/handlers"
	"go-cassandra-api/structs"
	"net/http"
	"unicode"

	"github.com/gin-gonic/gin"
)

func GetUniversitiesHandler(c *gin.Context) {
	universities, err := handlers.GetUniversities()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, universities)
}

func GetMajorsHandler(c *gin.Context) {
	majors, err := handlers.GetMajors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, majors)
}

func GetMajorByNameHandler(c *gin.Context) {
	major, err := handlers.GetMajorByName(capitalize(c.Param("name")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, major)
}

func capitalize(str string) string {
	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
