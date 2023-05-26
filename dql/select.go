package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Employee struct {
	ID      int    `db:"id"`
	Name    string `db:"name"`
	Address string `db:"address"`
	Age     int    `db:"age"`
	Salary  int    `db:"salary"`
}

func Connect() (*sql.DB, error) {
	dns := "host=localhost user=postgres password=fbcaa96d dbname=test_db_camp port=5432 sslmode=disable"

	db, err := sql.Open("postgres", dns)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	// buat koneksi ke database menggunakan func `Connect`
	db, err := Connect()
	if err != nil {
		panic(err)
	}

	// melakukan query untuk mendapatkan semua data dari tabel employee
	rows, err := db.Query("SELECT * FROM employees ORDER BY id ASC")
	if err != nil {
		panic(err)
	}

	// membuat struct baru untuk menampung data dari tabel employee
	var listEmployee []Employee

	// melakukan looping untuk menampung data dari rows ke struct Employee
	for rows.Next() {
		var employee Employee

		// kita tampung setiap baris data ke struct Employee
		err := rows.Scan(&employee.ID, &employee.Name, &employee.Age, &employee.Address, &employee.Salary)
		if err != nil {
			panic(err)
		}

		// kemudian kita tambahkan struct Employee ke listEmployee
		listEmployee = append(listEmployee, employee)
	}

	fmt.Println("Berhasil melakukan query tabel employee")
	fmt.Println(listEmployee)
}
