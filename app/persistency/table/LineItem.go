package table

import (
	"time"
)

type LineItem struct {
	ID          string    `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	OrderID     string    `json:"orderId" gorm:"type:uuid;not null;index:idx_order_id"`
	ProductID   string    `json:"productId" gorm:"type:uuid;not null"`
	Quantity    int64     `json:"quantity" gorm:"not null"`
	Price       int64     `json:"price" gorm:"not null"`
	Total       int64     `json:"total" gorm:"not null"`
	ProductName string    `json:"productName" gorm:"type:varchar(255)"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`

	// Relationship
	Product *Product `json:"product,omitempty" gorm:"foreignKey:ProductID;references:ID"`
}
