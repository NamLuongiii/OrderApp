package postgresql

import (
	"OrderApp/common/postgresql/table"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=orderapp port=5432 sslmode=disable"
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
