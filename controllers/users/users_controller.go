package users

import (
	"net/http"
	"strconv"

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

func FindUserById(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		badReqError := errors.BadRequestError("user id should be number")
		c.JSON(http.StatusBadRequest, badReqError)
		return
	}

	userFound, notFound := services.GetUserById(userId)

	if notFound != nil {
		c.JSON(http.StatusNotFound, notFound)
		return
	}

	c.JSON(http.StatusOK, userFound)
}

func UpdateUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.BadRequestError("user id should be number"))
		return
	}

	var user users.User

	if bindErr := c.ShouldBindJSON(&user); err != nil {
		badReqError := errors.BadRequestError(bindErr.Error())

		c.JSON(badReqError.Code, badReqError)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	updateResult, updateErr := services.UpdateUser(user, isPartial)
	if updateErr != nil {
		c.JSON(updateErr.Code, updateErr)
	}

	c.JSON(http.StatusOK, updateResult)
}
