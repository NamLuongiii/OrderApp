package table

import "time"

type Product struct {
	ID        string    `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	Price     int64     `json:"price" gorm:"not null"`
	SalePrice *int64    `json:"salePrice" gorm:"default:null"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

func (p Product) GetFinalPrice() int64 {
	if p.SalePrice != nil {
		return *p.SalePrice
	}
	return p.Price
}
