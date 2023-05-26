package main

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID           uint
	Name         string
	Email        string
	Age          uint8
	Birthday     string
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func main() {
	dsn := "host=localhost user=postgres password=fbcaa96d dbname=test_db_camp port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Create
	db.Create(&User{Name: "Aditira", Email: "test@mail.com", Age: 23, Birthday: "1998-02-21", MemberNumber: sql.NullString{String: "123", Valid: true}, ActivatedAt: sql.NullTime{Time: time.Now(), Valid: true}})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Insert Success")
}
