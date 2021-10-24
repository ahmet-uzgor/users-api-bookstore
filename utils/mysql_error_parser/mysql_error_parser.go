package mysqlerrorparser

import (
	"strings"

	"github.com/ahmet-uzgor/users-api-bookstore/utils/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	noRowsError = "no rows in result set"
)

func ParseError(err error) *errors.RestError {
	sqlError, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), noRowsError) {
			return errors.NotFoundError("no record with given id")
		}

		return errors.IntervalServerError("unparsed database error")
	}

	switch sqlError.Number {
	case 1062:
		return errors.BadRequestError(err.Error())
	}

	return errors.IntervalServerError("unparsed database error")
}
