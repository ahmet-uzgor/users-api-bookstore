package services

import "github.com/ahmet-uzgor/users-api-bookstore/domain/users"

func CreateUser(user users.User) (*users.User, error) {
	return &user, nil
}
