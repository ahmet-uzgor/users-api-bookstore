package app

import (
	"github.com/ahmet-uzgor/users-api-bookstore/controllers"
)

func mapUrls() {
	router.GET("/health", controllers.CheckHealth)
}
