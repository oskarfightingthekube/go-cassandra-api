package routes

import (
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.New()
	// get request to get all users
	r.GET("/users", GetUsersHandler)
	// get request to get a user by id
	r.GET("/users/:id", GetUserHandler)
	// post request to add a user
	r.POST("/adduser", AddUserHandler)
	// login user
	r.GET("/login", LoginHandler)
	/* ----------------------------------- */
	// r.GET("/universities", GetUniversitiesHandler)
	if err := r.Run(); err != nil {
		panic(err)
	}
}
