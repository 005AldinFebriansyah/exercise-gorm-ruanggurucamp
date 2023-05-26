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

	// Migrate the schema
	db.AutoMigrate(&Employee{}, &Department{})

	fmt.Println("Migration Success")
}
