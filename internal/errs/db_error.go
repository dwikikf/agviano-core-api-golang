package errs

import (
	"errors"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func TranslateError(err error) error {

	// Not Found
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrNotFound
	}

	// MySQL Duplicate Key
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		switch mysqlErr.Number {
		case 1062:
			return ErrConflict
		case 1452:
			return ErrConflictFK
		}
	}

	return err
}
