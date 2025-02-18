package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
	schema   = "active_gym_db"
)

func ConnectDB() (*sql.DB, error) {
	// Connect to database
	pgSqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable search_path=%s", host, port, user, password, dbname, schema)
	db, err := sql.Open("postgres", pgSqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected in " + dbname)
	return db, nil
}
