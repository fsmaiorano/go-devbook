package configuration

import (
	"fmt"
	"os"
	"strconv"
)

var (
	// ConnectionString is the connection string to the database
	ConnectionString = ""
	// The port to listen on
	ApiPort = 0
)

// Load environment variables
func Load() {
	var err error

	ApiPort, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		ApiPort = 9000
	}

	ConnectionString =
		fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"))
}
