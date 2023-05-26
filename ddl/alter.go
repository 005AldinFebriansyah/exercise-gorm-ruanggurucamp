package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "fbcaa96d"
	dbname   = "test_db_camp"
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	// connect to database using func `Connect`
	db, err := Connect()
	if err != nil {
		panic(err)
	}

	// rename table employee to employees
	_, err = db.Exec(`ALTER TABLE employee RENAME TO employees`)

	if err != nil {
		panic(err)
	}

	fmt.Println("Table employee renamed to employees")
}
