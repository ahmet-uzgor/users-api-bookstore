package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented yet")
}

func SearchUser(c *gin.Context) {}

func FindUserById(c *gin.Context) {}
