package routes

import (
	"go-cassandra-api/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})
	r.GET("/users", GetUsersHandler)
	if err := r.Run(); err != nil {
		panic(err)
	}
}

func GetUsersHandler(c *gin.Context) {
	users, err := handlers.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, users)
}
