package migration

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
)

func MysqlDriver(db *sql.DB) DriverFunc {
	return func() (string, database.Driver, error) {
		drv, err := mysql.WithInstance(db, &mysql.Config{})
		return mysqlScheme, drv, err
	}
}
