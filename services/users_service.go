package services

import (
	usersdb "github.com/ahmet-uzgor/users-api-bookstore/database/mysql/users_db"
	"github.com/ahmet-uzgor/users-api-bookstore/domain/users"
	"github.com/ahmet-uzgor/users-api-bookstore/utils/errors"
)

var (
	UsersService UserServiceInterface = &usersService{}
)

type usersService struct {
}

type UserServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestError)
	GetUserById(int64) (*users.User, *errors.RestError)
	UpdateUser(users.User, bool) (*users.User, *errors.RestError)
}

func (u *usersService) CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *usersService) GetUserById(userId int64) (*users.User, *errors.RestError) {
	if err := usersdb.Client.Ping(); err != nil {
		panic(err)
	}

	user := users.User{
		Id: userId,
	}

	if notFound := user.Get(); notFound != nil {
		return nil, notFound
	}

	return &user, nil
}

func (u *usersService) UpdateUser(user users.User, isPartial bool) (*users.User, *errors.RestError) {
	current := users.User{
		Id: user.Id,
	}
	if err := current.Get(); err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName == "" {
			user.FirstName = current.FirstName
		}
		if user.LastName == "" {
			user.LastName = current.LastName
		}
		if user.Email == "" {
			user.Email = current.FirstName
		}
	}

	if err := user.Update(); err != nil {
		return nil, err
	}

	return &user, nil
}
