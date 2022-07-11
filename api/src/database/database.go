package database

import (
	"api/src/configuration"
	"database/sql"
	"log"

	"github.com/denisenkom/go-mssqldb/azuread"
)

// Connect to the database
func Connect() (*sql.DB, error) {

	db, err := sql.Open(azuread.DriverName, configuration.ConnectionString)

	if err != nil {
		log.Fatal("Error connecting to database: " + err.Error())
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Error connecting to database: " + err.Error())
		db.Close()
		return nil, err
	}

	return db, nil
}
