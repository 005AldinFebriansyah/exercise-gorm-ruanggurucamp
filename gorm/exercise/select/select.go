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

	var result1 Employee
	db.Raw("SELECT id, name, age FROM employees WHERE id = ?", 3).Scan(&result1)
	fmt.Println(result1)

	db.Raw("SELECT id, name, age FROM employees WHERE name = ?", "Dito").Scan(&result1)
	fmt.Println(result1)

	var result2 Department
	db.Raw("SELECT id, department_name FROM departments WHERE id = ?", 5).Scan(&result2)
	fmt.Println(result2)

	db.Raw("SELECT id, department_name FROM departments WHERE department_name = ?", "IT").Scan(&result2)
	fmt.Println(result2)

	// Update tanpa kondisi
	db.Model(&Employee{}).Where("id = ?", 1).Update("name", "Jamhuri")

	//Delete
	db.Where("name = ?", "Imam").Delete(&Employee{})
}
