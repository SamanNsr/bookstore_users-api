package mysql_utils

import (
	"strings"

	"github.com/SamanNsr/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	errNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sql, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errNoRows) {
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}
	
	switch sql.Number {
	case 1062:
		return errors.NewBadRequestError("invalid data")			
	}
	return errors.NewInternalServerError("error proccessing request")
}