// The database package establish a new connection with MySQL database
package database

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func CreateNewMySQLStorage(config mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		panic("There was an error connecting to database")
	}

	return db, nil
}
