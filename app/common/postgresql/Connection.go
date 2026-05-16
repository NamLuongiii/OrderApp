package postgresql

import (
	"OrderApp/persistency/table"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection() (*gorm.DB, error) {
	// 1. Đọc DB_HOST từ môi trường, nếu trống thì mặc định là localhost (cho máy Mac)
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}

	// 2. Tự động ráp vào chuỗi DSN bằng fmt.Sprintf
	dsn := fmt.Sprintf("host=%s user=postgres password=postgres dbname=orderapp port=5432 sslmode=disable", dbHost)

	db, e := gorm.Open(postgres.Open(dsn))
	if e != nil {
		return nil, e
	}

	e = db.AutoMigrate(
		&table.Product{},
		&table.Order{},
		&table.LineItem{},
		&table.User{},
	)
	if e != nil {
		return nil, e
	}
	return db, nil
}
