package routes

import (
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
	r.GET("/users/:id", GetUserHandler)
	if err := r.Run(); err != nil {
		panic(err)
	}
}
