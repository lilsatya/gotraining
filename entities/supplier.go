package entities

import (
	"time"

	"gorm.io/gorm"
)

// Supplier Struct
type Supplier struct {
	ID        int64 `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name      string
	Location  string
	Phone     string
	DeletedAt gorm.DeletedAt
	CreatedAt time.Time
	UpdatedAt time.Time
}
