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

	// insert data to table `employees`
	_, err = db.Exec(`INSERT INTO employees
      VALUES (1, 'Rizki', 25, 'Jl. Kebon Jeruk', 2000000),
      (2, 'Andi', 27, 'Jl. Kebon Sirih', 3000000),
      (3, 'Budi', 30, 'Jl. Kebon Melati', 4000000),
      (4, 'Caca', 32, 'Jl. Kebon Anggrek', 5000000),
      (5, 'Deni', 35, 'Jl. Kebon Mawar', 6000000)`)

	if err != nil {
		panic(err)
	}

	fmt.Println("Insert data success")
}
