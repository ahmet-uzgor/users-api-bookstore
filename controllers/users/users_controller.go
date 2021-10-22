package users

import (
	"net/http"

	"github.com/ahmet-uzgor/users-api-bookstore/domain/users"
	"github.com/ahmet-uzgor/users-api-bookstore/services"
	"github.com/ahmet-uzgor/users-api-bookstore/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		badReqError := errors.BadRequestError(err.Error())

		c.JSON(badReqError.Code, badReqError)
		return
	}

	result, createUserErr := services.CreateUser(user)
	if createUserErr != nil {
		c.JSON(createUserErr.Code, createUserErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func SearchUser(c *gin.Context) {}

func FindUserById(c *gin.Context) {}
