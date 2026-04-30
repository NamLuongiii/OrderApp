package table

import "time"

type User struct {
	ID        string    `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	Email     string    `json:"email" gorm:"type:varchar(255);not null;uniqueIndex:idx_email"`
	Password  string    `json:"password" gorm:"type:varchar(255);not null"`
	Role      string    `json:"role" gorm:"type:varchar(255);not null" default:"USER"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
