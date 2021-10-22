package users

import (
	"fmt"
	"net/http"

	"github.com/ahmet-uzgor/users-api-bookstore/domain/users"
	"github.com/ahmet-uzgor/users-api-bookstore/services"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err.Error())
		return
	}

	result, createUserErr := services.CreateUser(user)
	if createUserErr != nil {
		fmt.Println(createUserErr.Error())
		return
	}
	c.JSON(http.StatusCreated, result)
}

func SearchUser(c *gin.Context) {}

func FindUserById(c *gin.Context) {}
