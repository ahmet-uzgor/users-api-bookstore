package users

import (
	"strings"

	"github.com/ahmet-uzgor/users-api-bookstore/utils/errors"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	CreatedDate string `json:"createdAt"`
}

func (user *User) Validate() *errors.RestError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.BadRequestError("invalid email address")
	}
	return nil
}
