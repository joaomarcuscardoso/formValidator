package config

import (
	"fmt"

	"gorm.io/gorm"
)

func InitializePostgres() (*gorm.DB, error) {
	fmt.Println("Initialize Postgres")
	// logger := GetLogger("postgres")
	/* 	dbPath := "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable" */

	return nil, nil
}
