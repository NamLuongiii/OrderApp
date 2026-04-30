package table

import "time"

type Product struct {
	ID        string    `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	Price     string    `json:"price" gorm:"not null"`         // hoặc float64 nếu có số thập phân
	SalePrice *string   `json:"salePrice" gorm:"default:null"` // pointer để nullable
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
