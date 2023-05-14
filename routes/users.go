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
			"message": "Invalid request, check your JSON",
		})
		return
	}
	if user.Email == "" || user.Login == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request, key values missing or empty",
		})
		return
	}
	if err := handlers.AddUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message":  "User created",
		"user_uid": handlers.UserID,
	})
}

func LoginHandler(c *gin.Context) {
	var user structs.LoginUser
	if err := c.BindJSON(&user); err != nil {
		fmt.Println(user)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request, check your JSON",
		})
		return
	}
	if user.Login == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request, key values missing or empty",
		})
		return
	}
	if _, err := handlers.LoginUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User logged in",
	})
}

func VoteHandler(c *gin.Context) {
	var vote structs.Vote
	if err := c.BindJSON(&vote); err != nil {
		fmt.Println(vote)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request, check your JSON",
		})
		return
	}
	if vote.Login == "" || vote.University_name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request, key values missing or empty",
		})
		return
	}
	if err := handlers.Vote(vote.Login, vote.University_name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Vote created",
	})
}

func GetMyVotesHandler(c *gin.Context) {
	var user structs.LoginUser
	if err := c.BindJSON(&user); err != nil {
		fmt.Println(user)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request, check your JSON",
		})
		return
	}
	if user.Login == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request, key values missing or empty",
		})
		return
	}
	votes, err := handlers.MyVotes(user.Login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, votes)
}
