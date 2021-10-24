package users

import (
	"fmt"

	usersdb "github.com/ahmet-uzgor/users-api-bookstore/database/mysql/users_db"
	dateutil "github.com/ahmet-uzgor/users-api-bookstore/utils/date_util"
	"github.com/ahmet-uzgor/users-api-bookstore/utils/errors"
)

var (
	mockUserDB  = make(map[int64]*User)
	insertQuery = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
)

func (user *User) Save() *errors.RestError {
	statement, err := usersdb.Client.Prepare(insertQuery)
	if err != nil {
		return errors.IntervalServerError(err.Error())
	}

	defer statement.Close()

	insertResult, dbErr := statement.Exec(user.FirstName, user.LastName, user.Email, dateutil.GetTodayDbString())
	if dbErr != nil {
		return errors.IntervalServerError(dbErr.Error())
	}

	userId, idErr := insertResult.LastInsertId()
	if idErr != nil {
		return errors.IntervalServerError(err.Error())
	}

	user.Id = userId

	return nil
}

func (user *User) Get() *errors.RestError {
	if mockUserDB[user.Id] != nil {
		fmt.Println(mockUserDB[user.Id])
		return nil
	}

	return errors.NotFoundError("user not found")
}
