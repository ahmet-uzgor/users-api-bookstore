package app

import (
	"github.com/ahmet-uzgor/users-api-bookstore/controllers"
)

func mapUrls() {
	router.GET("/health", controllers.CheckHealth)
	router.GET("users/:user_id", controllers.FindUserById)
	router.GET("users/search", controllers.SearchUser)
	router.POST("/users", controllers.CreateUser)
}
