package routes

import (
	"fmt"
	"go-cassandra-api/handlers"
	"go-cassandra-api/structs"
	"net/http"

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
