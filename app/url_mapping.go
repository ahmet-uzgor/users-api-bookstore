package app

import (
	"github.com/ahmet-uzgor/users-api-bookstore/controllers/health"
	"github.com/ahmet-uzgor/users-api-bookstore/controllers/users"
)

func mapUrls() {
	router.GET("/health", health.CheckHealth)
	router.GET("users/:user_id", users.FindUserById)
	router.GET("users/search", users.SearchUser)
	router.POST("/users", users.CreateUser)
	router.PUT("/users/:user_id", users.UpdateUser)
	router.PATCH("/users/:user_id", users.UpdateUser)
}
