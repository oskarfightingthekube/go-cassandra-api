package routes

import (
	"fmt"
	"go-cassandra-api/handlers"
	"go-cassandra-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

func GetUserHandler(c *gin.Context) {
	user, err := handlers.GetUser(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

func AddUserHandler(c *gin.Context) {
	var user structs.AddUser
	if err := c.BindJSON(&user); err != nil {
		fmt.Println(user)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	if user.Email == "" || user.Login == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	if err := handlers.AddUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created",
	})
}
