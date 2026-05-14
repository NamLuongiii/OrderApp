package table

import "time"

type Order struct {
	ID        string    `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Total     int64     `json:"total" gorm:"not null"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	Email     string    `json:"email" gorm:"type:varchar(255)"`
	Phone     string    `json:"phone" gorm:"type:varchar(50)"`
	Address   string    `json:"address" gorm:"type:text"`
	Note      string    `json:"note" gorm:"type:text"`
	Status    string    `json:"status" gorm:"type:varchar(50);default:'PENDING'"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`

	// Relationship
	LineItems []LineItem `json:"lineItems" gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
}
