package routes

import (
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	// get request to get all users
	r.GET("/users", GetUsersHandler)
	// get request to get a user by id
	r.GET("/users/:id", GetUserHandler)
	// post request to add a user
	r.POST("/adduser", AddUserHandler)
	// login user
	r.GET("/login", LoginHandler)

	/* ----------------------------------- */

	r.POST("/vote", VoteHandler)
	// r.POST("/addfavorite", AddFavoriteHandler)

	/* ----------------------------------- */

	r.GET("/universities", GetUniversitiesHandler)
	r.GET("/majors", GetMajorsHandler)
	r.GET("/majors/name", GetMajorByNameHandler)
	r.GET("/majors/type", GetMajorByTypeHandler)
	r.GET("/departments", GetDepartmentsHandler)
	r.GET("/departments/university", GetDepartmentByUniversityHandler)
	r.POST("/adduniversity", AddUniversityHandler)
	// r.POST("/addmajor", AddMajorHandler)
	// r.POST("/adddepartment", AddDepartmentHandler)

	if err := r.Run(); err != nil {
		panic(err)
	}
}
