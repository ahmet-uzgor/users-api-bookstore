package services

import (
	"github.com/ahmet-uzgor/users-api-bookstore/domain/users"
	"github.com/ahmet-uzgor/users-api-bookstore/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserById(userId int64) (*users.User, *errors.RestError) {
	user := users.User{
		Id: userId,
	}

	if notFound := user.Get(); notFound != nil {
		return nil, notFound
	}

	return &user, nil
}
