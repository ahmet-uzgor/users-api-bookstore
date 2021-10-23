package users

import (
	"fmt"

	"github.com/ahmet-uzgor/users-api-bookstore/utils/errors"
)

var (
	mockUserDB = make(map[int64]*User)
)

func (user *User) Save() *errors.RestError {
	if mockUserDB[user.Id] != nil {
		return errors.BadRequestError("user has already exits!")
	}

	if mockUserDB[user.Id].Email == user.Email {
		return errors.BadRequestError("email has used for another person")
	}

	mockUserDB[user.Id] = user

	return nil
}

func (user *User) Get() *errors.RestError {
	if mockUserDB[user.Id] != nil {
		fmt.Println(mockUserDB[user.Id])
		return nil
	}

	return errors.NotFoundError("user not found")
}
