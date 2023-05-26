package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Employee struct {
	ID           uint
	Name         string
	Email        string
	Address      string
	Age          uint8
	Birthday     string
	Level        string
	IDDepartment uint
}

type Department struct {
	ID             uint
	DepartmentName string
}

func main() {
	dsn := "host=localhost user=postgres password=fbcaa96d dbname=test_db_camp port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Create
	db.Create(&Employee{ID: 1, Name: "Aditira", Email: "adit@mail.com", Address: "Jl. Raya No. 1", Age: 23, Birthday: "1998-02-21", Level: "Manager", IDDepartment: 1})
	db.Create(&Employee{ID: 2, Name: "Dito", Email: "adit@mail.com", Address: "Jl. Raya No. 2", Age: 22, Birthday: "1998-07-21", Level: "Supervisor", IDDepartment: 1})
	db.Create(&Employee{ID: 3, Name: "Ajeng", Email: "ajeng@mail.com", Address: "Jl. Raya No. 3", Age: 23, Birthday: "1998-01-01", Level: "Staff", IDDepartment: 1})
	db.Create(&Employee{ID: 4, Name: "Rahayu", Email: "ayu@mail.com", Address: "Jl. Raya No. 4", Age: 25, Birthday: "1995-04-01", Level: "Manager", IDDepartment: 2})
	db.Create(&Employee{ID: 5, Name: "Ahmad", Email: "ahmad@mail.com", Address: "Jl. Raya No. 5", Age: 21, Birthday: "1999-09-01", Level: "Staff", IDDepartment: 2})
	db.Create(&Department{ID: 1, DepartmentName: "Administration"})
	db.Create(&Department{ID: 2, DepartmentName: "Marketing"})
	db.Create(&Department{ID: 3, DepartmentName: "IT"})
	db.Create(&Department{ID: 4, DepartmentName: "Shipping"})
	db.Create(&Department{ID: 5, DepartmentName: "Public Relation"})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Insert Success")
}
