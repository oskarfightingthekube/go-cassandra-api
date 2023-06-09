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

// rewrite GetMajorByNameHandler and take name: "name" from body json
func GetMajorByNameHandler(c *gin.Context) {
	var majors []structs.MajorsWithUniversity
	var major structs.MajorName
	if err := c.BindJSON(&major); err != nil {
		fmt.Println(major)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request, check your JSON",
		})
		return
	}
	if major.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request, key values missing or empty",
		})
		return
	}
	majors, err := handlers.GetMajorByName(major.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, majors)
}

func GetMajorByTypeHandler(c *gin.Context) {
	var majors []structs.MajorsWithUniversity
	var major structs.MajorType
	if err := c.BindJSON(&major); err != nil {
		fmt.Println(major)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request, check your JSON",
		})
		return
	}
	if major.Type == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request, key values missing or empty",
		})
		return
	}
	majors, err := handlers.GetMajorByType(major.Type)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, majors)
}

func GetDepartmentsHandler(c *gin.Context) {
	departments, err := handlers.GetDepartments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, departments)
}

func GetDepartmentByUniversityHandler(c *gin.Context) {
	var departments []structs.DepartmentWithUniversity
	var department structs.Department
	if err := c.BindJSON(&department); err != nil {
		fmt.Println(department)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request, check your JSON",
		})
		return
	}
	if department.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request, key values missing or empty",
		})
		return
	}
	departments, err := handlers.GetDepartmentByUniversity(department.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, departments)
}

func AddUniversityHandler(c *gin.Context) {
	var university structs.AddUniversity
	if err := c.BindJSON(&university); err != nil {
		fmt.Println(university)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request, check your JSON",
		})
		return
	}
	if university.Name == "" || university.Country == "" || university.City == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request, key values missing or empty",
		})
		return
	}
	if err := handlers.AddUniversity(university); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "University added successfully",
	})
}

// func AddFavoriteHandler(c *gin.Context) {
// 	var favorite structs.AddFavorite
// 	if err := c.BindJSON(&favorite); err != nil {
// 		fmt.Println(favorite)
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "Invalid request, check your JSON",
// 		})
// 		return
// 	}
// 	if favorite.Name == "" || favorite.Major == "" || favorite.Type == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "Invalid request, key values missing or empty",
// 		})
// 		return
// 	}
// 	if err := handlers.AddFavorite(favorite); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusCreated, gin.H{
// 		"message": "Favorite added successfully",
// 	})
// }
