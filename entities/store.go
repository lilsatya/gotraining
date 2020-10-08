package entities

import (
	"time"

	"gorm.io/gorm"
)

// Store Struct
type Store struct {
	ID        int64 `gorm:"primary_key;AUTO_INCREMENT"`
	Name      string
	Location  string
	DeletedAt gorm.DeletedAt
	CreatedAt time.Time
	UpdatedAt time.Time
}
