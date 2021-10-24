package users

import (
	usersdb "github.com/ahmet-uzgor/users-api-bookstore/database/mysql/users_db"
	dateutil "github.com/ahmet-uzgor/users-api-bookstore/utils/date_util"
	"github.com/ahmet-uzgor/users-api-bookstore/utils/errors"
)

var (
	insertUserQuery  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	getUserByIDQuery = "SELECT * FROM users WHERE id=?;"
)

func (user *User) Save() *errors.RestError {
	statement, err := usersdb.Client.Prepare(insertUserQuery)
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
	statement, err := usersdb.Client.Prepare(getUserByIDQuery)
	if err != nil {
		return errors.IntervalServerError(err.Error())
	}

	defer statement.Close()

	queryResult := statement.QueryRow(user.Id)
	if err := queryResult.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreatedDate); err != nil {
		return errors.IntervalServerError(err.Error())
	}

	return nil
}
